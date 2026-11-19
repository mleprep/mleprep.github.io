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

## IoU & Non-Maximum Suppression

**Key idea:** IoU = intersection area / union area of two boxes (clip negative overlaps to 0). NMS keeps the highest-scoring box, drops every remaining box whose IoU with it exceeds a threshold, and repeats — removing duplicate detections of the same object.

```python
import numpy as np

def iou(box, boxes):                      # box: (4,), boxes: (N, 4) as x1, y1, x2, y2
    x1 = np.maximum(box[0], boxes[:, 0]); y1 = np.maximum(box[1], boxes[:, 1])
    x2 = np.minimum(box[2], boxes[:, 2]); y2 = np.minimum(box[3], boxes[:, 3])
    inter = np.clip(x2 - x1, 0, None) * np.clip(y2 - y1, 0, None)
    area = lambda b: (b[..., 2] - b[..., 0]) * (b[..., 3] - b[..., 1])
    return inter / (area(box) + area(boxes) - inter)

def nms(boxes, scores, thresh=0.5):
    order = np.argsort(-scores)           # best first
    keep = []
    while order.size:
        i = order[0]
        keep.append(int(i))
        rest = order[1:]
        order = rest[iou(boxes[i], boxes[rest]) <= thresh]
    return keep

boxes = np.array([[0, 0, 10, 10], [1, 1, 11, 11], [20, 20, 30, 30], [21, 21, 29, 29], [0, 0, 10, 10.5]], float)
scores = np.array([0.9, 0.8, 0.75, 0.7, 0.6])
assert np.isclose(iou(boxes[0], boxes[1:2])[0], 81 / 119)      # 9x9 overlap / (100 + 100 - 81)
assert iou(boxes[0], boxes[2:3])[0] == 0
assert nms(boxes, scores, 0.5) == [0, 2]
print("ok", nms(boxes, scores, 0.5))
```

Follow-ups: soft-NMS (decay scores instead of deleting), class-wise NMS, and why DETR-style detectors don't need NMS (set prediction with Hungarian matching).

## Byte-Pair Encoding (BPE) Tokenizer

**Key idea:** start from characters (plus an end-of-word marker), repeatedly count adjacent symbol pairs weighted by word frequency and merge the most frequent pair into a new symbol. Encoding a new word replays the learned merges **in training order**.

```python
from collections import Counter

def merge_word(word, pair):
    out, i = [], 0
    while i < len(word):
        if i + 1 < len(word) and (word[i], word[i + 1]) == pair:
            out.append(word[i] + word[i + 1]); i += 2
        else:
            out.append(word[i]); i += 1
    return tuple(out)

def train_bpe(corpus, num_merges):
    words = Counter(tuple(w) + ("</w>",) for w in corpus.split())   # symbol tuple -> frequency
    merges = []
    for _ in range(num_merges):
        pairs = Counter()
        for w, f in words.items():
            for a, b in zip(w, w[1:]):
                pairs[a, b] += f
        if not pairs:
            break
        best = max(pairs, key=pairs.get)   # ties -> first seen
        merges.append(best)
        new = Counter()
        for w, f in words.items():
            new[merge_word(w, best)] += f
        words = new
    return merges

def encode(word, merges):
    syms = tuple(word) + ("</w>",)
    for pair in merges:
        syms = merge_word(syms, pair)
    return list(syms)

corpus = " ".join(["low"] * 5 + ["lower"] * 2 + ["newest"] * 6 + ["widest"] * 3)   # Sennrich et al.'s example
merges = train_bpe(corpus, 10)
assert merges[:5] == [("e", "s"), ("es", "t"), ("est", "</w>"), ("l", "o"), ("lo", "w")]
assert encode("lowest", merges) == ["low", "est</w>"]        # unseen word -> known subwords
assert encode("newer", merges)[0] == "new"
print("ok", merges[:5], encode("lowest", merges))
```

