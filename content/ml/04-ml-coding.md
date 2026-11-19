# ML Coding from Scratch

"Implement X in NumPy" questions. Each snippet is self-contained and ends with a quick check — copy it into a file and run it (`pip install numpy`).

Tips: vectorize (no Python loops over samples), mind shapes (write them in comments), use numerically stable formulas (log-sum-exp), and seed randomness.

## Softmax & Cross-Entropy (Numerically Stable)

**Key idea:** subtract the row max before `exp` (softmax is shift-invariant) to avoid overflow; compute log-softmax with log-sum-exp. Gradient of CE w.r.t. logits is `softmax − one_hot`.

```python
import numpy as np

def softmax(z):                       # z: (N, C)
    z = z - z.max(axis=1, keepdims=True)
    e = np.exp(z)
    return e / e.sum(axis=1, keepdims=True)

def cross_entropy(logits, y):         # y: (N,) int labels
    z = logits - logits.max(axis=1, keepdims=True)
    log_probs = z - np.log(np.exp(z).sum(axis=1, keepdims=True))
    loss = -log_probs[np.arange(len(y)), y].mean()
    grad = softmax(logits)
    grad[np.arange(len(y)), y] -= 1
    return loss, grad / len(y)        # dL/dlogits

# check
logits = np.array([[1000.0, 1001.0, 1002.0], [1.0, 2.0, 3.0]])
p = softmax(logits)
assert np.allclose(p[0], p[1]) and np.allclose(p.sum(1), 1)
loss, g = cross_entropy(logits, np.array([2, 0]))
assert np.isfinite(loss) and np.allclose(g.sum(1), 0)
print("ok", round(loss, 4))
```

## Linear Regression with Gradient Descent

**Key idea:** MSE loss `(1/N)‖Xw + b − y‖²`; gradient `dw = (2/N)·Xᵀ(ŷ − y)`. Compare with the closed form `(XᵀX)⁻¹Xᵀy`.

```python
import numpy as np

def fit_linear(X, y, lr=0.1, epochs=2000):
    N, D = X.shape
    w, b = np.zeros(D), 0.0
    for _ in range(epochs):
        err = X @ w + b - y           # (N,)
        w -= lr * 2 / N * (X.T @ err)
        b -= lr * 2 / N * err.sum()
    return w, b

rng = np.random.default_rng(0)
X = rng.normal(size=(500, 3))
true_w = np.array([2.0, -1.0, 0.5])
y = X @ true_w + 3 + 0.01 * rng.normal(size=500)
w, b = fit_linear(X, y)
assert np.allclose(w, true_w, atol=0.01) and abs(b - 3) < 0.01
print("ok", w.round(3), round(b, 3))
```

## Logistic Regression

**Key idea:** `p = σ(Xw + b)`, BCE loss, gradient `Xᵀ(p − y)/N` (+ L2 term). Clip probabilities inside the log.

```python
import numpy as np

def sigmoid(z):
    return 1 / (1 + np.exp(-z))

def fit_logistic(X, y, lr=0.5, epochs=1000, l2=1e-3):
    N, D = X.shape
    w, b = np.zeros(D), 0.0
    for _ in range(epochs):
        p = sigmoid(X @ w + b)
        w -= lr * (X.T @ (p - y) / N + l2 * w)
        b -= lr * (p - y).mean()
    return w, b

def bce(p, y, eps=1e-12):
    p = np.clip(p, eps, 1 - eps)
    return -(y * np.log(p) + (1 - y) * np.log(1 - p)).mean()

rng = np.random.default_rng(1)
X = np.vstack([rng.normal(-2, 1, (200, 2)), rng.normal(2, 1, (200, 2))])
y = np.r_[np.zeros(200), np.ones(200)]
w, b = fit_logistic(X, y)
acc = ((sigmoid(X @ w + b) > 0.5) == y).mean()
assert acc > 0.95
print("ok acc", acc, "loss", round(bce(sigmoid(X @ w + b), y), 4))
```

## K-Means

**Key idea:** alternate assignment (argmin distance, vectorized with broadcasting) and centroid update; k-means++ init; stop when centroids stop moving.

