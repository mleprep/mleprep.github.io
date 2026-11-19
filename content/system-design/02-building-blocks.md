# Building Blocks

The components every design is assembled from, and the trade-offs to talk about for each.

## Scaling: Vertical vs Horizontal

- **Vertical** (bigger machine): simple, no code changes; hard ceiling, single point of failure, expensive at the top end.
- **Horizontal** (more machines): near-linear scale and redundancy; requires **stateless** services (keep sessions in Redis/JWT, files in object storage) and a load balancer.

```viz
              +--> [app 1] --+
client -> LB -+--> [app 2] --+--> cache --> DB primary --> replicas
              +--> [app 3] --+
   stateless app tier: any request can hit any server
```

Scale the **stateless tier** first (easy), then the **data tier** (hard: replication → caching → sharding).

## Load Balancing

- **L4** (TCP/UDP, fast, no content awareness) vs **L7** (HTTP: route by path/header/cookie, TLS termination, retries).
- Algorithms: round-robin, weighted, **least connections**, latency-based, **consistent hashing** (sticky by key for caches).
- Health checks remove dead nodes; use **active-passive or active-active** LB pairs (or DNS / anycast) so the LB isn't a SPOF.
- Global: **GeoDNS / anycast** routes users to the nearest region.

## Caching

Where: client → CDN → API gateway → **application cache (Redis/Memcached)** → DB buffer pool.

| Strategy | How | Trade-off |
|---|---|---|
| **Cache-aside** (lazy) | app reads cache, on miss reads DB and fills | simple, most common; first read slow, possible stale |
| **Read-through** | cache library loads from DB on miss | cleaner app code |
| **Write-through** | write cache + DB synchronously | consistent; slower writes |
| **Write-back** | write cache, flush to DB async | fast writes; risk of loss |
| **Write-around** | write DB only, cache fills on read | avoids polluting cache |

- **Eviction**: LRU (default), LFU, TTL.
- **Pitfalls**: *thundering herd / cache stampede* (use request coalescing, locks, or probabilistic early refresh), *hot keys* (replicate the key, local in-process cache), *invalidation* (delete-on-write + TTL as a safety net), *cache penetration* (cache negative results, Bloom filter).

```viz
read(k):  v = cache.get(k)
          if v is None: v = db.get(k); cache.set(k, v, ttl)
write(k): db.put(k, v); cache.delete(k)      # delete, don't update -> avoids races
```

## CDN

- Edge servers cache static assets (images, video segments, JS) close to users → lower latency and origin load.
- **Pull** (fetch from origin on first miss; simple) vs **push** (upload proactively; good for large, predictable content).
- Control freshness with TTL / `Cache-Control`; bust with versioned URLs (`app.3f9c.js`).
- Video: chunked HLS/DASH segments are perfect for CDN caching.

## SQL vs NoSQL

| | Relational (Postgres, MySQL) | NoSQL |
|---|---|---|
| Model | tables, joins, schema | key-value (Redis, DynamoDB), document (Mongo), wide-column (Cassandra, Bigtable), graph (Neo4j) |
| Transactions | ACID, multi-row | usually single-key/partition; some offer more |
| Scaling | vertical + read replicas; sharding is manual | built-in partitioning, horizontal |
| Best for | money, inventory, relationships, ad-hoc queries | huge scale, simple access patterns, flexible schema, high write throughput |

**Rule of thumb**: start relational unless you have a clear reason (massive write volume, simple key access, schema flexibility, time-series). Model NoSQL tables **around queries**, not entities.

## Replication

