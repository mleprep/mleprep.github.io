# ML System Design

End-to-end ML design questions (recommendations, ranking, fraud, …) and the production concerns around them.

## ML System Design Framework

```viz
1. Clarify      business goal -> ML objective -> metric (online + offline); constraints (latency, scale, privacy)
2. Data         sources, labels (explicit/implicit, delayed?), sampling, leakage, privacy
3. Features     user / item / context / cross features; embeddings; feature store; freshness
4. Model        baseline (heuristic, logistic regression) -> GBDT -> deep model; multi-stage if large catalog
5. Training     splits (time-based), loss, imbalance, offline eval, hyperparameters
6. Serving      batch vs online, latency budget, caching, candidate generation -> ranking -> re-ranking
7. Evaluation   offline metrics -> A/B test on business metrics; guardrails
8. Monitoring   data/prediction drift, feedback loops, retraining cadence, fallbacks
```

Always start with a **simple baseline** and explain what you'd improve next. State trade-offs (latency vs accuracy, freshness vs cost, exploration vs exploitation).

## Recommendation System (YouTube / Netflix / TikTok)

**Goal**: maximize long-term engagement (e.g. watch time, satisfied sessions), not just clicks.

```viz
billions of items
   |  CANDIDATE GENERATION (recall, ~ms): two-tower ANN retrieval, co-visitation, trending,
   |                                       subscriptions, ...   -> ~1000 candidates
   v
   RANKING (precision): deep model with rich features, multi-task heads
   |                    (p(click), p(watch>30s), p(like), p(share)) -> weighted score -> ~100
   v
   RE-RANKING / POLICY: diversity, freshness, dedupe, business rules, fairness -> final ~20
   v
   logging -> training data (impressions + outcomes) -> retrain (daily / streaming)
```

- **Two-tower model**: user tower (history, context) and item tower (content, stats) map into the same embedding space; trained with in-batch negatives / sampled softmax; item embeddings indexed in an ANN store (HNSW / ScaNN).
- **Ranking features**: user–item interactions, item popularity & freshness, creator affinity, context (time, device), cross features.
- **Cold start**: content-based features, popularity priors, exploration (bandits) for new items; onboarding questions for new users.
- **Pitfalls**: position bias (log position, use it as a feature during training and fix it at serving), feedback loops / popularity bias, clickbait (optimize satisfaction signals), delayed labels.
- **Metrics**: offline recall@k, NDCG, AUC per head; online watch time, retention, diversity, DAU.

## Search Ranking

**Goal**: return the most relevant documents for a query in ~100–200 ms.

```viz
query -> understanding (spell fix, tokenization, intent, query expansion, embeddings)
      -> retrieval: inverted index (BM25) + dense ANN  -> ~thousands
      -> L1 ranker (cheap: GBDT / small model)         -> ~hundreds
      -> L2 ranker (cross-encoder / deep LTR)          -> top results
      -> blending (ads, freshness, diversity) -> SERP -> clicks logged
```

- **Learning to rank**: pointwise (regression/classification), pairwise (RankNet, LambdaMART on GBDTs), listwise (LambdaRank/ListNet) — optimize NDCG.
- **Labels**: human relevance ratings (expensive, high quality) + click logs (abundant, biased: position bias, presentation bias → counterfactual / click models, interleaving experiments).
- **Features**: text match (BM25, embedding similarity), document quality/authority, freshness, personalization, query–doc historical CTR.

## Ads Click-Through Rate Prediction

**Goal**: predict `p(click | user, ad, context)` accurately and **calibrated**, because it feeds the auction: `rank score = bid × pCTR (× quality)`.

- Massive sparse categorical features (user id, ad id, publisher, hashed crosses) → **embeddings** + feature interactions: logistic regression with hashed crosses (classic), Factorization Machines, Wide & Deep, DeepFM, DCN, DLRM.
- **Training**: billions of examples, online/continual learning (hourly or streaming) because ads and trends change fast; negative downsampling → re-calibrate (`p' = p / (p + (1 − p)/w)`).
- **Metrics**: log loss and calibration (expected/observed CTR ≈ 1) offline; revenue, CTR, advertiser ROI online.
- Latency: score hundreds of ads in ~10 ms → model distillation, caching embeddings, feature store lookups.

## Fraud Detection

**Goal**: flag fraudulent transactions in real time (~50 ms) with high recall at a tolerable false-positive rate.

