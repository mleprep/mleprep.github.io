# leetcode_review

A concise, browsable interview-prep site: **https://mleprep.github.io/**

| Section | What's inside |
|---|---|
| **Algorithms** | **446 problems** (96 Easy / 284 Medium / 66 Hard) covering **100% of Blind 75, NeetCode 150, Grind 169, LeetCode 75 and LeetCode Top Interview 150**, plus frequently asked company problems, grouped into 20 patterns. Each has the key idea, an ASCII sketch, complexity, **Python + Java** solutions with tests, and related-problem links. Premium-only problems are marked and include a short problem statement. |
| **System Design** | 76 articles: interview framework, estimation numbers and the zero-to-millions scaling walkthrough; 27 building blocks (caching, sharding, consistent hashing, CAP, queues, rate limiting, real-time transports, API design, Bloom filter / Count-Min / HyperLogLog, search, stream processing, geo indexes, distributed transactions, Raft, auth, choosing a database…); Hello Interview-style patterns (scaling reads/writes, contention, workflows, long-running tasks, large files); 37 designs (URL shortener, news feed, chat, Uber, YouTube, Dropbox, Kafka, payments, Google Docs, S3, stock exchange, Google Maps, nearby friends, digital wallet, Tinder, LeetCode, post search, auctions, food delivery, top-K trending, LLM serving…). |
| **ML / Deep Learning** | 72 articles: ML fundamentals Q&A, deep learning (incl. MoE, distributed training, GNNs), LLMs & GenAI (RLHF/DPO, LoRA, RAG, KV cache, agents), 15 tested NumPy implementations from scratch (attention, MLP & BatchNorm backprop, Adam, k-means, PCA, beam search…), ML system design (recsys, search, ads, fraud, ETA, support assistant, A/B testing). |

**Solve in the browser.** Every problem page has a *Solution / Solve in browser* switch. Solve mode hides the notes, gives you an
editor with a starter stub, and **Run tests** executes your Python against the problem's tests entirely on your machine — Python
runs in a Web Worker via [Pyodide](https://pyodide.org) (CPython compiled to WebAssembly), so there is no server. You get per-test
got/expected, runtime errors with line numbers, `print` output and a 10 s time limit; drafts are saved per problem and an accepted
solution marks the problem done. Tests accept any valid answer where LeetCode does (any order, any valid topological order, any
height-balanced BST…). The ML coding snippets have a **Run** button too (NumPy loads on demand). Java is not judged in the
browser: there is no free, reliable way to run `javac` client-side, so Java solutions stay read/copy/run-locally.

**Browse** (`#/browse`) filters every problem by level, status (todo/done), **study list** (Blind 75, NeetCode 150, Grind 169,
LeetCode 75, Top Interview 150 — optionally in the list's own order), pattern tag, topic and **company** (all time or last 6 months,
ranked by that company's ask frequency); every difficulty badge, tag and topic label on the site is a clickable filter, and filters
live in the URL so views can be bookmarked. Progress checkboxes are saved in your browser. Press `/` to search. The layout scales
fonts with screen size and shows code side-by-side with the explanation on wide screens.

**Algorithm topics:** Arrays & Hashing · Two Pointers · Sliding Window · Stack & Monotonic Stack · Binary Search · Linked List · Trees · Tries · Heap · Backtracking · Graphs · Advanced Graphs (Dijkstra, Bellman-Ford, Floyd-Warshall, MST, Tarjan, Hierholzer) · 1-D DP · 2-D DP · Greedy · Intervals · Math & Geometry · Bit Manipulation · Prefix Sums & Range Queries (difference arrays, Fenwick tree) · Design.

## Data sources

- **Company tags**: the public dataset [liquidslr/leetcode-company-wise-problems](https://github.com/liquidslr/leetcode-company-wise-problems)
  (358 companies cover all 446 problems). Frequency is the dataset's 0–100 score, relative within each company.
- **Study lists**: NeetCode's public problem data (Blind 75 / NeetCode 150), the Grind 169 site, and LeetCode's public study plans.

```bash
git clone --depth 1 https://github.com/liquidslr/leetcode-company-wise-problems.git /tmp/co
python tools/import_companies.py /tmp/co     # rewrites data/companies.json for the problems in content/dsa
python tools/import_lists.py                 # rewrites data/lists.json and prints any list problems the site lacks
```

## Run solutions locally

Every problem is also generated as a standalone runnable file with tests:

```bash
cd python && python 0001_two_sum.py          # prints OK when all asserts pass
cd java   && java P0001TwoSum.java            # Java 11+ single-file launch, prints OK
```

The `Solution` class inside each file is exactly what you paste into LeetCode. Python files import `python/lc.py`
(ListNode / TreeNode + `make_list`, `list_vals`, `make_tree`, `tree_vals`, `find_node`); Java files inline the same helpers
(`ListNode.of(1, 2, 3)`, `TreeNode.of(3, 9, 20, null, null, 15, 7)`).

## Repo layout

```
content/                 <- source of truth (markdown, one file per topic)
  dsa/NN-topic.md            problems: idea + ```viz + ```python / ```java (each with a "# test" / "// test" part)
  system-design/*.md         articles (one "## " heading per article)
  ml/*.md                    articles
build.py                 <- content -> data/*.json (site, incl. a starter stub per problem) + python/*.py + java/*.java
tools/import_companies.py <- public company-tag dataset -> data/companies.json
tools/import_lists.py    <- public study lists -> data/lists.json
index.html, assets/      <- static single-page site (no build tools; marked, highlight.js, CodeMirror, Pyodide from CDNs)
  assets/judge.py            the in-browser judge (also runs under CPython: `python assets/judge.py` self-checks every problem)
  assets/py-worker.js        Web Worker hosting Pyodide
data/, python/, java/    <- generated by build.py (committed so GitHub Pages and local runs just work)
```

## Edit or add content

1. Edit `content/dsa/<topic>.md` (the problem format is documented at the top of `build.py`).
2. `python build.py --test` regenerates everything and runs **all Python + Java test files** (892 for 446 problems) plus the ML
   NumPy snippets (skipped if numpy isn't installed) in parallel. Java needs `java` on PATH or `JAVA_HOME`. Add problem numbers to
   test only those: `python build.py --test 1 146`.
3. `python assets/judge.py` checks the browser judge: every reference solution must pass its tests and every starter stub must fail.
4. Preview: `python -m http.server` and open http://localhost:8000.
5. Commit and push to `main`; GitHub Pages serves the repo root.