- **Leader–follower**: writes to leader, reads from followers. Async = fast but **replication lag** (read-your-writes issues; route a user's reads to the leader for a few seconds after they write). Sync = durable but slower.
- **Multi-leader**: writes in several regions; must resolve conflicts (last-write-wins, CRDTs, app logic).
- **Leaderless** (Dynamo/Cassandra): write to W of N replicas, read from R; **R + W > N** gives overlap → read sees latest write. Repair with read-repair and anti-entropy (Merkle trees).

```viz
N=3, W=2, R=2:   write -> [A ✓] [B ✓] [C ✗]
                 read  <- [B]   [C]        B has the new version -> overlap guaranteed
```

Failover: detect leader death (heartbeats), elect a new leader (Raft/Paxos/ZooKeeper), fence the old one to avoid **split brain**.

## Sharding (Partitioning)

Split data across nodes by a **shard key**.

- **Range** (`A–F`, `G–M`…): range scans are efficient; risk of hot ranges (e.g. timestamps).
- **Hash** (`hash(key) % N` or consistent hashing): even spread; range queries scatter.
- **Directory / lookup service**: flexible, extra hop.

Choosing a key: high cardinality, even access, and matching the dominant query (e.g. `user_id` for per-user data). Watch out for:
- **Hot partitions** (celebrity users) → add a random suffix / split the key, cache.
- **Cross-shard queries & transactions** → denormalize, scatter-gather, or 2PC/sagas.
- **Resharding** → consistent hashing or many logical shards mapped to fewer physical nodes.

## Consistent Hashing

Place nodes and keys on a hash **ring**; a key belongs to the first node clockwise. Adding/removing a node moves only about `1/N` of the keys (with `hash % N`, almost all keys would move). **Virtual nodes** (each server appears at many ring positions) smooth out load and heterogeneity.

```viz
            0
       N3 .   . N1          key k1 -> walk clockwise -> N1
        .        .          add N4 between N1 and N2:
     k3 .         . k1        only keys in (N1, N4] move to N4
        .        .
       N2 .   . (N4)
           k2
```

Used by: Dynamo/Cassandra, Memcached clients, CDNs, load balancers for sticky routing.

## CAP and PACELC

- **CAP**: during a network **P**artition you must choose **C**onsistency (reject/block some requests) or **A**vailability (serve possibly stale data). Without partitions you can have both.
- **PACELC**: *if P → A or C; Else → Latency or Consistency*. E.g. Cassandra/Dynamo = PA/EL, Spanner = PC/EC.
- Practical framing: which data needs **strong consistency** (balances, inventory, unique usernames) vs **eventual** (likes, feeds, view counts)?

## Consistency Models

From strongest to weakest:
- **Linearizable**: every read sees the latest completed write (single-copy illusion). Needs consensus/leader; costs latency.
- **Sequential / causal**: everyone sees causally related writes in order (replies after the post).
- **Read-your-writes**, **monotonic reads**: per-session guarantees; often enough for UX.
- **Eventual**: replicas converge if writes stop.

Tools: consensus (**Raft**, Paxos) for leader election & replicated logs; **quorums**; **vector clocks**/version vectors to detect conflicts; **CRDTs** for merge-without-coordination (counters, sets).

## Message Queues & Event Streams

Decouple producers and consumers, absorb spikes, enable async work and fan-out.

| | Queue (SQS, RabbitMQ) | Log / stream (Kafka, Pulsar, Kinesis) |
|---|---|---|
| Consumption | message deleted after ack; competing consumers | append-only log; consumers track offsets; replayable |
| Ordering | best-effort / FIFO queues | ordered **per partition** |
| Use | task queues, jobs | event sourcing, analytics, CDC, fan-out to many systems |

Delivery semantics: **at-most-once**, **at-least-once** (default; make consumers **idempotent**), **exactly-once** (idempotent producer + transactional writes, or dedupe by message id). Add **dead-letter queues** and exponential back-off retries.

```viz
producers -> [ topic: partition 0 | partition 1 | partition 2 ] -> consumer group A (one consumer per partition)
                                                                -> consumer group B (independent offsets)
```

## Rate Limiting Algorithms

| Algorithm | Idea | Pros / cons |
|---|---|---|
| **Token bucket** | bucket refills at rate r up to capacity b; each request takes a token | allows bursts up to b; tiny state (tokens, last_refill) — most common |
| **Leaky bucket** | queue drains at fixed rate | smooth output; bursts get queued/dropped |
| **Fixed window counter** | count per minute window | simple; 2× burst at window edges |
| **Sliding window log** | store timestamps, count last 60 s | exact; memory heavy |
| **Sliding window counter** | weighted current + previous window | good approximation, cheap |

```python
def allow(key, rate, capacity, now):
    tokens, last = store.get(key, (capacity, now))
    tokens = min(capacity, tokens + (now - last) * rate)   # refill
    if tokens < 1:
        return False
    store[key] = (tokens - 1, now)
    return True
```

Distributed: keep counters in Redis with an atomic Lua script (or `INCR` + `EXPIRE`); return **HTTP 429** with `Retry-After`.

## Database Indexes: B-Tree vs LSM-Tree

- **B-tree** (Postgres, MySQL InnoDB): pages updated in place; great reads and range scans; random writes cost more.
- **LSM-tree** (Cassandra, RocksDB, LevelDB, HBase): writes go to an in-memory **memtable** + WAL, flushed as immutable sorted **SSTables**, merged by **compaction**. Excellent write throughput; reads may check several SSTables (mitigated by **Bloom filters** and sparse indexes).

```viz
write -> WAL (disk, sequential) + memtable (RAM)
                     | full
                     v
          SSTable L0  SSTable L0      -- compaction -->  L1 (bigger, sorted, non-overlapping) -> L2 ...
read  -> memtable -> bloom filter? -> SSTables newest to oldest
```

Other indexes: **hash** (point lookups), **inverted index** (full-text search: term → posting list), **geospatial** (geohash, quadtree, S2), **composite** (order matters: `(user_id, created_at)` serves "user's latest posts").

## Unique ID Generation

| Approach | Pros | Cons |
|---|---|---|
| DB auto-increment | simple, ordered | single point of failure / bottleneck |
| UUID v4 (128-bit random) | no coordination | large, unordered → poor index locality |
| Ticket server / ranges | simple, ordered within range | extra service |
| **Snowflake** (64-bit) | time-sortable, no coordination, compact | needs clock sync; worker id assignment |

```viz
Snowflake 64-bit id
| 1 sign | 41 bits: ms since epoch (~69 yrs) | 10 bits: machine id (1024) | 12 bits: sequence (4096/ms) |
```

Newer: **UUIDv7 / ULID** — time-ordered 128-bit IDs with random tail.

## Idempotency, Retries & Timeouts

- Networks fail mid-request → clients **retry** → servers must make operations **idempotent**.
- **Idempotency key**: client sends a unique key (`Idempotency-Key: <uuid>`); server stores key → result and returns the saved result on repeats. Essential for payments.
- Retries: **exponential back-off + jitter**, cap the number, and only retry idempotent/safe operations.
- **Timeouts** everywhere; **circuit breakers** stop hammering a failing dependency; **bulkheads** isolate resources.
- Distributed transactions: avoid 2PC across services; use **sagas** (sequence of local transactions with compensating actions) and the **outbox pattern** (write the event to an outbox table in the same DB transaction, relay it to the queue).

## Observability & Reliability

- **Metrics** (RED: rate, errors, duration; USE: utilization, saturation, errors), **logs** (structured, with request IDs), **traces** (distributed tracing across services).
- **SLI/SLO/SLA**: e.g. SLO "99.9% of feed requests < 300 ms over 30 days"; error budget drives release pace.
- Reliability patterns: redundancy across AZs/regions, graceful degradation (serve cached/partial results), load shedding, autoscaling, chaos testing, backups + tested restores.

## Real-Time Updates: Polling, SSE, WebSockets

How the server pushes fresh data (chat, live comments, notifications, dashboards, collaborative editing, driver locations).

| Transport | How | Pros | Cons / use when |
|---|---|---|---|
| **Short polling** | client asks every N s | trivial, works everywhere, stateless | wasted requests, latency ≤ N s; fine for slow-changing data |
| **Long polling** | server holds the request until data arrives (or ~30 s), client re-requests | near real time over plain HTTP | a request per message; reconnect storms |
| **Server-Sent Events** | one long-lived HTTP response streaming `text/event-stream` | simple, server → client, auto-reconnect with `Last-Event-ID`, works over HTTP/2 | one direction only; good for feeds, live comments, LLM token streaming |
| **WebSockets** | upgraded TCP connection, full duplex | low latency both ways | stateful connections: sticky routing, LB/proxy support, reconnection logic; chat, games, collaboration |
| **WebRTC** | peer-to-peer (STUN/TURN for NAT traversal) | lowest latency, offloads servers | complex; audio/video calls |

The server side is the real design problem:

```viz
publishers -> pub/sub (Redis pub/sub, Kafka, NATS) -> gateway servers holding connections -> clients
                                                     ^
                     connection registry: user_id -> gateway (Redis, with TTL heartbeats)
```

- Connections are **stateful**: track which gateway holds each user; gateways subscribe only to the topics their users need.
- Plan for **reconnects**: clients resume from the last event id / sequence number and fetch what they missed from durable storage.
- One gateway holds tens to hundreds of thousands of mostly idle connections (memory and file descriptors per connection); scale horizontally and drain connections gracefully on deploys.

## API Design

- **REST** (resources + HTTP verbs, JSON): the default for public APIs; cacheable GETs. **gRPC** (Protobuf over HTTP/2): typed, fast, streaming — great between internal services. **GraphQL**: the client picks fields from one endpoint — avoids over/under-fetching for rich UIs; harder caching and rate limiting (needs query-cost analysis).
- **Idempotency**: GET/PUT/DELETE are idempotent by definition; make POSTs safe to retry with an `Idempotency-Key` header.
- **Pagination**: offset/limit is simple but slow and unstable for deep pages on changing data; **cursor (keyset)** pagination — `?after=<last_id>` with `WHERE id < :cursor ORDER BY id DESC LIMIT 20` — is stable and O(page size).
- **Versioning** (`/v1/…` or a header), a consistent error format with machine-readable codes, `429 + Retry-After` for rate limits, `202 Accepted + job id` for long-running work.
- Authenticate every call (OAuth 2 bearer tokens / API keys), authorize per resource, limit request sizes.
- In interviews, define 3–5 endpoints with their key parameters — enough to anchor the data model, not a full spec.

## Probabilistic Data Structures

Trade a small, bounded error for huge memory savings — ideal for "have we seen it?", "how many?", "how often?" at scale.

| Structure | Answers | Error | Size |
|---|---|---|---|
| **Bloom filter** | "is x in the set?" | false positives only (never false negatives) | ~9.6 bits per element for 1% false positives |
| **Count-Min Sketch** | "how often did x occur?" | overestimates only | d × w counters, independent of the number of items |
| **HyperLogLog** | "how many distinct items?" | ~1.04/√m relative standard error | Redis: 12 KB → 0.81% error for billions of items |

```viz
Bloom filter (m bits, k hash functions):   add x  -> set bits h1(x) .. hk(x)
                                           query  -> all k bits set ? "maybe" : "definitely not"
false-positive rate ≈ (1 − e^(−kn/m))^k,   best k = (m/n)·ln 2
```

- **Bloom**: skip disk reads for absent keys (LSM SSTables in Cassandra/Bigtable/RocksDB), "already crawled this URL?", "username probably taken?". No deletes (use a counting Bloom filter or a cuckoo filter).
- **Count-Min**: d hash rows of w counters; an update increments one counter per row; the estimate is the minimum over the rows. Trending / heavy hitters = CMS + a min-heap of the current top k.
- **HyperLogLog**: hash each item, keep per bucket the longest run of leading zeros seen, combine buckets with a harmonic mean. Daily active users, unique visitors; sketches are mergeable (union across shards or days for free).

## Search & Inverted Indexes

**Inverted index**: term → posting list of document ids (+ positions, frequencies). Queries intersect/union posting lists; ranking with **BM25** (term frequency with saturation × inverse document frequency, normalized by document length), optionally followed by ML re-ranking.

```viz
doc 1 "red fox"   doc 2 "red panda"   ->   red: [1, 2]   fox: [1]   panda: [2]
query "red fox"  ->  [1, 2] ∩ [1]  ->  doc 1
```

- Text analysis: tokenize, lowercase, stem/lemmatize, stop words, synonyms, edge n-grams for prefix search.
- **Elasticsearch / OpenSearch**: an index is split into **shards** (each a Lucene index) with replicas; new documents become searchable on **refresh** (~1 s, near real time); immutable segments are merged in the background.
- Keep the **source of truth in the primary database** and feed the index through **CDC** (e.g. Debezium → Kafka → indexer) so it can always be rebuilt; search is eventually consistent.
- Partitioning: by document (every query fans out to all shards, gathers top-k — the default) vs by term (fewer shards per query, harder writes). Paginate with `search_after` cursors, not deep `from/size`.
- Not a primary database: no multi-document transactions, eventual consistency, reindex to change mappings.

## Batch vs Stream Processing

| | Batch (MapReduce, Spark, warehouse SQL) | Stream (Flink, Kafka Streams, Spark Structured Streaming) |
|---|---|---|
| Input | bounded data at rest (S3, HDFS, tables) | unbounded event streams (Kafka) |
| Latency | minutes to hours | milliseconds to seconds |
| Strengths | easy reprocessing, big joins, cheap | fresh results, alerting, real-time features |

Streaming concepts interviewers probe:
- **Windows**: tumbling (fixed, non-overlapping), hopping/sliding (overlapping), session (closed after a gap of inactivity).
- **Event time vs processing time**: aggregate by when the event happened; **watermarks** estimate "all events up to time T have arrived"; **allowed lateness** updates results for stragglers; very late events go to a side output.
- **State & exactly-once**: operators keep state (counts, joins) that is periodically **checkpointed**; after a failure, restore the checkpoint and replay Kafka from the saved offsets; idempotent or transactional sinks complete end-to-end exactly-once.
- **Lambda** = batch layer (correct) + speed layer (fresh), merged at query time — two codebases. **Kappa** = streaming only; reprocess history by replaying the log.

## Geospatial Indexing

"What's near me?" — map 2-D points to something an ordinary index can search.

| Index | Idea | Notes |
|---|---|---|
| **Geohash** | interleave latitude/longitude bits → base-32 string; shared prefix ≈ nearby | 5 chars ≈ 4.9 × 4.9 km, 6 ≈ 1.2 × 0.6 km, 7 ≈ 153 × 153 m; query the cell + its 8 neighbours (edge effects); works in any B-tree/KV store and Redis `GEOSEARCH` |
| **Quadtree** | split a cell into 4 once it holds more than K points | adapts to density (cities vs oceans); usually an in-memory structure |
| **S2 / H3** | hierarchical cells on the sphere (S2: Hilbert curve on a cube; H3: hexagons) | region coverings; H3's equidistant neighbours suit ride-hailing and surge pricing |
| **R-tree** (PostGIS GiST) | nested bounding rectangles | polygons and "within shape" queries |

Rules of thumb: static places (Yelp) → precompute an in-memory index and replicate it; moving objects (drivers, friends) → keep only the latest position in memory (Redis GEO, sharded by region) and stream the history elsewhere.

## Distributed Transactions

Changing data owned by different databases/services atomically is hard — first try to avoid it (keep data that must change together in one database).

- **Two-phase commit (2PC)**: the coordinator asks every participant to *prepare* (lock + vote), then tells all to *commit* or *abort*. Atomic but **blocking**: if the coordinator dies after prepare, participants hold locks until it recovers; as slow as the slowest participant. Used inside databases (Spanner, XA), rarely across microservices.
- **Saga**: a chain of local transactions with **compensating actions** for rollback (eventually consistent; intermediate states are visible). Orchestrated (central coordinator) or choreographed (events).
- **TCC (Try–Confirm–Cancel)**: each service reserves in *try* (e.g. freeze $50), then *confirm* or *cancel* — explicit holds instead of locks; common in payments and wallets.
- **Transactional outbox**: write the business row and an event row in one local transaction; a relay publishes the event (at least once) → never "DB committed but message lost".
- Make every step **idempotent** and record processed message ids — retries and duplicate deliveries will happen.

## Consensus & Leader Election

**Consensus** lets a group of nodes agree on a value or a sequence of commands despite failures — the basis of leader election, replicated logs, distributed locks and configuration stores (etcd, ZooKeeper, Consul).

- **Raft** in one breath: nodes are followers, candidates or the leader; a follower that hears no heartbeat within a **randomized timeout** starts an election for a new **term** and wins with a **majority** of votes; the leader appends client commands to its log and replicates them; an entry is **committed** once stored on a majority and is then applied in order.
- **2f + 1 nodes tolerate f failures** (3 → 1, 5 → 2). Every write needs a majority round trip, so keep clusters small and use them for **metadata and coordination**, not bulk data.
- **Split brain**: an old leader cut off by a partition may still believe it leads → **leases** plus **fencing tokens** (monotonic epoch numbers checked by the storage layer).
- Paxos is the classic equivalent; ZooKeeper uses ZAB; etcd, Consul, CockroachDB, TiKV and Kafka's KRaft use Raft.

## Microservices, API Gateway & Service Discovery

- **Monolith first** for small teams; split along business domains when independent deployment, scaling or team ownership matter. Costs: network calls, partial failures, distributed data, harder debugging.
- **API gateway**: single entry point for authentication, rate limiting, routing, TLS termination and request aggregation (backend-for-frontend).
- **Service discovery**: instances register themselves (Consul, etcd, Kubernetes Services/DNS); health checks remove bad instances; clients or load balancers resolve names to healthy instances.
- **Service mesh** (Envoy sidecars with Istio/Linkerd): mTLS, retries, timeouts, circuit breaking and telemetry without application changes.
- **Database per service** keeps services decoupled; share data through APIs or events, not by reading each other's tables.
- **Resilience**: timeouts on every call, retries with jitter (idempotent operations only), circuit breakers, bulkheads, fallbacks.

## Authentication & Authorization

- **AuthN** (who are you): passwords hashed with bcrypt/scrypt/Argon2 (never a fast hash), MFA, SSO via **OIDC** (OpenID Connect, built on OAuth 2.0).
- **Sessions vs tokens**: server-side sessions (opaque id in a cookie, state in Redis) are easy to revoke; **JWTs** (signed, self-contained claims) avoid a lookup per request but are hard to revoke → short expiry plus refresh tokens.
- **OAuth 2.0** grants: authorization code + PKCE for user-facing apps, client credentials for service-to-service; **scopes** limit what a token can do.
- **AuthZ** (what may you do): RBAC (roles), ABAC (attribute policies) or relationship-based access (Google Zanzibar style: "user is an editor of the doc through a group") for sharing-heavy products — enforced server side on every request.
- TLS everywhere (mTLS between services), encryption at rest, secrets in a vault/KMS (never in code or repos), least privilege, audit logs.

## Choosing a Database

| Need | Typical choice | Why |
|---|---|---|
| Transactions, relations, flexible queries | **PostgreSQL / MySQL** (+ read replicas) | ACID, joins, mature; scales further than most people think |
| Huge-scale key-value, predictable latency | **DynamoDB / Cassandra / ScyllaDB** | partitioned and horizontally scalable; model tables around queries |
| Cache, counters, leaderboards, rate limits | **Redis** | in-memory data structures (sorted sets, hashes, streams) |
| Full-text and faceted search | **Elasticsearch / OpenSearch** | inverted index, relevance ranking |
| Analytics over billions of rows | **ClickHouse / BigQuery / Snowflake / Druid / Pinot** | columnar storage, fast aggregations |
| Metrics and time series | **Prometheus / InfluxDB / TimescaleDB** | compression, downsampling, time-range queries |
| Files, images, video, backups | **S3 / GCS** object storage + CDN | cheap, durable, effectively unlimited |
| Highly connected data, multi-hop queries | **Neo4j** (or relational tables with care) | graph traversals |
| Global strong consistency | **Spanner / CockroachDB** | distributed SQL on consensus replication |
| Embeddings / semantic search | **pgvector, Elasticsearch kNN, Milvus, Pinecone** | ANN indexes (HNSW, IVF-PQ) |

Justify the choice by **access pattern, consistency needs and scale** — not by brand.
