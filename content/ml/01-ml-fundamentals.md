# ML Fundamentals

The classical machine learning questions that come up in almost every ML interview, with crisp answers.

## Bias–Variance Trade-off

**Q: What are bias and variance, and how do they relate to under/overfitting?**

Expected test error decomposes as **bias² + variance + irreducible noise**.
- **Bias**: error from wrong assumptions (model too simple) → **underfitting**: high train *and* test error.
- **Variance**: sensitivity to the particular training sample (model too flexible) → **overfitting**: low train error, high test error.

```viz
error
  |\                         /  total
  | \  bias²              _/
  |  \__              __/
  |     \____  ______/  variance
  |          \/
  +-----------------------------> model complexity
        under   sweet    over
        fit     spot     fit
```

| Symptom | Fix |
|---|---|
| High bias | bigger model, more features, less regularization, train longer |
| High variance | more data, regularization, dropout, early stopping, simpler model, ensembling (bagging) |

Note: very large neural nets can show **double descent** — test error falls again past the interpolation threshold.

## Regularization (L1, L2, Dropout, Early Stopping)

**Q: How does regularization reduce overfitting? L1 vs L2?**

Add a penalty that discourages complex solutions: `loss = data_loss + λ·R(w)`.
- **L2 (ridge / weight decay)**: `R = Σw²` → shrinks all weights smoothly; equivalent to a Gaussian prior; handles correlated features well.
- **L1 (lasso)**: `R = Σ|w|` → drives weights **exactly to zero** (sparse, feature selection); Laplace prior. The L1 ball has corners on the axes, where the loss contours tend to touch it.
- **Elastic net**: mix of both.
- **Dropout**: randomly zero activations during training (scale by `1/(1−p)`); trains an implicit ensemble and prevents co-adaptation.
- **Early stopping**: stop when validation loss stops improving.
- **Data augmentation**, **label smoothing**, **noise injection** also regularize.

```viz
      w2                         w2
       |   L2 ball (circle)       |   L1 ball (diamond)
     .-+-.     contours           /\        contours touch at
    (  +  )  touch anywhere     /  +\       a corner -> w1 = 0
     '-+-'  -> small w1, w2     \  + /      (sparse)
       |                          \/
```

## Train / Validation / Test, Cross-Validation & Leakage

**Q: How do you set up evaluation so the numbers are trustworthy?**

- **Train** to fit, **validation** to tune hyperparameters/early stopping, **test** touched once for the final estimate.
- **k-fold CV** for small data: average over k train/validate splits. **Stratify** for imbalanced labels; **group k-fold** when samples are correlated (same user); **time-based splits** for temporal data (never train on the future).
- **Data leakage** = information unavailable at prediction time sneaks into training: target-derived features, normalizing using statistics of the full dataset, duplicate users across splits, future data. Symptom: offline metrics that are "too good" and collapse in production.

## Evaluation Metrics

**Q: Precision vs recall, ROC-AUC vs PR-AUC — when to use which?**

```viz
                  predicted +    predicted -
actual +             TP              FN
actual -             FP              TN

precision = TP / (TP + FP)      "of what I flagged, how much was right"
recall    = TP / (TP + FN)      "of all positives, how many did I catch"   (= TPR, sensitivity)
FPR       = FP / (FP + TN)
F1        = 2PR / (P + R)
```

- **Precision** matters when false positives are costly (spam filter hiding real mail); **recall** when misses are costly (cancer screening, fraud).
- **ROC-AUC**: TPR vs FPR over thresholds; = probability a random positive scores above a random negative. Insensitive to class balance — can look great on very imbalanced data.
- **PR-AUC / average precision**: better for rare positives.
- **Log loss** & **calibration** (reliability diagrams, Platt/isotonic scaling) when probabilities are used downstream (ads bidding).
- Ranking: **NDCG@k**, **MAP**, **MRR**, recall@k.
- Regression: **MAE** (robust to outliers), **RMSE** (penalizes large errors), **R²**, MAPE.
- Always tie the offline metric to the **business metric** and verify with A/B tests.

## Handling Class Imbalance

**Q: Your positive class is 0.1% of the data. What do you do?**

