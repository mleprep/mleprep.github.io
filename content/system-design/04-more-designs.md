# More Designs

Thirteen more frequently asked designs, in the same format: requirements → estimates → architecture → deep dives.

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