```viz
transaction -> feature service (real-time aggregates: #tx last 1h per card/device/IP, velocity,
               geo distance from last tx, merchant risk) + static features
           -> rules engine (hard blocks, allow lists)
           -> model (GBDT / GNN on entity graph) -> score -> approve / step-up auth / block / manual review
labels: chargebacks (arrive weeks later) + analyst reviews -> retraining
```

- **Extreme imbalance** (0.1%): PR-AUC, recall at fixed FPR, cost-weighted metrics (fraud $ caught vs friction).
- **Delayed & noisy labels**, **adversarial drift** (fraudsters adapt) → frequent retraining, monitoring, anomaly detectors for novel patterns.
- **Graph features**: shared devices/cards/addresses link fraud rings.
- Explainability for analysts and regulators (SHAP, reason codes).

## Harmful Content Moderation

**Goal**: detect policy-violating posts (hate, violence, spam, nudity) across text, images and video at platform scale.

- **Multimodal** classifiers (text encoder + image/video encoder, fused) with multi-label heads per policy.
- **Pipeline**: hash matching for known bad content (PhotoDNA-style perceptual hashes) → fast classifiers on every post → expensive models / LLMs on borderline cases → human review queue prioritized by severity × reach.
- Thresholds per policy and region; precision-focused auto-removal, recall-focused human review.
- Active learning from reviewer decisions; adversarial robustness (text obfuscation, image perturbations); fairness audits across languages and groups.
- Metrics: prevalence (views of violating content per 10k views), proactive detection rate, appeal overturn rate.

## Feature Stores & Training–Serving Skew

**Q: What is training–serving skew and how do you prevent it?**

Skew = the model sees different feature values in production than in training (different code paths, stale data, time-travel leakage).
- **Feature store**: one feature definition, computed by batch/stream pipelines, written to an **offline store** (historical, for training with **point-in-time correct joins**) and an **online store** (low-latency KV for serving). Examples: Feast, Tecton.
- **Log features at serving time** and train on the logged values when possible.
- Monitor feature distributions online vs training (PSI / KL), null rates, freshness.

## A/B Testing & Online Experimentation

**Q: How do you know your new model is better?**

- Randomize **units** (users, not requests, to avoid inconsistent experiences); choose primary metric + **guardrails** (latency, crashes, revenue).
- **Power analysis** up front: sample size `n ≈ 16σ²/δ²` per arm for 80% power at α = 0.05 (δ = minimum detectable effect).
- Run for full weekly cycles; avoid **peeking** (or use sequential testing); correct for multiple comparisons.
- **Variance reduction**: CUPED (use pre-experiment covariates).
- Pitfalls: novelty/primacy effects, network effects/interference (use cluster or switchback randomization), sample ratio mismatch (SRM — check it!), Simpson's paradox.
- Faster alternatives for ranking: **interleaving**; for exploration: **multi-armed bandits**.

## Model Monitoring & Drift

**Q: What can go wrong after deployment, and how do you catch it?**

- **Data drift** (input distribution changes: P(x)), **concept drift** (relationship changes: P(y|x)), **label drift** (P(y)).
- Monitor: input feature stats, prediction distribution, calibration, latency/errors, business KPIs; and model quality once labels arrive (possibly delayed).
- Responses: alerts → investigate (upstream pipeline bug is the most common cause!) → retrain on recent data (scheduled or triggered) → shadow/canary deploy new models → automatic rollback.
- **Feedback loops**: the model influences its own future training data (recommendations, fraud blocking) → keep exploration traffic / holdouts to collect unbiased labels.


## People You May Know (Friend Recommendation)

**Goal**: suggest connections a user is likely to accept, increasing healthy network growth.

```viz
candidates: friends-of-friends (2-hop), same school/company/contacts upload, co-engagement
     -> features: mutual friends count, Adamic-Adar, graph embeddings (node2vec/GNN), profile similarity,
                  recency of interactions, geography
     -> ranker (GBDT / two-tower + MLP) predicting p(request sent) and p(accept)
     -> re-rank: diversity, spam/fake-account filters, privacy rules, de-dupe shown recently
```

Labels: invitations sent/accepted (impressions without action = negatives, careful with position bias). Graph computations (2-hop counts) are heavy → precompute in batch daily, refresh hot users incrementally. Metrics: accept rate, connections per DAU, long-term retention; guardrails against spammy invites.

