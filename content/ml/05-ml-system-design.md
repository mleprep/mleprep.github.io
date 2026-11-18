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