```python
import numpy as np

def kmeans(X, k, iters=100, seed=0):
    rng = np.random.default_rng(seed)
    C = [X[rng.integers(len(X))]]                          # k-means++ init
    for _ in range(1, k):
        d2 = ((X[:, None, :] - np.array(C)[None]) ** 2).sum(-1).min(1)
        C.append(X[rng.choice(len(X), p=d2 / d2.sum())])
    C = np.array(C)
    for _ in range(iters):
        dist = ((X[:, None, :] - C[None]) ** 2).sum(-1)    # (N, k)
        labels = dist.argmin(1)
        newC = np.array([X[labels == j].mean(0) if np.any(labels == j) else C[j] for j in range(k)])
        if np.allclose(newC, C):
            break
        C = newC
    return C, labels

rng = np.random.default_rng(0)
centers = np.array([[0, 0], [5, 5], [0, 5]])
X = np.vstack([c + 0.3 * rng.normal(size=(100, 2)) for c in centers])
C, labels = kmeans(X, 3)
assert all(np.min(np.linalg.norm(C - c, axis=1)) < 0.2 for c in centers)
print("ok", np.round(C[np.lexsort(C.T[::-1])], 2).tolist())
```

## K-Nearest Neighbors

**Key idea:** pairwise squared distances via `‖a‖² + ‖b‖² − 2a·b` (one matrix multiply), `argpartition` for the k smallest, majority vote.

```python
import numpy as np

def knn_predict(X_train, y_train, X_test, k=5):
    d2 = (X_test ** 2).sum(1)[:, None] + (X_train ** 2).sum(1)[None, :] - 2 * X_test @ X_train.T
    idx = np.argpartition(d2, k, axis=1)[:, :k]          # (M, k) nearest indices
    votes = y_train[idx]
    return np.array([np.bincount(v).argmax() for v in votes])

rng = np.random.default_rng(2)
X = np.vstack([rng.normal(-2, 1, (100, 2)), rng.normal(2, 1, (100, 2))])
y = np.r_[np.zeros(100, int), np.ones(100, int)]
pred = knn_predict(X, y, np.array([[-2.0, -2.0], [2.0, 2.0]]), k=7)
assert pred.tolist() == [0, 1]
print("ok", pred)
```

## Scaled Dot-Product & Multi-Head Attention

**Key idea:** `softmax(QKᵀ/√d)V` with an optional causal mask (`-inf` above the diagonal). Multi-head = reshape `(N, T, D) → (N, h, T, D/h)`, attend per head, merge, project.

```python
import numpy as np

def attention(Q, K, V, causal=False):          # (..., T, d)
    d = Q.shape[-1]
    scores = Q @ np.swapaxes(K, -1, -2) / np.sqrt(d)   # (..., T, T)
    if causal:
        T = scores.shape[-1]
        scores = np.where(np.tril(np.ones((T, T), bool)), scores, -np.inf)
    scores = scores - scores.max(-1, keepdims=True)
    w = np.exp(scores)
    w /= w.sum(-1, keepdims=True)
    return w @ V, w

def multi_head(x, Wq, Wk, Wv, Wo, h, causal=True):    # x: (N, T, D)
    N, T, D = x.shape
    split = lambda t: t.reshape(N, T, h, D // h).transpose(0, 2, 1, 3)   # (N, h, T, D/h)
    out, _ = attention(split(x @ Wq), split(x @ Wk), split(x @ Wv), causal)
    return out.transpose(0, 2, 1, 3).reshape(N, T, D) @ Wo

rng = np.random.default_rng(0)
N, T, D, h = 2, 5, 16, 4
x = rng.normal(size=(N, T, D))
W = [rng.normal(size=(D, D)) / np.sqrt(D) for _ in range(4)]
y = multi_head(x, *W, h=h)
assert y.shape == (N, T, D)
_, w = attention(x, x, x, causal=True)
assert np.allclose(np.triu(w[0], 1), 0) and np.allclose(w.sum(-1), 1)   # no peeking at the future
x2 = x.copy(); x2[:, -1] += 10                                           # changing the last token
assert np.allclose(multi_head(x2, *W, h=h)[:, :-1], y[:, :-1])           # doesn't affect earlier outputs
print("ok", y.shape)
```

## Two-Layer MLP with Backprop