## Visual Search (Search by Image)

**Goal**: given a photo, return visually similar items (Pinterest Lens, Google Lens, product search).

```viz
offline: catalog images -> embedding model (CNN/ViT, trained contrastively on (query, clicked item) pairs)
                        -> vectors -> ANN index (HNSW / IVF-PQ), sharded + replicated
online:  query image -> object detection / crop -> embed (GPU) -> ANN top-k -> re-rank
         (metadata filters, freshness, personalization) -> results
```

Key points: train embeddings with **contrastive / triplet loss** using engagement data; hard-negative mining; quantization (PQ) to fit billions of vectors in memory; index refresh for new items; evaluate with recall@k on labeled pairs + online CTR; dedupe near-identical images.

## ETA Prediction (Maps / Food Delivery)

**Goal**: predict arrival time accurately (MAE / percentile errors) at request time, < 50 ms.

- **Features**: route segments with real-time speeds (from GPS probes), historical speeds by time-of-week, weather, events, road type; for delivery: restaurant prep time, courier state.
- **Model**: route-level baseline = sum of segment times (graph routing) + ML correction (GBDT or GNN over road graph); quantile regression for P90 promises.
- **Training data**: completed trips (label = actual duration); watch for distribution shift (holidays, construction).
- **Serving**: precompute segment speeds every minute in a streaming job; cache popular routes.
- Metrics: MAE, MAPE, % within ±2 min, calibration of P90 (late rate ≈ 10%).

## LLM-Powered Support Assistant

**Goal**: answer customer questions from help-center docs and account data, escalate to humans when needed; reduce resolution time without hurting satisfaction.

```viz
user message -> safety + intent classifier -> retrieval (help docs, past tickets) + account tools (order status API)
             -> LLM with system prompt, retrieved context, tool results -> answer with citations
             -> confidence / policy check -> respond | hand off to human agent (with summary)
feedback: thumbs, resolution, CSAT, agent edits -> eval set -> prompt / retrieval / fine-tune iterations
```

Design points: **RAG** over curated docs (freshness pipeline when docs change); **tools** for personal data with strict auth scoping; guardrails (no unsupported promises, PII redaction, prompt-injection defenses); offline eval set of real conversations graded by rubric + LLM-as-judge validated by humans; online metrics: containment rate, CSAT, escalation quality, hallucination reports; cost and latency budgets (smaller model for routing, bigger for hard cases).

## News Feed Ranking (Facebook / LinkedIn)

**Goal**: order the posts from a user's network (and recommended content) to maximize meaningful engagement and long-term retention, not just clicks.

```viz
candidates: posts from friends/follows/groups since last visit (+ recommended, ads slots)  ~ thousands
     -> light ranker (small model / GBDT)                                                  ~ hundreds
     -> heavy ranker: multi-task DNN -> p(click), p(like), p(comment), p(share), p(hide), p(dwell > 10s)
     -> value model: score = Σ w_i · p_i   (weights set by product goals; negative for hide/report)
     -> re-rank: diversity (not 5 posts from one author), freshness, integrity demotions, ads blending
```

- **Features**: viewer–author affinity (past interactions, closeness), post features (type, age, early engagement velocity), viewer context (device, time, session depth), content embeddings.
- **Labels**: implicit engagement events with weights; negative feedback (hide, unfollow, report) matters as much as positives; delayed signals (did the user come back tomorrow?) via long-term value models.
- **Training**: logged impressions, time-based splits, frequent retraining (daily or streaming); handle **position bias** and the feedback loop (exploration slots, counterfactual evaluation).
- **Serving**: the feed is computed on request (fresh) with a latency budget of a few hundred ms; heavy features precomputed in a feature store; cache candidates.
- **Metrics**: offline AUC/NDCG per head and calibration; online time spent, meaningful interactions, retention, negative feedback rate, creator-side health.

## Similar Listings (Airbnb / Item-to-Item Recommendations)

**Goal**: on a listing page, show other listings the user is likely to book next ("similar homes"), for a large catalog with constant churn.

```viz
training: user browsing sessions  [l12, l7, l40, l3 (booked)]  -> skip-gram on sessions (listing2vec)
          positives = co-viewed within a window; global context = the booked listing; negatives = same-market samples
serving:  listing embedding -> ANN index per market -> top 100 nearby in embedding space
          -> filters (dates available, capacity, location radius) -> light re-rank (price, quality) -> top 12
```

