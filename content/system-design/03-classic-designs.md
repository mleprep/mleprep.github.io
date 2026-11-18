# Classic Designs

The most frequently asked design questions. Each one follows the same structure: requirements → estimates → API & data → architecture → deep dives.

## URL Shortener (TinyURL / bit.ly)

**Requirements**: shorten a long URL, redirect short → long with low latency, optional custom alias & expiry, analytics. Read-heavy (100:1).

**Estimates**: 100 M new URLs/day ≈ 1.2k writes/s; 10 B redirects/day ≈ 115k reads/s. 5 yrs × 100 M × 365 ≈ 180 B URLs × ~500 B ≈ 90 TB. A 7-char base62 code gives 62⁷ ≈ 3.5 × 10¹² codes — plenty.

**API**: `POST /urls {longUrl, alias?, ttl?} → shortUrl`, `GET /{code} → 301/302 Location: longUrl`.

```viz
            write path                                  read path
client -> API -> ID generator (ranges / Snowflake)   client -> LB -> redirect svc -> Redis (code -> url)
              -> base62(id) = "aZ3k9Q1"                                          | miss
              -> KV store {code: longUrl, owner, exp}                            v
                                                                            KV store (Cassandra/Dynamo)
                         analytics: redirect svc -> Kafka -> stream agg -> OLAP
```

**Deep dives**
- **Code generation**: (a) hash(longUrl) → take 7 chars, handle collisions by retry/salt; (b) **counter + base62** — no collisions; give each app server a pre-allocated ID range (from ZooKeeper / a ticket DB) so there's no per-request coordination. Add a random offset or shuffle if sequential codes must not be guessable.
- **301 vs 302**: 301 is cached by browsers (less load, loses analytics); 302 keeps every hit on your servers.
- **Storage**: simple key → value lookups at huge scale → partitioned KV store keyed by `code`. Cache the hot 20% of URLs (Zipf distribution) in Redis.
- **Expiry**: TTL in the store + lazy deletion on read + background cleanup.

## Rate Limiter

**Requirements**: limit requests per user/IP/API key (e.g. 100/min), low added latency (<1 ms), distributed across many API servers, informative 429 responses, configurable rules.

```viz
client -> API gateway / middleware --check--> Redis cluster (counters, Lua scripts)
                 |  allowed                       ^
                 v                                | rules (cached locally, from config store)
             backend services
```

**Deep dives**
- **Algorithm**: token bucket (bursty but bounded) or sliding window counter (accurate, cheap) — see Building Blocks.
- **Atomicity**: read-modify-write races across servers → do it in one Redis Lua script, or `INCR` with `EXPIRE` for fixed windows.
- **Placement**: at the API gateway (centralized) or as a sidecar/library; client-side limits are only advisory.
- **Scale & latency**: shard counters by key; keep a small local token cache and sync periodically (allows slight overshoot) to avoid a Redis hop per request.
- **Failure mode**: if Redis is down, **fail open** (availability) or **fail closed** (protection) — decide per API.
- Headers: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `Retry-After`.

## Key-Value Store (Dynamo-style)

**Requirements**: `put(key, value)`, `get(key)`; highly available, partition-tolerant, horizontally scalable, tunable consistency.

```viz
client -> any node (coordinator)
               | consistent hashing ring with virtual nodes
               v
   key k -> preference list [N1, N2, N3]  (N = 3 replicas)
   write: send to all 3, succeed after W = 2 acks
   read : ask R = 2, return newest version, repair stale replicas
```

**Deep dives**
- **Partitioning**: consistent hashing + virtual nodes.
- **Replication**: N replicas on distinct nodes/AZs; quorum `R + W > N` for overlap; **sloppy quorum + hinted handoff** to stay writable during failures.
- **Conflicts**: vector clocks to detect concurrent writes; resolve via last-write-wins (timestamps) or return siblings to the client.
- **Failure detection**: gossip protocol; **anti-entropy** with Merkle trees to sync replicas cheaply.
- **Storage engine**: LSM-tree (memtable + WAL + SSTables + compaction), Bloom filters for reads.

