# More Designs

More frequently asked designs, in the same format: requirements → estimates → architecture → deep dives.

## Collaborative Editor (Google Docs)

**Requirements**: many users edit the same document concurrently, see each other's changes in ~100 ms, offline edits merge cleanly, version history, cursors/presence.

```viz
clients ==WebSocket==> collab servers (one "owner" per document, via consistent hashing on doc_id)
                             |  apply op -> transform/merge -> assign revision -> broadcast
                             v
                   op log (append-only, per doc)  --->  snapshotter (periodic snapshots) -> blob store
                   presence (cursors) in memory / Redis pub-sub
```

**Deep dives**
- **Concurrency control**: **OT** (operational transformation) — the server transforms concurrent ops against each other (`insert(5,"a")` vs `delete(3)` → shift index) and imposes a total order; or **CRDTs** (e.g. RGA, Yjs/Automerge): each character has a unique, ordered id, so replicas merge without a central server (great for offline/P2P, more metadata).
- Single writer per document (route all sessions of a doc to one server) simplifies ordering; fail over by replaying the op log from the last snapshot.
- Storage: op log + periodic snapshots; history = replay; compaction of old ops.
- Client keeps pending ops, acks with server revision, rebases local changes (that's OT client logic).

## Distributed Cache (Redis Cluster / Memcached)

**Requirements**: sub-millisecond `get/set/delete` with TTL, horizontally scalable to TBs, highly available, eviction when full.

```viz
app (client lib with cluster map) --hash(key)--> shard k (primary) --async repl--> replica(s)
                                  16384 hash slots -> nodes (Redis Cluster) / consistent hashing ring
config service / gossip: slot map, failure detection, replica promotion
```

**Deep dives**
- **Partitioning**: consistent hashing or fixed hash slots; resharding moves slots, clients get `MOVED` redirects.
- **Replication & failover**: async replicas, automatic promotion; possible loss of the last writes (cache = acceptable).
- **Eviction**: LRU/LFU approximations (sampled), TTL expiry (lazy on access + active sampling).
- **Hot keys**: client-side near-cache, key replication with random suffixes, request coalescing.
- **Consistency with DB**: cache-aside + delete-on-write; leases (Facebook's memcache paper) to prevent stale sets and thundering herds.
- Memory efficiency: slab allocator (Memcached) vs jemalloc + compact encodings (Redis).

## Proximity Service (Yelp / Nearby Places)

**Requirements**: find businesses within radius r of a location, filter/sort by distance & rating; 100 M places; read-heavy (~5k QPS reads, rare writes).

```viz
client -> LB -> search svc -> geo index (read replicas, in memory)
                                   geohash prefix "9q8yy" -> [place ids]
                           -> place details cache -> places DB (sharded by place_id)
business owners -> business svc -> places DB -> (async, nightly or CDC) -> rebuild geo index
```

**Deep dives**
- **Geohash**: interleave lat/lng bits into a base32 string; shared prefix ≈ nearby. Query the cell + its 8 neighbours (handles boundary effects); choose precision from the radius.
- **Quadtree**: recursively split cells with > K places; good for uneven density; built in memory at startup.
- **Google S2 / Uber H3**: hierarchical cells on a sphere; region coverings for arbitrary shapes.
- Index is small (100 M × ~24 B ≈ a few GB) → fits in memory; replicate for reads; writes can be eventually consistent.
- Ranking: distance + rating + popularity + personalization.

## Metrics Monitoring & Alerting

**Requirements**: collect metrics (CPU, latency, business counters) from 100k hosts, store with downsampling, dashboards, alert rules with low latency.

```viz
agents (push) / scrapers (pull) -> Kafka -> stream processors (aggregate, downsample)
                                               |
                                               v
                       time-series DB (e.g. Prometheus/M3/InfluxDB): series = metric + labels
                         | query API                     | rule evaluator (every 15 s)
                         v                               v
                     dashboards (Grafana)        alert manager -> dedupe/group -> PagerDuty, Slack, email
```

**Deep dives**
- **Data model**: series id = `name{label=value,…}`; points `(timestamp, value)`; beware **cardinality explosions** (user_id as a label!).
- **Storage**: append-only, compressed (delta-of-delta timestamps, XOR floats — Gorilla) ~1.4 bytes/point; time-partitioned blocks; retention tiers (raw 7 d, 1-min rollups 30 d, 1-h rollups 1 y).
- **Push vs pull**: pull (Prometheus) knows which targets are down; push works through firewalls/short-lived jobs.
- **Alerting**: avoid flapping (`for: 5m`), group and deduplicate, route by severity; alert on SLO burn rates.

## Distributed Job Scheduler

**Requirements**: schedule one-off and cron jobs (millions), run each **at least once** close to its time, retries, priorities, visibility of status.

```viz
API -> jobs DB (job definitions, next_run_at, status)
            |  scheduler shards (each owns a partition of jobs): poll "next_run_at <= now + 1 min"
            v
     delay/ready queue (Kafka / SQS / Redis sorted set by time)
            |
            v
     worker pool -> execute -> heartbeat / result -> update job status, compute next cron time
                                       | failure -> retry with back-off -> DLQ after N attempts
```

**Deep dives**: partition jobs by id across scheduler instances (leader election per partition) so no job is triggered twice by different schedulers; **idempotent jobs** + leases (`visibility timeout`) handle worker crashes; time-bucketed wheels / sorted sets for efficient "due soon" queries; clock skew tolerance; fairness and priority queues per tenant.

## Ad Click Aggregation

**Requirements**: count clicks per ad per minute in near real time (for billing and dashboards), ~10 B clicks/day, exact (billing!), handle late and duplicate events.

```viz
click logs -> Kafka (partition by ad_id) -> stream aggregator (Flink): tumbling 1-min windows,
                                             watermarks for late events, dedupe by click_id
                                                   |
                                                   v
                           aggregated (ad_id, minute, count) -> OLAP store (Druid/ClickHouse/Pinot)
raw events -> data lake (S3) -> nightly batch recomputation -> reconcile with streaming results
```

**Deep dives**: **exactly-once** = Kafka transactions / checkpointed state (Flink) + idempotent sinks; **event time vs processing time** with watermarks and allowed lateness; **hot ads** → pre-aggregate with key salting then merge; **Lambda vs Kappa** architecture (batch reconciliation for billing correctness); top-N ads per minute with a min-heap per window.

## Hotel Reservation (Booking / Airbnb)

**Requirements**: search availability, reserve rooms without double booking, payment, cancellation; moderate QPS but high correctness; allow some overbooking policy.

```viz
search -> search svc -> availability cache / search index (eventually consistent)
reserve -> reservation svc -> SQL: room_type_inventory(hotel, room_type, date, total, reserved)
              BEGIN; SELECT ... FOR UPDATE  (or UPDATE ... SET reserved = reserved + 1
                                             WHERE reserved < total * 1.1)   -- 10% overbooking
              INSERT reservation (idempotency key); COMMIT
           -> payment svc -> confirm / timeout releases inventory
```

**Deep dives**: inventory is per **room type per date** (not per room); concurrency via row locks, optimistic locking (version column) or conditional updates; **idempotency key** for the reserve call (users double-click); sharding by `hotel_id`; availability served from caches with short TTL while the final check hits the DB; holds with expiry for unpaid bookings.

## Stock Exchange (Order Matching)

**Requirements**: accept limit/market orders, match by **price-time priority**, ultra-low latency (µs–ms), strict ordering, fairness, auditability, market data feed.

```viz
brokers -> gateway (auth, risk checks) -> sequencer (assigns global sequence number)
                                              |
                                              v
                       matching engine (single-threaded, in memory, one per symbol)
                         order book: bids (max-heap/price levels) | asks (min-heap/price levels)
                                              |  fills / events
                                              v
                    event log (replicated, append-only) -> market data publisher, reporting, settlement
```

**Deep dives**: deterministic single-threaded matching per symbol (no locks); order book = price → FIFO queue of orders (sorted map + linked lists; O(1) cancel via order-id map); **event sourcing**: state = replay of the sequenced log → hot standby replicas replay the same log for failover; low-latency networking (kernel bypass, multicast market data); risk checks before sequencing.

## Object Storage (Amazon S3)

**Requirements**: store trillions of immutable objects (bytes to TBs) in buckets, 11 nines durability, high availability, PUT/GET/DELETE/LIST, versioning.

```viz
client -> API gateway (auth, rate limit) -> metadata service (bucket, key -> object id, version, placement)
                                                |  sharded KV / NewSQL
                                                v
                                   data service: placement -> storage nodes
      PUT: split into chunks -> erasure code (e.g. 8 data + 4 parity) -> spread across AZs/racks
      GET: read any 8 of 12 fragments -> reassemble
background: scrubbing (checksums), repair of lost fragments, garbage collection, lifecycle tiering
```

**Deep dives**: separate **metadata** (small, needs strong consistency for read-after-write) from **data** (huge, immutable, append-only files on disks); **erasure coding** vs 3× replication (1.5× vs 3× storage for similar durability); multipart upload for big objects; LIST via a sorted metadata index (prefix scans); durability math from independent failure domains.

## Email Service (Gmail)

**Requirements**: send/receive email (SMTP), store mailboxes, search, labels/threads, spam filtering, attachments; billions of users, mostly reads.

```viz
outbound: client -> web/API -> send queue -> SMTP outgoing servers -> internet (DKIM signing, retries)
inbound:  internet -> SMTP incoming -> spam/virus filtering -> mailbox svc
                                                                 | metadata (Bigtable/Cassandra: user_id, folder, msg_id)
                                                                 | bodies + attachments -> object store (dedupe by hash)
                                                                 | search index (per-user inverted index)
          real-time sync to clients via push (WebSocket / IMAP IDLE / push notifications)
```

**Deep dives**: shard everything by **user_id** (mailbox is per user, no cross-user queries); per-user search index; threading by `Message-ID`/`In-Reply-To` headers; spam = ML classifier + reputation (SPF/DKIM/DMARC); large attachments stored once and referenced; eventual consistency across devices with sync tokens.

## Live Comments (Facebook Live / Twitch Chat)

**Requirements**: viewers of a live video see new comments in real time (< 1–2 s); millions of viewers on popular streams; comments also stored.

```viz
commenter -> comment svc -> store (Cassandra, partition = video_id) + publish to pub/sub topic(video_id)
                                                                          |
                          gateway servers (hold SSE/WebSocket connections) subscribe per video they serve
                                                                          |
                                                               fan-out to their local viewers
```

**Deep dives**: two-level fan-out — the pub/sub layer fans out to **gateway servers** (hundreds), each gateway fans out to its **local connections** (thousands); gateways subscribe only to videos their viewers watch (dynamic subscription registry); for mega-streams sample/aggregate comments and rank them; SSE is enough (server → client), WebSockets if clients send often; backpressure for slow clients.

## LLM Inference Serving (ChatGPT-like)

**Requirements**: chat completions with streaming tokens, low time-to-first-token (TTFT), high throughput on expensive GPUs, multiple models, rate limits per user/tier, conversation history.

```viz
client --SSE stream--> API gateway (auth, quotas, rate limits)
                           -> router (model, region, load) -> inference cluster (GPU pods)
                                                                continuous batching scheduler
                                                                prefill -> decode loop (KV cache in GPU memory)
                           <- tokens streamed back as generated
side services: conversation store, safety filters (input/output), prompt cache, usage metering/billing, eval logging
```

**Deep dives**
- **Metrics**: TTFT (prefill-bound), inter-token latency / tokens per second (decode, memory-bandwidth-bound), cost per 1M tokens, GPU utilization.
- **Batching**: continuous (in-flight) batching; separate or chunked prefill to protect decode latency; **paged KV cache** (vLLM) to avoid fragmentation.
- **Scaling the model**: tensor parallelism within a node, pipeline parallelism across nodes; quantization (FP8/INT4); speculative decoding.
- **Prompt/prefix caching** for shared system prompts; route a conversation to the same replica to reuse its KV cache.
- **Capacity**: GPUs are scarce → priority queues per tier, load shedding, autoscaling with warm pools (model load takes minutes).

## Distributed Lock Service (ZooKeeper / Chubby / etcd)

**Requirements**: mutual exclusion across processes on different machines, leader election, safety under crashes and network partitions.

```viz
clients -> lock service cluster (3 or 5 nodes, Raft/ZAB consensus, replicated log)
   acquire: create ephemeral sequential node /locks/job-000042 (tied to client session)
            lowest sequence number holds the lock; others watch their predecessor
   release: delete node, or session expires (heartbeat lost) -> automatic release
   every grant carries a monotonically increasing FENCING TOKEN
storage service rejects writes with a token older than the last one seen
```

**Deep dives**: consensus (majority quorum) keeps the lock table consistent through node failures; **leases** with heartbeats free locks of dead clients — but a paused client (GC, network) may still think it holds the lock → **fencing tokens** make the protected resource reject stale holders; Redis single-node locks (`SET key val NX PX 30000`) are fine for efficiency locks, not for correctness (Redlock is debated); watch the predecessor only (avoid herd effect).

## Google Maps (Navigation & Routing)

**Requirements**: render the map at any zoom, find the fastest route between two points with live traffic, accurate ETA, turn-by-turn navigation with re-routing; ~1 B users, very read-heavy for tiles.

**Estimates**: the world at zoom level 21 is ~4.4 trillion 256×256 tiles (4²¹) — petabytes naive, but most tiles are ocean or empty land and compress to almost nothing; tiles are **static** → precompute and serve from a CDN. Active navigators send GPS points every few seconds (batched) → hundreds of thousands of location writes/s.

```viz
map tiles:   client --(zoom, x, y)--> CDN --miss--> tile store (S3)  <- offline tile/vector-tile pipeline
routing:     client -> routing svc -> routing tiles (road graph partitioned by region, cached in memory)
                                    -> live traffic layer (edge speeds) -> ETA model
location:    navigating clients --batched GPS--> Kafka -> map-matching -> traffic aggregation (per edge speed)
                                                                      -> ETA training data, re-route triggers
```

**Deep dives**
- **Tiles**: precomputed per zoom level, addressed by `zoom/x/y`, cached by CDN; **vector tiles** (geometry, styled on the client) are smaller and allow rotation/3-D.
- **Routing at continent scale**: plain Dijkstra/A* explores too much; production systems use **Contraction Hierarchies** or **hierarchical routing**: road graph partitioned into tiles with precomputed shortcuts between boundary nodes, so a query touches local roads near the endpoints and highways in between.
- **Live traffic**: map-match GPS traces to road segments, aggregate speeds per edge per minute (stream processing), update edge weights; precomputed shortcuts must be refreshed (customizable CH).
- **ETA** = route travel time from current speeds + ML correction (historical patterns by time of week, incidents) — see ETA Prediction in ML system design.
- **Re-routing**: the client follows the route locally; the server pushes better routes when traffic changes significantly.

## Nearby Friends

**Requirements**: users who opt in see friends within ~5 miles with their distance, updated about every 30 s; history kept for analytics; inactive friends disappear after a few minutes.

**Estimates**: 100 M DAU, 10% online concurrently → 10 M users × 1 update / 30 s ≈ **330k location updates/s**; each update must reach the user's online nearby friends.

```viz
client ==WebSocket==> WS servers --update--> location cache (Redis: user -> lat,lng,ts, TTL 10 min)
                          |                 -> location history (Cassandra) for analytics
                          +--publish--> pub/sub channel "user:<id>"  (Redis pub/sub cluster, sharded by channel)
friends' WS servers subscribe to each online friend's channel; on a message: compute distance, push if within radius
```

**Deep dives**
- **Pub/sub fan-out** instead of querying a geo index: each online user subscribes (via their WS server) to channels of their friends; channels are cheap, and the pub/sub tier is sharded by consistent hashing on the channel id.
- **Distance check** happens on the subscriber's WS server (it knows its own user's location) — no global geo index needed because the candidate set is just the friend list.
- **Scale of fan-out**: 330k updates/s × online friends each → millions of messages/s → many pub/sub nodes; it's CPU-bound, so shard and autoscale.
- **Stateful WS servers**: on deploy or failure, clients reconnect and re-subscribe; drain connections gradually.
- Extension "nearby random people": add a geohash-based channel per cell and subscribe to your cell + neighbours.

## Digital Wallet (Balance Transfers)

**Requirements**: users hold balances and transfer money to each other; **no money created or lost**, every transfer applied exactly once, full audit trail; high throughput (target ~1 M transfers/s at peak in Alex Xu's version).

```viz
transfer(A -> B, $5, idempotency key)
  -> sequencer / partition by account -> command log (replicated with Raft)
  -> state machine: validate (A has $5?) -> emit events "A -5", "B +5" -> apply to balances (in memory + snapshots)
  -> read side (CQRS): balances, statements, reconciliation
```

**Deep dives**
- **Sharding** accounts across many databases makes a transfer a **distributed transaction** when A and B live on different shards: options are 2PC (blocking), **TCC** (try: freeze A's $5, confirm/cancel on both sides) or a **saga** with compensation.
- **Event sourcing** (the reproducible design): store *commands* and resulting *events* in an append-only log; current balances are a deterministic replay of the log. Gives a perfect audit trail, easy recovery (snapshot + replay) and the ability to rebuild read models.
- **Exactly-once**: client idempotency keys, dedupe table keyed by request id inside the same transaction as the balance change.
- **Reliability**: replicate the event log with Raft across zones; balance state is derived, so any replica can rebuild it.
- **Reconciliation** with the external ledger/bank daily; double-entry bookkeeping (every transfer = debit + credit summing to zero).

## Dating App (Tinder)

**Requirements**: create a profile with preferences (age, distance, gender); get a stack of nearby candidates; swipe left/right; a **match** when both swiped right (notify both instantly); never show the same profile twice.

**Estimates**: 20 M DAU × 100 swipes/day = 2 B swipes/day ≈ **25k swipes/s** (peaks 3–5×).

```viz
feed:   client -> feed svc -> candidate generation: geo + filters (Elasticsearch geo_distance / geohash index)
                           -> drop already-swiped (per-user swiped set / Bloom filter) -> rank -> cached stack (Redis)
swipe:  client -> swipe svc -> swipes store (Cassandra, partition = swiper_id)
                           -> right swipe? atomic check of the reverse swipe -> match -> matches DB + push notification
```

**Deep dives**
- **Match consistency**: A and B may swipe right on each other at the same moment. Store the pair under one key (`min(A,B):max(A,B)`) and update it atomically (Redis Lua script or a single-partition conditional write) so exactly one of the two swipes creates the match.
- **Feed generation**: precompute a stack of ~100–200 candidates per user (batch + refresh when the user moves or the stack runs low) so swiping is instant; ranking blends distance, activity, mutual-interest likelihood.
- **Avoid repeats**: check candidates against the user's swipe history; a per-user Bloom filter keeps it cheap (a rare false positive just hides one profile).
- **Location updates** are infrequent and only update the user's geo document; privacy: show rounded distances.

## Online Judge (LeetCode)

**Requirements**: browse problems, submit code in many languages, run it against hidden tests with time/memory limits, return the verdict in seconds; contests with **10k+ simultaneous submitters** and a live leaderboard.

```viz
client -> API -> submissions DB (status: queued) -> queue per language (SQS/Kafka)
                                                    |
                                     sandboxed workers (container / gVisor / Firecracker microVM)
                                     compile -> run each test with CPU/mem/time limits, no network
                                                    |
                         verdict -> submissions DB -> client (poll or WebSocket) ; contest -> leaderboard
test cases: object store, cached on workers        leaderboard: Redis sorted set (score = solved, penalty)
```

**Deep dives**
- **Isolation is the core problem**: untrusted code must not escape, hog resources or reach the network → microVMs or gVisor containers, seccomp syscall filters, cgroups for CPU/memory, read-only filesystem, wall-clock and CPU timeouts, output size limits; destroy or recycle the sandbox after each run.
- **Scale**: queue decouples bursts from execution; pre-warmed worker pools per language (cold starts are slow); autoscale on queue depth before contests start.
- **Fairness**: per-user rate limits; run tests in parallel within a submission; stop at the first failing test.
- **Leaderboard**: sorted set keyed by contest with a composite score (problems solved, then lower penalty time); recompute only on accepted submissions.
- **Client-side execution** (what this site does with Pyodide) is great for practice — no servers, instant feedback — but a contest judge must run on the server because the client can't be trusted.

## Post Search (Facebook / Twitter Search)

**Requirements**: search posts by keywords, sorted by **recency** or **popularity (likes)**; new posts searchable within ~1 minute; billions of posts, high query volume, posts get liked constantly.

```viz
write:  posts -> Kafka -> indexing workers: tokenize -> append post_id to each term's posting list
                                             (sharded by term; one list ordered by time, one by like bucket)
likes:  like events -> counter svc -> re-index a post only when its count crosses a bucket (1, 10, 100, 1k …)
query:  "taylor swift" -> fetch both terms' lists -> intersect -> top N -> fetch exact like counts -> re-rank
        popular queries cached (TTL ~ 1 min)
```

**Deep dives**
- **Inverted index** (term → post ids) partitioned by term; keep postings sorted by time so "most recent" is a prefix read; very hot terms are cached and replicated.
- **Like-count churn**: re-indexing on every like would be a write storm — index coarse popularity buckets and fetch exact counts only for the top candidates at query time.
- **Freshness vs cost**: an in-memory real-time index for recent posts plus a larger on-disk index for older ones; queries merge both.
- **Scale**: billions of posts × ~10 terms each — compress posting lists (delta + varint encoding); move cold index segments to cheaper storage.
- **Privacy filtering**: the result set must respect visibility (friends-only posts) → filter at query time against the viewer's graph, or index visibility attributes.

## Online Auction (eBay)

**Requirements**: list items with an end time; users place bids; a bid must beat the current highest; everyone watching sees the price update live; the highest bid at the end wins; bids must never be lost.

```viz
bid -> API (idempotency key) -> bids partitioned by auction_id (Kafka) -> bid processor (one per partition)
                                   -> conditional update: UPDATE auctions SET price=?, leader=? WHERE id=? AND price < ?
                                   -> bids log (all bids, durable) -> publish "new price" -> SSE to watchers
auction end: delayed job at end_time -> close auction -> notify winner -> payment
```

**Deep dives**
- **Consistency of the highest bid**: either an atomic conditional update in the database (optimistic, retry on conflict) or serialize bids per auction through a partitioned queue (one consumer per auction → no races, natural ordering).
- **Durability**: write every bid to a durable log *before* acknowledging; the auction's state can be rebuilt from it.
- **Real-time updates**: SSE/WebSockets to watchers via pub/sub per auction; popular auctions have many watchers → two-level fan-out (pub/sub → gateway servers).
- **Ending correctly**: use the server's timestamp for bid acceptance; close via a scheduled job (delay queue / timer wheel); many sites **extend the end time** after late bids to discourage sniping.
- **Proxy bidding**: users give a maximum; the system bids the minimum increment on their behalf.

## Food Delivery (DoorDash / Uber Eats)

**Requirements**: browse restaurants near the customer with menus and ETAs; place and pay for orders; restaurants accept and prepare; **dispatch** assigns couriers; live courier tracking; order status updates.

```viz
browse:  client -> search svc (geo index of restaurants + filters) -> menu cache/CDN -> ETA estimates
order:   client -> order svc (state machine: placed -> accepted -> preparing -> picked up -> delivered)
                -> payment (authorize; capture on delivery) -> restaurant tablet notification
dispatch: every few seconds per region: open orders × available couriers -> optimizer -> offers to couriers
couriers: location every ~5 s -> location svc (in-memory geo index by region) -> tracking via WebSocket/SSE
```

**Deep dives**
- **Dispatch is an assignment problem**: batch orders and couriers per region every few seconds and solve for minimum total delivery time (min-cost matching / heuristics), considering food prep time (don't make couriers wait), order batching (two orders, one courier) and fairness.
- **ETA** combines prep time prediction, courier travel time and traffic (an ML model per stage).
- **Order state machine** with idempotent transitions; restaurants and couriers confirm steps; timeouts trigger reassignment.
- **Peak hours**: surge pricing / delivery fees by region to balance supply and demand; precompute restaurant availability.
- Location pipeline, geo index and real-time tracking mirror the Ride Sharing design.

## Top-K Trending (YouTube Top Videos / Heavy Hitters)

**Requirements**: the top K most viewed videos for the last hour, day, month and all time; views arrive at very high rate; results can be a minute stale.

**Estimates**: 70 B views/day ≈ **800k views/s** — far too many to update a database row per view.

```viz
view events -> Kafka (partition by video_id) -> stream counters (Flink): count per video per 1-minute bucket
                                                  -> per-partition top-K (min-heap) -> merge -> global top-K
                                                  -> window rollups: last hour = 60 minute-buckets, day = 24 hour-buckets
precomputed lists (per window) -> cache -> API "top 100 today"
```

**Deep dives**
- **Exact counting** works because the stream is partitioned by video id: each partition owns its videos' counts, keeps a local top-K, and a final stage merges the per-partition lists (the global top K must appear in some partition's top K).
- **Sliding windows**: keep per-minute counts and add the newest bucket / subtract the expiring one; long windows use coarser buckets (hours, days).
- **Approximate** alternative when memory is tight: **Count-Min Sketch + heap** per window — bounded memory, small overestimates.
- **Hot keys**: a viral video concentrates load on one partition → pre-aggregate in the producers (count locally for a second, then send) or salt the key and merge.
- All-time counts are updated in batches to the video database (e.g. every minute), never per view.