1. Use the right metric (PR-AUC, recall@precision, not accuracy).
2. **Resample**: undersample negatives (fast; re-calibrate probabilities afterwards) or oversample positives / SMOTE.
3. **Class weights** in the loss (`weight ∝ 1/frequency`), or **focal loss** (down-weights easy examples).
4. Tune the **decision threshold** on validation data for the required precision/recall.
5. Get more positives: active learning, weak labels, data augmentation.
6. Consider anomaly-detection framing when positives are extremely rare.

## Gradient Descent & Optimizers

**Q: Explain SGD, momentum, and Adam. How do you choose the learning rate?**

- **Batch GD**: exact gradient, expensive. **SGD / mini-batch**: noisy but cheap, the noise helps escape sharp minima; batch size 32–4096.
- **Momentum**: `v = βv + g; w -= lr·v` → accelerates along consistent directions, damps oscillation.
- **RMSProp**: divides by a running RMS of gradients → per-parameter step sizes.
- **Adam**: momentum (1st moment `m`) + RMSProp (2nd moment `v`) with bias correction:
  `m = β1·m + (1−β1)g`, `v = β2·v + (1−β2)g²`, `w -= lr · m̂ / (√v̂ + ε)`. **AdamW** decouples weight decay (the default for transformers).
- **Learning rate**: the most important hyperparameter. Use **warmup** then **cosine/linear decay**; LR range test; too high → divergence/loss spikes, too low → slow, stuck.

```viz
loss surface (narrow valley)
SGD:       /\/\/\/\/\/\--> zig-zags across the valley
Momentum:  ~~~~~~~~~~~~--> builds speed along the valley floor
```

## Linear & Logistic Regression

**Q: Derive the loss for logistic regression. Why not MSE?**

- **Linear regression**: `ŷ = Xw + b`, minimize MSE; closed form `w = (XᵀX)⁻¹Xᵀy` (use GD or regularize when `XᵀX` is ill-conditioned). Assumes linearity, independent errors, homoscedasticity.
- **Logistic regression**: `p = σ(w·x + b)`, `σ(z) = 1/(1+e^{−z})`. Maximizing the Bernoulli likelihood ⇒ minimize **binary cross-entropy**:
  `L = −[y log p + (1−y) log(1−p)]`, with gradient `∂L/∂w = (p − y)·x` (same elegant form as linear regression).
- Why not MSE with a sigmoid: the loss becomes **non-convex** and gradients vanish when the prediction is confidently wrong.
- Coefficients are interpretable as **log-odds**; multi-class → softmax regression.

## Decision Trees, Random Forests & Gradient Boosting

**Q: Bagging vs boosting? Why does XGBoost win on tabular data?**

- **Decision tree**: greedy splits maximizing impurity reduction (**Gini** `1−Σp²` or **entropy** `−Σp log p`; variance for regression). Interpretable, handles mixed types, but high variance.
- **Random forest (bagging)**: many deep trees on bootstrap samples + random feature subsets per split → averaging **reduces variance**. Robust, little tuning, out-of-bag error estimate.
- **Gradient boosting (GBDT: XGBoost, LightGBM, CatBoost)**: add shallow trees **sequentially**, each fitting the **negative gradient (residuals)** of the loss → **reduces bias**. Needs care: learning rate (shrinkage), number of trees, depth, subsampling, early stopping.

```viz
bagging:   data -> [tree][tree][tree] (parallel, independent) -> average
boosting:  data -> tree1 -> residuals -> tree2 -> residuals -> tree3 ... -> sum (sequential)
```

GBDTs dominate tabular data: handle heterogeneous features, missing values, non-linear interactions, monotonic constraints, fast training; little preprocessing (no scaling needed).

## K-Means & Clustering

**Q: How does k-means work and what are its limitations?**

Alternate **assign** (each point → nearest centroid) and **update** (centroid = mean of its points) until assignments stop changing. Minimizes within-cluster sum of squares; converges to a local optimum.
- Init with **k-means++** (spread-out seeds); run several restarts.
- Choose k with the **elbow** method, **silhouette** score, or domain needs.
- Limitations: assumes spherical, similarly sized clusters; sensitive to scale (standardize) and outliers.
- Alternatives: **GMM** (soft assignments, elliptical clusters, EM), **DBSCAN** (density-based, arbitrary shapes, finds outliers), hierarchical clustering.