**Key idea:** forward `X → ReLU(XW1+b1) → logits`, softmax-CE gradient `p − y`, then chain rule backwards. Train on a non-linear problem (XOR-like) to prove it learns.

```python
import numpy as np

rng = np.random.default_rng(0)
X = rng.uniform(-1, 1, (400, 2))
y = ((X[:, 0] * X[:, 1]) > 0).astype(int)              # XOR-like quadrants

H, C, lr = 32, 2, 0.5
W1 = rng.normal(0, np.sqrt(2 / 2), (2, H)); b1 = np.zeros(H)   # He init
W2 = rng.normal(0, np.sqrt(2 / H), (H, C)); b2 = np.zeros(C)

for step in range(2000):
    z1 = X @ W1 + b1; a1 = np.maximum(z1, 0)            # forward
    logits = a1 @ W2 + b2
    p = np.exp(logits - logits.max(1, keepdims=True)); p /= p.sum(1, keepdims=True)
    dlogits = p.copy(); dlogits[np.arange(len(y)), y] -= 1; dlogits /= len(y)   # backward
    dW2 = a1.T @ dlogits; db2 = dlogits.sum(0)
    da1 = dlogits @ W2.T; dz1 = da1 * (z1 > 0)
    dW1 = X.T @ dz1; db1 = dz1.sum(0)
    for P, G in ((W1, dW1), (b1, db1), (W2, dW2), (b2, db2)):
        P -= lr * G

acc = (p.argmax(1) == y).mean()
assert acc > 0.95
print("ok train acc", acc)
```

## Decision Tree Split (Gini)

**Key idea:** for each feature, sort values and sweep thresholds while maintaining left/right class counts → O(N log N) per feature instead of O(N²).

```python
import numpy as np

def gini(counts):
    n = counts.sum()
    return 0.0 if n == 0 else 1.0 - ((counts / n) ** 2).sum()

def best_split(X, y, n_classes):
    best = (np.inf, None, None)                        # (weighted gini, feature, threshold)
    N = len(y)
    for f in range(X.shape[1]):
        order = np.argsort(X[:, f])
        xs, ys = X[order, f], y[order]
        left, right = np.zeros(n_classes), np.bincount(ys, minlength=n_classes).astype(float)
        for i in range(1, N):
            left[ys[i - 1]] += 1; right[ys[i - 1]] -= 1
            if xs[i] == xs[i - 1]:
                continue                               # can't split between equal values
            score = (i * gini(left) + (N - i) * gini(right)) / N
            if score < best[0]:
                best = (score, f, (xs[i] + xs[i - 1]) / 2)
    return best

X = np.array([[2.0, 1.0], [3.0, 5.0], [1.0, 2.0], [8.0, 1.5], [9.0, 4.0], [7.0, 3.0]])
y = np.array([0, 0, 0, 1, 1, 1])
score, f, t = best_split(X, y, 2)
assert f == 0 and 3 < t < 7 and score == 0
print("ok feature", f, "threshold", t)
```

## ROC-AUC from Scratch

**Key idea:** AUC = probability a random positive outranks a random negative = normalized **Mann–Whitney U** statistic computed from ranks (ties get average ranks). O(N log N).

```python
import numpy as np

def roc_auc(y_true, scores):
    y_true, scores = np.asarray(y_true), np.asarray(scores, float)
    order = np.argsort(scores)
    ranks = np.empty(len(scores))
    s = scores[order]
    i = 0
    while i < len(s):                                  # average ranks for ties
        j = i
        while j + 1 < len(s) and s[j + 1] == s[i]:
            j += 1
        ranks[order[i:j + 1]] = (i + j) / 2 + 1
        i = j + 1
    n_pos = y_true.sum(); n_neg = len(y_true) - n_pos
    return (ranks[y_true == 1].sum() - n_pos * (n_pos + 1) / 2) / (n_pos * n_neg)

assert roc_auc([0, 0, 1, 1], [0.1, 0.4, 0.35, 0.8]) == 0.75
assert roc_auc([0, 1], [0.5, 0.5]) == 0.5
assert roc_auc([0, 0, 1], [0.1, 0.2, 0.9]) == 1.0
print("ok")
```

## Conv2D Forward (im2col)

**Key idea:** unfold every receptive field into a row (**im2col**), then the convolution is one matrix multiply with the flattened kernels. Output size `(H + 2P − K)/S + 1`.