- **Session-based embeddings** (word2vec on click sessions) capture what users actually consider similar — better than raw attribute similarity. Airbnb's trick: add the eventually-booked listing as global context and sample negatives from the **same market** so embeddings learn within-city distinctions.
- **Cold start** for new listings: average the embeddings of nearby listings with similar type/price/capacity until data accrues.
- **Freshness**: retrain embeddings daily; index refresh; availability changes constantly → filter at query time, not in the embedding.
- **Metrics**: offline — rank of the eventually booked listing among candidates (hit rate@k); online — CTR on the carousel, bookings, revenue.

## Street View Blurring (Object Detection System)

**Goal**: automatically blur faces and license plates in billions of street-level images, with **very high recall** (privacy) and acceptable precision (don't blur the whole image).

```viz
images (batch, petabytes) -> preprocessing (resize, tiles) -> detector (faces, plates) -> NMS
                          -> blur regions -> publish; sampled images -> human review -> new labels -> retrain
```

- **Model**: a one-stage or two-stage detector with a feature pyramid (small, distant faces/plates); tiles for high-resolution panoramas; classes: face, license plate.
- **Data**: human-annotated boxes, augmentations (blur, lighting, occlusion, weather), hard-negative mining (posters, statues), active learning on low-confidence cases.
- **Thresholds**: tune for recall ≥ ~99% on a held-out set, accept extra false positives; per-class thresholds.
- **Offline batch inference**: no latency constraint but huge throughput → GPU fleet, batching, quantized models; reprocess when the model improves.
- **Metrics**: recall/precision per class at IoU ≥ 0.5, mAP; user reports of unblurred faces as an online signal.

## Gmail Smart Compose (Text Autocomplete)

**Goal**: suggest the rest of the sentence as the user types an email; suggestions must be fast (< ~100 ms per keystroke), useful, and safe.

```viz
context: subject + previous email (if reply) + text typed so far + (opt.) locale, time
   -> tokenizer -> small transformer LM (decoder, or encoder for context + decoder)
   -> beam search, few tokens, stop at phrase boundary -> confidence threshold -> show grey suggestion
user presses Tab = accept (positive) / keeps typing = reject (negative)
```

- **Model**: a compact language model trained on (de-identified) email text; latency dominates → small model, quantization, caching per-keystroke state (KV cache / incremental decoding), run on TPU/GPU servers or on device.
- **Decoding**: beam search for short, high-probability completions; only show when confidence is high (bad suggestions annoy more than none).
- **Personalization**: user-level adaptation (e.g. a lightweight personalized n-gram/adapter blended with the global model).
- **Safety & privacy**: never memorize/leak private data (differential privacy, de-duplication, filtering rare sequences), filter offensive or biased suggestions (e.g. gendered pronoun assumptions).
- **Metrics**: offline perplexity, exact-match@k on held-out emails; online acceptance rate, characters saved per email, latency p99.

## Text-to-Image Generation Service

**Goal**: generate high-quality images from text prompts at scale, safely, with a few seconds of latency.

```viz
prompt -> safety filter (blocked concepts) -> text encoder (CLIP/T5) -> latent diffusion model (DiT / U-Net)
       -> 20-50 denoising steps with classifier-free guidance -> VAE decoder -> (super-resolution)
       -> output safety classifier (NSFW, likeness) -> watermark -> image
```

- **Model**: latent diffusion — train a VAE to compress images, then a diffusion transformer to denoise latents conditioned on text via cross-attention; trained on large filtered, recaptioned image–text data (better captions → better prompt following).
- **Serving cost**: each image = dozens of network evaluations → step-reduction (DPM-Solver, distillation, consistency models), batching, fp16/fp8, smaller resolution + super-resolution, caching prompt embeddings; GPU autoscaling with queues and per-user quotas.
- **Controls**: guidance scale, seed, negative prompts, image-to-image, inpainting, ControlNet-style conditioning (pose, edges).
- **Safety**: prompt classifiers, output classifiers, block real-person likeness/IP misuse, invisible watermarking (C2PA provenance metadata).
- **Evaluation**: FID / CLIP score offline, human preference win rates (side by side), prompt-adherence benchmarks; online: regeneration rate, downloads.