## PCA & Dimensionality Reduction

**Q: What does PCA do? How is it computed?**

Find orthogonal directions of **maximum variance**: eigenvectors of the covariance matrix (or via **SVD** of the centered data `X = UΣVᵀ`; components = rows of `Vᵀ`, explained variance ∝ `σ²`). Project onto the top-k to compress, denoise or visualize.
- Center (and usually standardize) features first.
- Linear only; for visualization of non-linear structure use **t-SNE** / **UMAP**; for learned compression use autoencoders.
- Use cases: speeding up models, removing multicollinearity, visualization, embedding compression.

## Feature Engineering & Missing Data

**Q: How do you prepare features for a model?**

- **Numerical**: scaling (standardize for linear models/NNs; not needed for trees), log-transform skewed data, clip outliers, bucketize.
- **Categorical**: one-hot (low cardinality), **target encoding** (with out-of-fold to avoid leakage), **hashing trick**, **learned embeddings** (high cardinality: user/item ids).
- **Text**: TF-IDF, pretrained embeddings. **Time**: cyclical encodings (sin/cos of hour), lags, rolling stats.
- **Crosses / interactions** for linear models.
- **Missing values**: drop, impute (mean/median/mode, model-based), add a "was missing" indicator (missingness is often informative); trees handle missing natively in XGBoost/LightGBM.

## Generative vs Discriminative Models

**Q: Compare Naive Bayes and logistic regression.**

- **Discriminative** models learn `p(y|x)` directly (logistic regression, SVM, NNs): usually better accuracy with enough data.
- **Generative** models learn `p(x, y) = p(x|y)p(y)` (Naive Bayes, GMM, HMM, VAEs, diffusion): can generate data, handle missing features, and work well with little data.
- **Naive Bayes** assumes features are conditionally independent given the class → super fast, strong baseline for text; probabilities are poorly calibrated.
- Ng & Jordan: Naive Bayes reaches its (higher) asymptotic error faster; logistic regression wins with more data.

## SVM & Kernels

**Q: What is a support vector machine? What is the kernel trick?**

- Find the hyperplane that **maximizes the margin** between classes; only the points on the margin (**support vectors**) define it. **Soft margin** (hinge loss + C) tolerates misclassification.
- **Kernel trick**: replace dot products with `K(x, x')` (RBF, polynomial) to separate data in an implicit high-dimensional space without computing it.
- Good for medium-sized, high-dimensional data; scales poorly (O(n²–n³)) to millions of samples.

## KNN and the Curse of Dimensionality

**Q: When does nearest-neighbour search break down?**

KNN predicts from the k closest training points (no training, expensive inference). In high dimensions, distances concentrate — **all points become almost equally far apart**, and the data needed to cover the space grows exponentially. Remedies: dimensionality reduction, learned embeddings, feature selection. At scale, use **approximate nearest neighbour** indexes (HNSW, IVF-PQ, ScaNN, FAISS).


## MLE vs MAP and Bayes' Theorem

**Q: What's the difference between maximum likelihood and MAP estimation?**

- **Bayes**: `p(θ | D) = p(D | θ) p(θ) / p(D)` — posterior ∝ likelihood × prior.
- **MLE**: `argmax p(D | θ)`; e.g. MSE regression = MLE under Gaussian noise, cross-entropy = MLE for categorical labels.
- **MAP**: `argmax p(D | θ) p(θ)` = MLE + regularizer: Gaussian prior ⇒ **L2**, Laplace prior ⇒ **L1**.
- Full Bayesian inference keeps the whole posterior (uncertainty estimates) instead of a point estimate.

Classic worked example (base-rate fallacy): a test is 99% sensitive and 95% specific for a disease with 1% prevalence. `P(sick | +) = 0.99·0.01 / (0.99·0.01 + 0.05·0.99) ≈ 16.7%` — most positives are false positives when the condition is rare.

## Hyperparameter Tuning

**Q: How do you tune hyperparameters efficiently?**