## News Feed (Twitter / Instagram / Facebook)

**Requirements**: post, follow, view home timeline (reverse-chronological or ranked), low-latency feed reads; eventual consistency is fine.

**Estimates**: 200 M DAU, 2 posts/day → 5k writes/s; 20 feed loads/day → ~50k reads/s. Average 200 followers; celebrities have 100 M.

```viz
post:  client -> post svc -> posts DB (sharded by post_id) + media -> object store/CDN
                         -> Kafka "new_post" -> fan-out workers
                                                 | for each follower (non-celebrity author)
                                                 v
                                     feed cache: Redis list per user (post ids, capped ~800)

read:  client -> feed svc -> Redis feed(user) ──merge── celebrity posts pulled at read time
                          -> hydrate ids via post cache -> rank -> page with cursor
```

**Deep dives**
- **Fan-out on write (push)**: precompute each follower's feed → fast reads, expensive writes for big accounts.
- **Fan-out on read (pull)**: build the feed at read time from followees → cheap writes, slow reads.
- **Hybrid** (the classic answer): push for normal users, pull for celebrities (and for inactive users, don't push at all).
- **Pagination**: cursor-based (`max_id`) instead of offsets.
- **Ranking**: candidate generation → feature lookup → ML scoring → re-ranking with diversity rules.
- **Data**: follow graph in a sharded store (`follower → followees` and reverse index); posts keyed by `post_id` (Snowflake = time-sortable).

## Chat System (WhatsApp / Messenger)

**Requirements**: 1:1 and group messages, online presence, delivery/read receipts, offline delivery, message history across devices. Low latency.

```viz
client A ==WebSocket==> chat server 1 --> message svc --> Kafka --> chat server 7 ==WS==> client B
                              |                |                         ^
                              |                v                         | which server holds B?
                              |        messages DB (Cassandra,           |
                              |        partition = conversation_id,  session registry (Redis:
                              |        clustering = message_id)       user -> server)
                              v
                    presence svc (heartbeats, TTL in Redis)     B offline -> push notification svc
```

**Deep dives**
- **Transport**: WebSockets (bi-directional, persistent); long polling as a fallback. Chat servers are stateful (hold connections) → route via the session registry.
- **Ordering**: per-conversation monotonically increasing `message_id` (Snowflake or per-chat sequence); clients sort and dedupe.
- **Delivery guarantees**: client-generated message id for idempotency; server ack → "sent"; recipient ack → "delivered"; read event → "read". Retry until acked.
- **Offline**: store in DB, push notification; on reconnect client syncs from its last seen `message_id` per conversation.
- **Groups**: small groups → fan-out messages to each member's inbox; large channels → members pull.
- **E2E encryption** (Signal protocol): server only relays ciphertext.

## Notification System

**Requirements**: send push (APNs/FCM), SMS, email; triggered by services or scheduled; user preferences & opt-outs; rate limits; at-least-once delivery with dedupe.

```viz
services -> notification API -> validate, prefs, templates -> Kafka topics per channel
                                                               |        |        |
                                                           push wkrs  sms wkrs  email wkrs
                                                               |        |        |
                                                             APNs/FCM  Twilio   SES/SendGrid
                                     retries + DLQ <-----------+--------+--------+
                                     delivery status / tracking -> DB / analytics
```

**Deep dives**: per-channel queues isolate failures; idempotency key per notification to avoid duplicates on retry; respect quiet hours/time zones and frequency caps; template service; priority queues (OTP codes before marketing); track opens/clicks.

## Web Crawler

**Requirements**: crawl billions of pages, politeness (robots.txt, per-host rate), freshness (recrawl), dedupe content, scalable and robust to traps.

```viz
seed URLs -> URL frontier (priority queues + per-host queues)
                 |
                 v
     fetchers (async HTTP, DNS cache) -> robots.txt cache
                 |
                 v
     parser -> content dedupe (SimHash / checksum) -> store page (blob store)
                 |
                 v
     link extractor -> URL normalizer -> "seen URL?" (Bloom filter + DB) -> back to frontier
```

**Deep dives**
- **Frontier**: front queues by priority (PageRank, freshness), back queues one per host to enforce politeness delays.
- **Dedupe**: URLs via normalized form + Bloom filter; near-duplicate content via SimHash.
- **Scale**: partition by host hash so each worker owns some hosts (politeness is local); billions of pages/month ≈ thousands of pages/s.
- **Traps**: depth limits, URL length limits, detect infinite calendars/session IDs.

## Search Autocomplete (Typeahead)

**Requirements**: top-k (≈5–10) suggestions per prefix within ~100 ms, ranked by popularity, updated periodically (not necessarily real time).

```viz
query logs -> Kafka -> aggregation job (hourly/daily) -> (phrase, count)
                                                            |
                                                            v
                                         build trie, store top-k at every node
                                                            |
                         serialized shards (by prefix range) -> trie servers (in memory)
client -- debounce 100ms --> CDN/browser cache --> API --> trie servers
```

**Deep dives**
- **Trie with cached top-k per node** → lookup is O(prefix length), no subtree traversal at query time.
- Shard by first characters (balanced by traffic, not alphabet); replicate for reads.
- Cache responses aggressively (browser + CDN) — popular prefixes are very repetitive.
- Filters for offensive terms; personalization and trending via a real-time layer merged with the batch layer.

## Video Streaming (YouTube / Netflix)

**Requirements**: upload, transcode, stream at multiple qualities with adaptive bitrate, high availability; massive read bandwidth.

```viz
upload: client -> pre-signed URL -> object store (raw)
                                       | event
                                       v
              transcoding DAG (split into chunks -> encode 240p..4K, H.264/VP9/AV1 -> thumbnails)
                                       |
                                       v
              object store (segments + HLS/DASH manifests) -> CDN edges worldwide
                                       |
              metadata DB (title, owner, status) + search index

watch:  player -> API (metadata, manifest URL) -> CDN segments, switching bitrate by bandwidth
```

**Deep dives**: parallel chunked transcoding on a worker fleet (queue-driven DAG); **adaptive bitrate streaming** (2–10 s segments); CDN costs dominate → popularity-based placement (hot content at edges, long tail from origin); resumable uploads; view counts via stream aggregation (approximate, batched writes).

## Ride Sharing (Uber / Lyft)

**Requirements**: riders request rides; match with nearby available drivers in seconds; live location tracking; ETA & pricing; trip state machine.

**Estimates**: 1 M active drivers sending location every 4 s → 250k location updates/s.

```viz
driver app --loc every 4s--> location svc --> in-memory geo index (geohash/H3 cell -> drivers)
                                         \-> Kafka (history, ETA models)
rider app -> request -> matching svc -> query nearby cells -> rank by ETA -> offer to driver
                              |                                              (accept / timeout -> next)
                              v
                          trip svc (state machine: requested -> matched -> en route -> in trip -> done)
                              -> payments, notifications
```

**Deep dives**
- **Geo index**: geohash / Google S2 / Uber H3 cells; search the rider's cell + neighbours; keep it in memory (Redis GEO or custom), sharded by region/city.
- **High write rate**: location updates are ephemeral — keep only the latest in memory, stream history to Kafka.
- **Matching consistency**: lock a driver while offering (avoid double-assignment) with a short TTL.
- **Surge pricing**: supply/demand per cell computed in near-real time.

## File Sync (Dropbox / Google Drive)

**Requirements**: upload/download files, sync across devices, sharing, version history; large files; offline edits.

```viz
client (watcher + chunker) --metadata--> metadata svc -> SQL (files, versions, chunk lists, ACLs)
        |                                     |
        |--chunks (4 MB, by hash)--> block svc -> object store (content-addressed: sha256 -> chunk)
        |                                     |
        <--------- notification svc (long poll / WebSocket: "file X changed") --------
```

**Deep dives**: split files into **content-defined chunks**, upload only changed chunks (delta sync) and dedupe identical chunks globally by hash; metadata needs strong consistency (SQL, sharded by user/namespace); conflict handling = keep both copies ("conflicted copy"); resumable uploads with pre-signed URLs.

## Distributed Message Queue (Kafka)

**Requirements**: durable, ordered (per key), high-throughput pub/sub; consumers replay from offsets; horizontal scaling.

```viz
producer --(key -> partition = hash(key) % P)--> broker leaders
topic "orders":  P0 [0 1 2 3 4 ...]  leader B1, followers B2 B3
                 P1 [0 1 2 ...]      leader B2, followers B3 B1
consumer group G: C1 <- P0, C2 <- P1        (committed offsets per partition)
controller (KRaft/ZooKeeper): metadata, leader election
```

**Deep dives**: append-only segment files + OS page cache + zero-copy `sendfile` → very high throughput; replication with **ISR** (in-sync replicas), `acks=all` for durability; ordering only within a partition (choose the key carefully); retention by time/size or log compaction (keep latest per key); consumer rebalancing; exactly-once via idempotent producers + transactions.

## Payment System

**Requirements**: charge customers via PSPs (Stripe/Adyen), never double-charge, never lose money, full audit trail, reconciliation; correctness over latency.

```viz
checkout -> payment svc --(idempotency key)--> payment DB (SQL: payment, state machine)
                 |                                   |
                 v                                   v  outbox -> Kafka
         PSP (card network) <--webhook/callback--  ledger svc (double-entry, append-only)
                                                     |
                                       nightly reconciliation vs PSP settlement files
```

**Deep dives**
- **Idempotency** everywhere: client idempotency keys; PSP calls carry a unique payment id; retries are safe.
- **State machine**: `created → authorized → captured → settled` (or `failed / refunded`); transitions are transactional.
- **Double-entry ledger**: every movement = debit + credit rows summing to zero; immutable, append-only.
- **Consistency**: SQL with transactions; outbox pattern to emit events; sagas for multi-step flows (reserve inventory → charge → ship, with compensations).
- **Reconciliation**: compare internal ledger with PSP/bank reports, flag mismatches.
- Security: PCI-DSS scope minimization (tokenize cards via PSP), encryption, fraud checks.

## Leaderboard / Top-K

**Requirements**: real-time ranking of players by score; get top N, get a user's rank; millions of players.

```viz
game servers -> score update -> Redis sorted set  ZINCRBY lb:2026-11 <delta> <user>
                                                   ZREVRANGE lb 0 9 WITHSCORES   (top 10)
                                                   ZREVRANK lb <user>            (my rank)
               -> Kafka -> durable store (scores history) for rebuilds
```

**Deep dives**: Redis sorted sets (skip list) give O(log n) updates and rank queries; shard by score range or by game/region when one set gets too big; for approximate heavy hitters on huge streams use **Count-Min Sketch + min-heap**; time-windowed leaderboards = one key per window with TTL.

## Ticket Booking (Ticketmaster)

**Requirements**: browse events, hold seats, pay, no double booking; extreme spikes when popular events open.

```viz
users -> CDN (static event pages) -> virtual waiting room (queue tokens, admission rate)
                                          |
                                          v
       booking svc -> reserve seat: UPDATE seats SET status='held', hold_until=now()+10m
                                    WHERE id=? AND status='available'   (atomic, row lock)
                   -> payment -> confirm (status='booked') | expire hold -> 'available'
```

**Deep dives**: correctness via DB transactions / conditional updates (optimistic locking) or a distributed lock with TTL (Redis) per seat; **holds with expiry**; a **virtual waiting room** to admit users at a sustainable rate during spikes; cache read-mostly data (seat maps) and serve availability with short TTLs; idempotent payment confirmation.
