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
