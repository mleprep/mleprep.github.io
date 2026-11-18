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