- **Grid search** wastes budget on unimportant dimensions; **random search** covers important ones better (Bergstra & Bengio).
- **Bayesian optimization** (Gaussian process / TPE surrogate + acquisition function) for expensive trainings; **Hyperband / ASHA** stop bad trials early.
- Tune the most sensitive first: learning rate (log scale!), batch size, regularization strength, model size.
- Always tune on a validation set (or CV), never on the test set; fix seeds and report variance across seeds.

## Ensembles & Stacking

**Q: Why do ensembles work, and what is stacking?**

Averaging models with **uncorrelated errors** reduces variance (`Var(mean) = ρσ² + (1−ρ)σ²/M`). Techniques:
- **Bagging** (bootstrap + average: random forest), **boosting** (sequential bias reduction: GBDT), **snapshot/seed ensembles** for neural nets.
- **Stacking**: train a meta-model on **out-of-fold** predictions of base models (never in-fold → leakage).
- Costs: latency and complexity at serving time → often **distill** the ensemble into one model.

## Statistics: Hypothesis Tests, p-values & Confidence Intervals

**Q: What does a p-value mean? How do you test whether a change helped?**

- **Null hypothesis H₀** (no effect) vs alternative. The **p-value** is the probability of seeing data at least this extreme **if H₀ were true** — not the probability that H₀ is true, and not the size of the effect.
- Reject H₀ when p < α (often 0.05). **Type I error** (false positive) rate = α; **Type II error** (false negative) rate = β; **power** = 1 − β (typically 0.8).
- **Central limit theorem**: means of many independent samples are approximately normal with standard error `σ/√n` — why z/t-tests work on averages even when the data isn't normal.
- **95% confidence interval** for a mean: `x̄ ± 1.96·σ/√n`. Interpretation: the *procedure* captures the true mean 95% of the time.
- Tests: z/t-test for means (Welch's t-test for unequal variances), two-proportion z-test for conversion rates, χ² for categorical counts, Mann–Whitney for skewed data, **bootstrap** when no formula fits (resample with replacement, take percentiles).
- **Multiple comparisons**: testing 20 metrics at α = 0.05 expects one false positive → Bonferroni (α/m) or Benjamini–Hochberg (controls the false discovery rate).
- Statistical vs practical significance: with huge samples, tiny useless effects become "significant" — report effect sizes and intervals.

## Information Theory: Entropy, Cross-Entropy, KL

**Q: How are entropy, cross-entropy and KL divergence related, and why do we minimize cross-entropy?**

- **Entropy** `H(p) = −Σ p log p`: average surprise / optimal code length; maximal for the uniform distribution.
- **Cross-entropy** `H(p, q) = −Σ p log q`: cost of coding data from p with a code built for q.
- **KL divergence** `KL(p‖q) = Σ p log(p/q) = H(p, q) − H(p) ≥ 0`, zero iff p = q; **not symmetric**.
- Training with cross-entropy against fixed labels p minimizes `KL(p‖q_model)` (H(p) is constant) = maximum likelihood.
- Forward KL `KL(p‖q)` is *mass-covering* (q spreads to cover all of p's modes); reverse KL `KL(q‖p)` is *mode-seeking* (used in variational inference; also the KL penalty in RLHF).
- **Mutual information** `I(X;Y) = H(X) − H(X|Y)`: how much knowing Y reduces uncertainty about X — feature selection, InfoNCE is a lower bound on it.
- Perplexity of a language model = `exp(cross-entropy per token)`.

## Loss Functions Cheat Sheet

**Q: Which loss would you use, and why?**

| Task | Loss | Notes |
|---|---|---|
| Regression | **MSE** | Gaussian noise MLE; punishes outliers heavily |
| Regression, outliers | **MAE**, **Huber** | MAE → median; Huber = quadratic near 0, linear far away |
| Quantiles / P90 ETAs | **pinball (quantile) loss** | asymmetric weights τ and 1 − τ |
| Binary classification | **binary cross-entropy** (log loss) | calibrated probabilities |
| Multi-class | **softmax cross-entropy** | + label smoothing to reduce over-confidence |
| Multi-label | sigmoid BCE per label | labels are independent |
| Imbalanced / dense detection | **focal loss** `−(1−p)^γ log p` | down-weights easy examples (RetinaNet) |
| Max-margin | **hinge** `max(0, 1 − y·f(x))` | SVMs |
| Ranking | pairwise (RankNet, BPR), listwise (LambdaRank, softmax) | optimize order, not scores |
| Embeddings | **contrastive / triplet / InfoNCE** | pull positives together, push negatives apart |
| Distributions | **KL divergence** | distillation (match teacher probabilities), VAEs |

## Bandits & Exploration

**Q: How do you choose between options while still learning which is best (ads, recommendations, layouts)?**

The explore–exploit trade-off: always showing the current best wastes no traffic now but may never discover something better.
- **ε-greedy**: exploit with probability 1 − ε, explore uniformly otherwise. Simple; explores blindly.
- **UCB**: pick the arm maximizing `mean + c·√(ln t / n_arm)` — optimism under uncertainty; rarely tried arms get a bonus.
- **Thompson sampling**: keep a posterior per arm (Beta(successes + 1, failures + 1) for clicks), sample from each, play the best sample. Strong in practice, handles delayed feedback well.
- **Contextual bandits** (LinUCB, neural): choose based on user/context features — personalized recommendations and ads.
- Compared with A/B tests: bandits minimize **regret** during the experiment; A/B tests give cleaner estimates of the effect. Use bandits for many short-lived options (headlines, creatives), A/B for decisions you need to measure rigorously.
- Exploration traffic also produces unbiased logged data for training future models (counterfactual / off-policy evaluation with propensity weights).

## Collaborative Filtering & Matrix Factorization

**Q: How do classic recommenders work before deep learning?**

- **Content-based**: recommend items similar (by features/embeddings) to what the user liked. Handles new items; limited serendipity.
- **User–user / item–item CF**: similarity from co-interaction (cosine over rating vectors); item–item ("people who bought X also bought Y") is stable and scalable (Amazon).
- **Matrix factorization**: `rating ≈ uᵤ · vᵢ (+ biases)`, learned on observed entries with L2 regularization via SGD or **ALS** (each step is ridge regression in closed form, parallelizes well). Netflix Prize lesson: biases + MF + ensembling.
- **Implicit feedback** (clicks, watches): no negatives — treat unobserved as weak negatives with confidence weights (weighted ALS) or sample negatives (BPR, two-tower with in-batch negatives).
- **Cold start**: new users → popularity, onboarding, context; new items → content features, exploration.
- Modern systems keep the factorization idea as the **two-tower retrieval model** feeding a heavier ranker.

## Model Interpretability

**Q: How do you explain a model's predictions?**

- **Global**: feature importance (gain/split counts in trees — biased toward high-cardinality features), **permutation importance** (drop in score when a feature is shuffled; model-agnostic), **partial dependence plots** (average prediction as one feature varies).
- **Local** (one prediction): **SHAP** (Shapley values: fair attribution of the prediction to features; exact and fast for trees with TreeSHAP), **LIME** (fit a simple model around the point), counterfactuals ("approved if income were $5k higher").
- Deep nets: saliency/gradient maps, integrated gradients, attention visualization (attention ≠ explanation).
- Caveats: correlated features split credit arbitrarily; explanations describe the **model**, not causality.
- Regulated domains (credit, healthcare) need reason codes, monotonic constraints and fairness audits (compare error rates across groups).

## Time-Series Forecasting

**Q: How would you forecast demand / traffic?**

- **Baselines first**: last value, seasonal naive (same hour last week), moving averages — surprisingly hard to beat.
- Decompose into **trend + seasonality + residual**; check **stationarity** (differencing, log transform).
- Classical: exponential smoothing (Holt-Winters), **ARIMA/SARIMA**, Prophet (trend + seasonality + holidays).
- ML: GBDTs on **lag features** (t−1, t−7, t−364), rolling means, calendar features, holidays, promotions — the workhorse for many related series; deep models (DeepAR, temporal fusion transformers, N-BEATS) and pretrained time-series foundation models for large collections.
- **Evaluation**: rolling-origin **backtesting** (train on the past, test on the next window, slide forward) — never random splits; metrics MAE, MAPE (breaks near zero), sMAPE, quantile loss for intervals.
- Forecast **intervals**, not just points (quantile regression), and watch for regime changes (holidays, outages, pandemics).