```python
import numpy as np

def conv2d(x, w, b, stride=1, pad=0):
    # x: (N, C, H, W), w: (F, C, K, K), b: (F,)
    N, C, H, W = x.shape
    F, _, K, _ = w.shape
    x = np.pad(x, ((0, 0), (0, 0), (pad, pad), (pad, pad)))
    Ho, Wo = (H + 2 * pad - K) // stride + 1, (W + 2 * pad - K) // stride + 1
    cols = np.empty((N, Ho, Wo, C * K * K))
    for i in range(Ho):
        for j in range(Wo):
            patch = x[:, :, i * stride:i * stride + K, j * stride:j * stride + K]
            cols[:, i, j, :] = patch.reshape(N, -1)
    out = cols @ w.reshape(F, -1).T + b                # (N, Ho, Wo, F)
    return out.transpose(0, 3, 1, 2)

x = np.arange(16, dtype=float).reshape(1, 1, 4, 4)
w = np.ones((1, 1, 2, 2))                              # 2x2 box sum
y = conv2d(x, w, np.zeros(1), stride=2)
assert y.shape == (1, 1, 2, 2) and y[0, 0].tolist() == [[10, 18], [42, 50]]
assert conv2d(np.ones((2, 3, 5, 5)), np.ones((4, 3, 3, 3)), np.zeros(4), pad=1).shape == (2, 4, 5, 5)
print("ok")
```


## Batch Normalization (Forward & Backward)

**Key idea:** normalize per feature over the batch, scale/shift with γ, β. The backward pass has a compact closed form; verify it with a **numerical gradient check** (a must-know debugging technique).

```python
import numpy as np

def bn_forward(x, gamma, beta, eps=1e-5):
    mu, var = x.mean(0), x.var(0)
    xhat = (x - mu) / np.sqrt(var + eps)
    return gamma * xhat + beta, (xhat, var, eps, gamma)

def bn_backward(dout, cache):
    xhat, var, eps, gamma = cache
    N = dout.shape[0]
    dgamma, dbeta = (dout * xhat).sum(0), dout.sum(0)
    dxhat = dout * gamma
    dx = (N * dxhat - dxhat.sum(0) - xhat * (dxhat * xhat).sum(0)) / (N * np.sqrt(var + eps))
    return dx, dgamma, dbeta

rng = np.random.default_rng(0)
x, g, b = rng.normal(size=(8, 3)), rng.normal(size=3), rng.normal(size=3)
dout = rng.normal(size=(8, 3))
y, cache = bn_forward(x, g, b)
assert np.allclose(y.mean(0), b) and np.allclose(y.std(0), np.abs(g), atol=1e-3)
dx, dg, db = bn_backward(dout, cache)
num = np.zeros_like(x); h = 1e-5                     # numerical gradient check
for i in np.ndindex(x.shape):
    xp, xm = x.copy(), x.copy(); xp[i] += h; xm[i] -= h
    num[i] = ((bn_forward(xp, g, b)[0] - bn_forward(xm, g, b)[0]) * dout).sum() / (2 * h)
assert np.allclose(dx, num, atol=1e-6)
print("ok, max grad error", float(np.abs(dx - num).max()))
```

## PCA via SVD

**Key idea:** center the data, take the SVD `X = UΣVᵀ`; principal directions are the rows of `Vᵀ`, explained variance is `σ² / (n − 1)`.

```python
import numpy as np

def pca(X, k):
    Xc = X - X.mean(0)
    U, S, Vt = np.linalg.svd(Xc, full_matrices=False)
    var = S ** 2 / (len(X) - 1)
    return Xc @ Vt[:k].T, Vt[:k], var[:k] / var.sum()

rng = np.random.default_rng(0)
z = rng.normal(size=(500, 1))
X = np.hstack([z, 2 * z, -z]) + 0.01 * rng.normal(size=(500, 3))   # essentially 1-D data
proj, comps, ratio = pca(X, 1)
assert ratio[0] > 0.999
assert np.allclose(np.abs(comps[0]), np.array([1, 2, 1]) / np.sqrt(6), atol=1e-3)
print("ok explained", round(float(ratio[0]), 5))
```

## Adam Optimizer