Production BPE (GPT-2/tiktoken) works on **bytes** (no unknown tokens), pre-splits text with a regex (so merges don't cross word/number boundaries) and uses a priority queue for fast training.

## Sampling: Temperature, Top-k, Top-p

**Key idea:** divide logits by the temperature, optionally keep only the k largest (top-k) or the smallest set of tokens whose probability mass reaches p (nucleus / top-p), renormalize and sample.

```python
import numpy as np

def softmax(z):
    e = np.exp(z - z.max())
    return e / e.sum()

def next_token_probs(logits, temperature=1.0, top_k=None, top_p=None):
    z = np.asarray(logits, float) / max(temperature, 1e-8)
    if top_k is not None:
        kth = np.sort(z)[-top_k]
        z = np.where(z < kth, -np.inf, z)
    p = softmax(z)
    if top_p is not None:
        order = np.argsort(-p)
        cum = np.cumsum(p[order])
        cutoff = np.searchsorted(cum, top_p) + 1          # smallest prefix with mass >= top_p
        keep = np.zeros_like(p, bool)
        keep[order[:cutoff]] = True
        p = np.where(keep, p, 0.0)
        p /= p.sum()
    return p

def sample(logits, rng, **kw):
    p = next_token_probs(logits, **kw)
    return rng.choice(len(p), p=p)

logits = [2.0, 1.0, 0.0, -1.0]             # softmax ≈ [0.644, 0.237, 0.087, 0.032]
assert np.count_nonzero(next_token_probs(logits, top_k=2)) == 2
assert np.count_nonzero(next_token_probs(logits, top_p=0.5)) == 1     # 0.644 already >= 0.5
assert np.count_nonzero(next_token_probs(logits, top_p=0.9)) == 3     # 0.644 + 0.237 < 0.9
assert next_token_probs(logits, temperature=0.01)[0] > 0.999           # T -> 0 = greedy
hot = next_token_probs(logits, temperature=10.0)
assert hot.max() - hot.min() < 0.1                                      # high T -> near uniform
rng = np.random.default_rng(0)
draws = [sample(logits, rng, top_k=2) for _ in range(2000)]
assert set(draws) == {0, 1}
print("ok", np.round(next_token_probs(logits, top_p=0.9), 3))
```

## LayerNorm & RMSNorm (Forward & Backward)

**Key idea:** LayerNorm normalizes each **row** (token) over its features, then applies learned scale γ and shift β. Backward: `dx = (dx̂ − mean(dx̂) − x̂·mean(dx̂·x̂)) / σ` with `dx̂ = dy·γ`. RMSNorm (LLaMA) skips the mean: `y = γ·x / rms(x)`.

```python
import numpy as np

def layernorm_forward(x, gamma, beta, eps=1e-5):      # x: (N, D)
    mu = x.mean(-1, keepdims=True)
    var = x.var(-1, keepdims=True)
    xhat = (x - mu) / np.sqrt(var + eps)
    return gamma * xhat + beta, (xhat, var, gamma, eps)

def layernorm_backward(dy, cache):
    xhat, var, gamma, eps = cache
    dgamma, dbeta = (dy * xhat).sum(0), dy.sum(0)
    dxhat = dy * gamma
    dx = (dxhat - dxhat.mean(-1, keepdims=True) - xhat * (dxhat * xhat).mean(-1, keepdims=True)) / np.sqrt(var + eps)
    return dx, dgamma, dbeta

def rmsnorm(x, gamma, eps=1e-6):
    return gamma * x / np.sqrt((x * x).mean(-1, keepdims=True) + eps)

rng = np.random.default_rng(0)
x, gamma, beta = rng.normal(size=(4, 6)), rng.normal(size=6), rng.normal(size=6)
y, cache = layernorm_forward(x, np.ones(6), np.zeros(6))
assert np.allclose(y.mean(-1), 0) and np.allclose(y.std(-1), 1, atol=1e-3)

# gradient check against finite differences of loss = sum(y * w)
w = rng.normal(size=(4, 6))
y, cache = layernorm_forward(x, gamma, beta)
dx, dgamma, dbeta = layernorm_backward(w, cache)
num, h = np.zeros_like(x), 1e-6
for i in np.ndindex(x.shape):
    xp, xm = x.copy(), x.copy()
    xp[i] += h; xm[i] -= h
    num[i] = ((layernorm_forward(xp, gamma, beta)[0] - layernorm_forward(xm, gamma, beta)[0]) * w).sum() / (2 * h)
assert np.abs(num - dx).max() < 1e-6
r = rmsnorm(x, np.ones(6))
assert np.allclose(np.sqrt((r * r).mean(-1)), 1, atol=1e-4)
print("ok, max grad error", np.abs(num - dx).max())
```

## Positional Encodings: Sinusoidal & RoPE

**Key idea:** sinusoidal encodings add `sin/cos(pos / 10000^(2i/d))` to the embeddings. **RoPE** instead rotates each pair of query/key dimensions by an angle proportional to the position, so the dot product `q_m · k_n` depends only on the **offset** `m − n`.

```python
import numpy as np

def sinusoidal(n_pos, d):
    pos = np.arange(n_pos)[:, None]
    angle = pos / 10000 ** (np.arange(0, d, 2)[None, :] / d)
    pe = np.zeros((n_pos, d))
    pe[:, 0::2], pe[:, 1::2] = np.sin(angle), np.cos(angle)
    return pe

def rope(x, pos, base=10000.0):           # x: (..., d), d even; rotate pairs (x[2i], x[2i+1])
    d = x.shape[-1]
    ang = np.asarray(pos)[..., None] * base ** (-np.arange(0, d, 2) / d)
    cos, sin = np.cos(ang), np.sin(ang)
    x1, x2 = x[..., 0::2], x[..., 1::2]
    out = np.empty_like(x)
    out[..., 0::2] = x1 * cos - x2 * sin
    out[..., 1::2] = x1 * sin + x2 * cos
    return out

pe = sinusoidal(50, 16)
assert pe.shape == (50, 16) and np.allclose(pe[0], np.tile([0, 1], 8))
rng = np.random.default_rng(0)
q, k = rng.normal(size=8), rng.normal(size=8)
assert np.isclose(rope(q, 3) @ rope(k, 7), rope(q, 10) @ rope(k, 14))   # only the offset (4) matters
assert np.isclose(np.linalg.norm(rope(q, 5)), np.linalg.norm(q))        # rotation preserves norms
X = rng.normal(size=(5, 8))
assert np.allclose(rope(X, np.arange(5))[2], rope(X[2], 2))              # batched == per position
print("ok", round(float(rope(q, 3) @ rope(k, 7)), 4))
```

## Causal Self-Attention with a KV Cache

**Key idea:** during generation, keys and values of past tokens never change, so cache them; each new token computes only its own query/key/value and attends over the cache — O(T) work per step instead of recomputing O(T²) attention. The result must equal full causal attention.

```python
import numpy as np

def causal_attention(x, Wq, Wk, Wv):           # x: (T, d) -> (T, d_v), single head
    q, k, v = x @ Wq, x @ Wk, x @ Wv
    s = q @ k.T / np.sqrt(k.shape[-1])
    s = np.where(np.tril(np.ones(s.shape, bool)), s, -np.inf)   # token t sees 0..t
    w = np.exp(s - s.max(-1, keepdims=True))
    return (w / w.sum(-1, keepdims=True)) @ v

class KVCache:
    def __init__(self):
        self.k, self.v = [], []

    def step(self, x_t, Wq, Wk, Wv):           # x_t: (d,) — only the newest token
        q = x_t @ Wq
        self.k.append(x_t @ Wk)
        self.v.append(x_t @ Wv)
        K, V = np.stack(self.k), np.stack(self.v)
        s = K @ q / np.sqrt(K.shape[-1])
        w = np.exp(s - s.max())
        return (w / w.sum()) @ V

rng = np.random.default_rng(0)
T, d = 6, 8
x = rng.normal(size=(T, d))
Wq, Wk, Wv = (rng.normal(size=(d, d)) / np.sqrt(d) for _ in range(3))
full = causal_attention(x, Wq, Wk, Wv)
cache = KVCache()
incremental = np.stack([cache.step(x[t], Wq, Wk, Wv) for t in range(T)])
assert np.allclose(full, incremental)
print("ok, cache holds", len(cache.k), "keys; memory per token per layer = 2 * d values")
```

Real KV-cache memory: `2 · layers · kv_heads · head_dim · bytes` per token — which is why GQA/MQA (fewer KV heads) and paged allocation matter.

## Matrix Factorization (ALS) for Recommendations

**Key idea:** approximate the sparse user × item rating matrix as `U Vᵀ` with k-dimensional factors, fit only on observed entries with L2 regularization. **Alternating least squares**: fixing V makes each user's factor a small ridge regression (closed form), then swap.

The objective is non-convex: from a small random start, ALS sometimes stalls in a bad local minimum (try it — some seeds give a held-out error 50× worse). **Spectral initialization** — the top-k SVD of the observed matrix rescaled by `1/(fraction observed)` — starts near the right subspace and fixes this.

```python
import numpy as np

def als(R, mask, k=3, lam=0.1, iters=30):
    P, s, Qt = np.linalg.svd(R / mask.mean(), full_matrices=False)   # spectral init
    U, V = P[:, :k] * np.sqrt(s[:k]), Qt[:k].T * np.sqrt(s[:k])
    reg = lam * np.eye(k)
    for _ in range(iters):
        for u in range(R.shape[0]):                     # ridge regression per user
            Vu = V[mask[u]]
            U[u] = np.linalg.solve(Vu.T @ Vu + reg, Vu.T @ R[u, mask[u]])
        for i in range(R.shape[1]):                     # ... and per item
            Ui = U[mask[:, i]]
            V[i] = np.linalg.solve(Ui.T @ Ui + reg, Ui.T @ R[mask[:, i], i])
    return U, V

rng = np.random.default_rng(1)
true = rng.normal(size=(40, 3)) @ rng.normal(size=(25, 3)).T    # rank-3 "taste" matrix
mask = rng.random(true.shape) < 0.5                              # observe half the ratings
U, V = als(np.where(mask, true, 0), mask)
pred = U @ V.T
rmse = lambda m: np.sqrt(((pred - true)[m] ** 2).mean())
assert rmse(mask) < 0.1 and rmse(~mask) < 0.3                    # generalizes to unseen ratings
print("ok, train rmse", round(float(rmse(mask)), 3), "held-out rmse", round(float(rmse(~mask)), 3))
```

Implicit feedback (clicks/views) uses weighted ALS with confidence weights, or a two-tower model trained with sampled negatives.

## Ranking Metrics: Precision@k, MAP, MRR, NDCG

**Key idea:** ranking metrics reward putting relevant items **near the top**. NDCG discounts graded relevance by `log2(position + 1)` and normalizes by the ideal ordering; AP averages precision at each relevant hit; MRR looks only at the first relevant hit.

```python
import numpy as np

def precision_at_k(hits, k):
    return float(np.mean(hits[:k]))

def dcg(rels, k, gain="linear"):
    r = np.asarray(rels, float)[:k]
    g = r if gain == "linear" else 2 ** r - 1
    return float((g / np.log2(np.arange(2, r.size + 2))).sum())

def ndcg(rels, k, gain="linear"):
    ideal = dcg(sorted(rels, reverse=True), k, gain)
    return dcg(rels, k, gain) / ideal if ideal > 0 else 0.0

def average_precision(hits):            # over the relevant items that were retrieved
    hits = np.asarray(hits)
    if not hits.any():
        return 0.0
    prec = np.cumsum(hits) / np.arange(1, hits.size + 1)
    return float((prec * hits).sum() / hits.sum())

def mrr(queries):
    return float(np.mean([1 / (np.argmax(h) + 1) if np.any(h) else 0.0 for h in queries]))

assert np.isclose(ndcg([3, 2, 3, 0, 1, 2], 6), 0.961, atol=1e-3)     # the Wikipedia example
assert ndcg([3, 3, 2, 2, 1, 0], 6) == 1.0
assert np.isclose(average_precision([1, 0, 1, 0, 0, 1]), (1 + 2 / 3 + 3 / 6) / 3)
assert mrr([[0, 1, 0], [1, 0, 0], [0, 0, 0]]) == 0.5
assert precision_at_k([1, 0, 1, 1], 2) == 0.5
print("ok, ndcg", round(ndcg([3, 2, 3, 0, 1, 2], 6), 4), "exp-gain ndcg", round(ndcg([3, 2, 3, 0, 1, 2], 6, "exp"), 4))
```

## Dropout (Forward & Backward)

**Key idea:** *inverted* dropout zeros each activation with probability p during training and scales survivors by `1/(1−p)`, so the expected activation is unchanged and inference is a no-op. Backward multiplies the gradient by the same mask.

```python
import numpy as np

def dropout_forward(x, p, rng, train=True):
    if not train or p == 0:
        return x, None
    mask = (rng.random(x.shape) >= p) / (1 - p)
    return x * mask, mask

def dropout_backward(dout, mask):
    return dout if mask is None else dout * mask

rng = np.random.default_rng(0)
x = np.ones((1000, 1000))
y, mask = dropout_forward(x, 0.3, rng)
assert abs(y.mean() - 1.0) < 0.01                       # expectation preserved
assert abs((y == 0).mean() - 0.3) < 0.01                # ~30% dropped
assert np.array_equal(dropout_forward(x, 0.3, rng, train=False)[0], x)   # identity at inference
assert np.array_equal(dropout_backward(np.ones_like(x), mask), mask)
print("ok, kept fraction", round(float((y != 0).mean()), 3))
```

## LoRA Linear Layer

**Key idea:** freeze the pretrained `W (d_out × d_in)` and learn a low-rank update `ΔW = (α/r)·B·A` with `A: r × d_in`, `B: d_out × r`. B starts at **zero** so training starts exactly at the base model; only `r·(d_in + d_out)` parameters train; after training, merge `W + ΔW` for zero inference overhead.

```python
import numpy as np

class LoRALinear:
    def __init__(self, W, r, alpha, rng):
        self.W = W                                           # frozen
        self.A = rng.normal(scale=0.1, size=(r, W.shape[1]))  # trainable
        self.B = np.zeros((W.shape[0], r))                   # trainable, zero init
        self.s = alpha / r

    def forward(self, x):                                    # x: (N, d_in)
        self.xa = x @ self.A.T
        return x @ self.W.T + self.s * self.xa @ self.B.T

    def backward(self, x, dout, lr):                         # SGD step on A and B only
        dB = self.s * dout.T @ self.xa
        dA = self.s * (dout @ self.B).T @ x
        self.B -= lr * dB
        self.A -= lr * dA

    def merged(self):
        return self.W + self.s * self.B @ self.A

rng = np.random.default_rng(0)
d_in, d_out, r = 32, 16, 4
W = rng.normal(size=(d_out, d_in)) / np.sqrt(d_in)
delta = rng.normal(size=(d_out, 2)) @ rng.normal(size=(2, d_in)) / np.sqrt(d_in)   # a rank-2 "new task"
x = rng.normal(size=(256, d_in))
y = x @ (W + delta).T
layer = LoRALinear(W, r=r, alpha=8, rng=rng)
assert np.allclose(layer.forward(x), x @ W.T)               # starts as the base model
loss0 = ((layer.forward(x) - y) ** 2).mean()
for _ in range(1500):
    out = layer.forward(x)
    layer.backward(x, 2 * (out - y) / out.size, lr=0.5)
loss = ((layer.forward(x) - y) ** 2).mean()
assert loss < loss0 * 0.01
assert np.allclose(layer.forward(x), x @ layer.merged().T)  # merging is exact
print(f"ok, loss {loss0:.3f} -> {loss:.5f}; trainable {r * (d_in + d_out)} vs full {d_in * d_out}")
```
