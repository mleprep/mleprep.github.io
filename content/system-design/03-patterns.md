# Common Patterns

Most designs are a few recurring problems in disguise. Recognize the pattern, then reach for its standard toolbox — and say the trade-offs out loud.

## Scaling Reads

**When**: read-heavy systems (feeds, product pages, profiles, URL redirects) — typically 10:1 to 1000:1 reads to writes.

Work down this ladder, cheapest first:

```viz
1. optimize the query      indexes (composite, covering), avoid N+1, denormalize hot joins
2. cache                   CDN for static/public, Redis/Memcached for app data, in-process for tiny hot sets
3. read replicas           scale reads horizontally; accept replication lag
4. precompute              materialized views, fan-out-on-write timelines, search indexes
5. shard                   when the working set no longer fits one primary + replicas
```

- **Cache invalidation**: cache-aside with delete-on-write + TTL; versioned keys (`user:42:v7`) avoid races.
- **Replica lag**: route a user's reads to the primary for a few seconds after they write (read-your-writes), or carry a version/LSN token.
- **Hot keys** (a viral post): replicate the key across shards (`key#1..#N`, read a random copy), in-process caches with short TTL, request coalescing so one miss fetches for many waiters.
- **Thundering herd** on expiry: jittered TTLs, probabilistic early refresh, a lock so only one request recomputes.

## Scaling Writes

**When**: high ingest (metrics, clicks, location updates, logs, likes).

- **Shard / partition** by a key with high cardinality and even load; avoid monotonically increasing keys on range-partitioned stores (all writes hit the last shard) — hash them or prefix with a bucket.
- **Batch & buffer**: accumulate in memory or a queue and write in bulk; Kafka absorbs bursts and decouples producers from the database.
- **Write-optimized storage**: LSM-tree stores (Cassandra, RocksDB, ScyllaDB) turn random writes into sequential ones; time-series DBs compress heavily.
- **Pre-aggregate**: don't store every like as a row you count later — increment counters (sharded counters for hot ones: `likes:post42:{0..15}`, sum on read) or aggregate in a stream processor.
- **Drop what you don't need**: ephemeral data (driver location every 4 s) can live only in memory with the latest value; sample high-volume analytics.
- **Backpressure**: when consumers fall behind, queues grow — monitor lag, autoscale consumers, shed low-priority work.

## Dealing with Contention

**When**: many users compete for the same thing — the last concert seat, an item in a flash sale, a username, a bid, a bank balance.

| Technique | How | Good for |
|---|---|---|
| **Atomic conditional update** | `UPDATE seats SET holder=? WHERE id=? AND holder IS NULL` (check rowcount) | single-row races; simplest correct answer |
| **Pessimistic lock** | `SELECT … FOR UPDATE` inside a transaction | short critical sections on one DB |
| **Optimistic concurrency** | version column; `UPDATE … WHERE version = v`; retry on conflict | low contention, many readers |
| **Distributed lock / lease** | Redis `SET key val NX PX ttl`, ZooKeeper/etcd with fencing tokens | coordinating across services; holds with expiry |
| **Serialize through a queue** | partition by resource id, one consumer per partition applies requests in order | extreme contention (auctions, order books) |
| **Reservation / hold** | temporary hold with expiry, confirm on payment | multi-step checkout (tickets, hotels) |

Discuss: what happens when the lock holder crashes (TTL/lease), double-submits (idempotency key), and whether a **virtual waiting room** should limit how many users even reach the contended step.

## Multi-Step Processes (Workflows & Sagas)

**When**: one user action spans several services or slow steps — order → payment → inventory → shipping; a ride's lifecycle; onboarding flows.

- **State machine** per entity (`created → paid → shipped`), persisted, with legal transitions only.
- **Saga**: a sequence of local transactions, each with a **compensating action** (refund, release inventory) if a later step fails. *Orchestration* (a coordinator tells each service what to do — easier to reason about) vs *choreography* (services react to each other's events — looser coupling, harder to trace).
- **Durable execution / workflow engines** (Temporal, AWS Step Functions, Cadence): write the workflow as code; the engine persists every step, retries with back-off, survives crashes and supports timers ("cancel the hold after 10 minutes").
- **Outbox pattern** to publish events atomically with the state change; **idempotent** steps so retries are safe.
- Avoid distributed 2PC across services: it blocks on coordinator failure and couples availability.

## Long-Running Tasks

**When**: work that takes seconds to hours — video transcoding, report generation, ML training, bulk imports, sending a million emails.

```viz
client --POST /jobs--> API --enqueue--> queue (SQS / Kafka / Redis) --> worker pool --> result store (S3/DB)
   |       202 Accepted + job_id                                      | heartbeat, progress %
   +--GET /jobs/{id} (poll) or webhook / WebSocket push when done <---+
```

- Return **202 + job id** immediately; never hold an HTTP request open for minutes.
- Workers take a **lease** (visibility timeout) and heartbeat; if they die the job reappears. Make jobs **idempotent** (or checkpoint progress) because at-least-once delivery means repeats.
- **Retries** with exponential back-off, then a **dead-letter queue** for inspection.
- Split huge jobs into chunks (fan-out / fan-in) so they parallelize and partial failures retry cheaply.
- Separate queues/priorities per job type or tenant so one heavy customer can't starve others; autoscale workers on queue depth.

## Handling Large Files (Blobs)

**When**: uploads/downloads of images, video, documents, backups — anything bigger than a few MB.

- **Never stream large files through your app servers.** The API issues a **pre-signed URL**; the client uploads directly to object storage (S3/GCS) and downloads come from a **CDN**.
- **Multipart / resumable upload**: split into chunks (5–100 MB), upload in parallel, retry only failed parts, then "complete" the upload.
- Keep **metadata** (owner, size, checksum, status, storage key) in a database; the blob store holds bytes. Status goes `pending → uploaded → processed` via the object-store event notification.
- **Content-addressed chunks** (key = SHA-256 of the chunk) give dedupe and delta sync (Dropbox).
- Post-processing (thumbnails, transcoding, virus scanning) runs asynchronously off the upload event.
- Security: short-lived signed URLs scoped to one key and method; validate type/size after upload.
