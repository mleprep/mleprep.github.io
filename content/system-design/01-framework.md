# Framework & Estimation

How to run a 45-minute system design interview, and the numbers you should know by heart.

## The 4-Step Framework

```viz
 1. Requirements (5 min)      2. Estimates + API (5 min)
    functional / non-func.  ->   QPS, storage, bandwidth; endpoints
            |                                   |
            v                                   v
 4. Deep dives (15-20 min) <-  3. High-level design (10 min)
    bottlenecks, scaling,        boxes & arrows, data model,
    failure modes, trade-offs    happy path end to end
```

**1. Clarify requirements** — never start drawing first.
- *Functional*: the 3–5 core features ("users can post, follow, see a feed"). Explicitly park the rest.
- *Non-functional*: scale (DAU, read/write ratio), latency (p99), availability vs consistency, durability, geo-distribution.

**2. Back-of-the-envelope + API**
- QPS = DAU × actions/day ÷ 86,400 (≈ 10⁵ s/day). Peak ≈ 2–10× average.
- Storage = writes/day × object size × retention.
- Define 3–5 API calls (`POST /tweets`, `GET /feed?cursor=`).

**3. High-level design**: client → LB → stateless services → cache → DB / object store / queue. Walk one request through it.

**4. Deep dives**: pick the 2–3 hardest parts (fan-out, hot keys, consistency, sharding key, failure handling). Always state **trade-offs** and alternatives.

**Signals interviewers look for:** driving the conversation, quantitative reasoning, knowing *why* a component is there, discussing failure modes, and adapting when requirements change.

## Numbers Everyone Should Know

| Operation | Latency |
|---|---|
| L1 cache reference | 1 ns |
| Main memory reference | 100 ns |
| Compress 1 KB (fast codec) | 2 µs |
| Read 1 MB sequentially from memory | 3 µs |
| SSD random read | 16–100 µs |
| Round trip within a datacenter | 0.5 ms |
| Read 1 MB sequentially from SSD | 50 µs – 1 ms |
| Disk seek (HDD) | 2–10 ms |
| Round trip CA ↔ Europe | 150 ms |

| Power of 2 | ≈ | Name |
|---|---|---|
| 2¹⁰ | 1 thousand | KB |
| 2²⁰ | 1 million | MB |
| 2³⁰ | 1 billion | GB |
| 2⁴⁰ | 1 trillion | TB |
| 2⁵⁰ | 1 quadrillion | PB |

**Rules of thumb**
- 1 day ≈ 86,400 s ≈ 10⁵ s; 1 M requests/day ≈ 12 QPS.
- A single modern server: ~10k–100k simple QPS (cache), ~1k–10k QPS for DB-backed reads.
- Redis: ~100k ops/s per node; Postgres/MySQL: a few k writes/s per primary.
- Availability: 99.9% ≈ 8.8 h/yr downtime; 99.99% ≈ 53 min/yr; 99.999% ≈ 5 min/yr.

## Worked Estimation Example

> "Design Twitter: 300 M monthly users, 150 M daily, each posts 2 tweets and reads the feed 20 times a day."

```viz
writes : 150M x 2  / 1e5 s  ≈   3,000 tweets/s     (peak ~ 10k/s)
reads  : 150M x 20 / 1e5 s  ≈  30,000 feed reads/s (peak ~100k/s)   -> read-heavy, cache!
storage: 300M tweets/day x ~300 B ≈ 90 GB/day text  ≈ 33 TB/yr  (+ media in object store)
media  : 10% tweets with 1 MB image -> 30 TB/day   -> CDN + blob storage
```

Conclusions you should *say out loud*: read:write = 10:1 → precompute feeds and cache aggressively; media dominates storage → object store + CDN; text fits in a sharded key-value / wide-column store.

## Requirements Checklist

Ask about these before designing:

- **Users & scale**: DAU/MAU, growth, geographic distribution.
- **Traffic shape**: read/write ratio, peak vs average, bursty events (celebrity posts, flash sales).
- **Data**: size per object, retention, hot vs cold, relationships (graph? relational?).
- **Latency**: interactive (<200 ms) vs async/batch.
- **Consistency**: is stale data OK (feeds, likes) or not (payments, inventory)?
- **Availability / durability**: can we lose data? multi-region?
- **Security & privacy**: auth, PII, abuse/rate limiting.
- **Out of scope**: say it explicitly (analytics, admin tools, ML ranking …).