**Key idea:** first moment (momentum) + second moment (per-parameter scale) with **bias correction** for the zero initialization. Test it on an ill-conditioned quadratic where plain SGD with the same step is slow.

```python
import numpy as np

class Adam:
    def __init__(self, lr=0.1, b1=0.9, b2=0.999, eps=1e-8):
        self.lr, self.b1, self.b2, self.eps, self.t, self.m, self.v = lr, b1, b2, eps, 0, 0, 0

    def step(self, w, g):
        self.t += 1
        self.m = self.b1 * self.m + (1 - self.b1) * g
        self.v = self.b2 * self.v + (1 - self.b2) * g * g
        mhat = self.m / (1 - self.b1 ** self.t)
        vhat = self.v / (1 - self.b2 ** self.t)
        return w - self.lr * mhat / (np.sqrt(vhat) + self.eps)

A = np.diag([100.0, 1.0])                  # f(w) = 0.5 wᵀAw, minimum at 0
grad = lambda w: A @ w
w, opt = np.array([1.0, 1.0]), Adam(lr=0.05)
for _ in range(500):
    w = opt.step(w, grad(w))
assert np.linalg.norm(w) < 1e-2
print("ok", w.round(5))
```

## Gaussian Naive Bayes

**Key idea:** per class, model each feature as an independent Gaussian. Predict `argmax_c log p(c) + Σ_j log N(x_j; μ_cj, σ²_cj)` — work in log space to avoid underflow.

```python
import numpy as np

class GaussianNB:
    def fit(self, X, y):
        self.classes = np.unique(y)
        self.mu = np.array([X[y == c].mean(0) for c in self.classes])
        self.var = np.array([X[y == c].var(0) + 1e-9 for c in self.classes])
        self.logprior = np.log(np.array([(y == c).mean() for c in self.classes]))
        return self

    def predict(self, X):
        # (N, C): log prior + sum of per-feature Gaussian log-likelihoods
        ll = -0.5 * (np.log(2 * np.pi * self.var)[None] + (X[:, None, :] - self.mu[None]) ** 2 / self.var[None]).sum(-1)
        return self.classes[(ll + self.logprior).argmax(1)]

rng = np.random.default_rng(3)
X = np.vstack([rng.normal([0, 0], 1, (200, 2)), rng.normal([3, 3], 1, (200, 2))])
y = np.r_[np.zeros(200, int), np.ones(200, int)]
acc = (GaussianNB().fit(X, y).predict(X) == y).mean()
assert acc > 0.95
print("ok acc", acc)
```

## Beam Search Decoding

**Key idea:** keep the `b` highest-scoring partial sequences (sum of log-probs); expand each with every token, keep the top `b`; finished beams (EOS) are set aside. Greedy decoding is beam size 1 — and can miss the best sequence.

```python
import numpy as np

def beam_search(step_logprobs, beam=3, max_len=5, eos=0):
    """step_logprobs(prefix) -> log-prob vector over the vocabulary for the next token."""
    beams, done = [((), 0.0)], []
    for _ in range(max_len):
        cand = []
        for seq, score in beams:
            lp = step_logprobs(seq)
            for tok in np.argsort(lp)[-beam:]:
                cand.append((seq + (int(tok),), score + float(lp[tok])))
        cand.sort(key=lambda c: c[1], reverse=True)
        beams = []
        for seq, score in cand[:beam]:
            (done if seq[-1] == eos else beams).append((seq, score))
        if not beams:
            break
    return max(done + beams, key=lambda c: c[1])

# toy model: first token 1 looks best (0.6) but leads to a weak continuation; token 2 (0.4) leads to a sure EOS
P = {(): [0.0, 0.6, 0.4], (1,): [0.3, 0.4, 0.3], (2,): [1.0, 0.0, 0.0]}
def model(seq):
    p = np.array(P.get(seq, [1.0, 0.0, 0.0])) + 1e-12
    return np.log(p)
greedy = beam_search(model, beam=1, max_len=2)
best = beam_search(model, beam=3, max_len=2)
assert greedy[0] == (1, 1) and best[0] == (2, 0)          # beam finds the higher-probability sequence
print("ok greedy", greedy[0], round(np.exp(greedy[1]), 3), "beam", best[0], round(np.exp(best[1]), 3))
```
