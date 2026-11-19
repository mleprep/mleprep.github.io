# SDE & MLE Prep

Interview prep for software engineers (SDE) and machine learning engineers (MLE) in one site: algorithms with an in-browser
judge, system design, and ML / deep learning. **https://mleprep.github.io/**

| Section | What's inside |
|---|---|
| **Algorithms** | **477 problems** (106 Easy / 299 Medium / 72 Hard) covering **100% of Blind 75, NeetCode 150, NeetCode 250, Grind 169, LeetCode 75 and LeetCode Top Interview 150**, plus frequently asked company problems, grouped into 20 patterns. Each has the key idea, an ASCII sketch, complexity, **Python + Java + Go** solutions with tests, and related-problem links. Premium-only problems are marked and include a short problem statement. |
| **System Design** | 76 articles: interview framework, estimation numbers and the zero-to-millions scaling walkthrough; 27 building blocks (caching, sharding, consistent hashing, CAP, queues, rate limiting, real-time transports, API design, Bloom filter / Count-Min / HyperLogLog, search, stream processing, geo indexes, distributed transactions, Raft, auth, choosing a database…); Hello Interview-style patterns (scaling reads/writes, contention, workflows, long-running tasks, large files); 37 designs (URL shortener, news feed, chat, Uber, YouTube, Dropbox, Kafka, payments, Google Docs, S3, stock exchange, Google Maps, nearby friends, digital wallet, Tinder, LeetCode, post search, auctions, food delivery, top-K trending, LLM serving…). |
| **ML / Deep Learning** | 102 articles: ML fundamentals Q&A (incl. statistics & hypothesis tests, information theory, loss functions, bandits, collaborative filtering, interpretability, forecasting), deep learning (object detection, model compression, diffusion, ViT/CLIP, MoE, distributed training, GNNs), LLMs & GenAI (RLHF/DPO, LoRA, RAG, KV cache, agents, vector search, reasoning models, safety, multimodal), **25 tested NumPy implementations** you can run in the browser (attention, KV cache, MLP/BatchNorm/LayerNorm backprop, Adam, BPE, top-p sampling, RoPE, NMS, ALS, NDCG, LoRA…), and 18 ML system designs (recsys, search, ads, fraud, feed ranking, similar listings, Street View blurring, Smart Compose, text-to-image, RAG assistant, A/B testing…). |

**Solve in the browser.** Every problem page has a *Solution / Solve in browser* switch. Solve mode hides the notes, gives you an
editor with a starter stub, and **Run tests** executes your **Python or Go** against the problem's tests entirely on your machine,
in a Web Worker, so there is no server:

- **Python** runs on [Pyodide](https://pyodide.org) (CPython compiled to WebAssembly, ~10 MB on first use).
- **Go** runs on [Yaegi](https://github.com/traefik/yaegi), a Go interpreter written in Go, compiled to WebAssembly (~3 MB on
  first use; source in `tools/gojudge`). Write it LeetCode style: no `package` clause or imports (`sort`, `strings`, `heap`… are
  imported automatically). Go 1.22 semantics, plus the `min`/`max` builtins and `slices.Sort`/`Max`/`BinarySearch`… Known
  interpreter limits: a few thousand nested recursive calls, and generic functions whose type argument must be inferred from a
  function-literal argument.

You get per-test got/expected, runtime errors (with the line or function), `print`/`fmt.Println` output and a 10 s time limit;
drafts are saved per problem and language, and an accepted solution marks the problem done. Tests accept any valid answer where
LeetCode does (any order, any valid topological order, any height-balanced BST…). The ML coding snippets have a **Run** button too
(NumPy loads on demand). Java is not judged in the browser: there is no free, reliable way to run `javac` client-side, so Java
solutions stay read/copy/run-locally.

**Browse** (`#/browse`) filters every problem by level, status (todo/done), **study list** (Blind 75, NeetCode 150, NeetCode 250,
Grind 169, LeetCode 75, Top Interview 150 — optionally in the list's own order), pattern tag, topic and **company** (all time or last 6 months,
ranked by that company's ask frequency); every difficulty badge, tag and topic label on the site is a clickable filter, and filters
live in the URL so views can be bookmarked. Progress checkboxes are saved in your browser. Press `/` to search. The layout scales
fonts with screen size and shows code side-by-side with the explanation on wide screens.

**Algorithm topics:** Arrays & Hashing · Two Pointers · Sliding Window · Stack & Monotonic Stack · Binary Search · Linked List · Trees · Tries · Heap · Backtracking · Graphs · Advanced Graphs (Dijkstra, Bellman-Ford, Floyd-Warshall, MST, Tarjan, Hierholzer) · 1-D DP · 2-D DP · Greedy · Intervals · Math & Geometry · Bit Manipulation · Prefix Sums & Range Queries (difference arrays, Fenwick tree) · Design.

## Data sources

- **Company tags**: the public dataset [liquidslr/leetcode-company-wise-problems](https://github.com/liquidslr/leetcode-company-wise-problems)
  (358 companies cover 476 of the 477 problems). Frequency is the dataset's 0–100 score, relative within each company.
- **Study lists**: NeetCode's public problem data and site (Blind 75 / NeetCode 150 / NeetCode 250), the Grind 169 site, and LeetCode's public study plans.

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
cd go     && go run 0001_two_sum.go lc.go     # Go 1.22+, prints OK
```

The `Solution` class inside each file is exactly what you paste into LeetCode. Python files import `python/lc.py`
(ListNode / TreeNode + `make_list`, `list_vals`, `make_tree`, `tree_vals`, `find_node`); Java files inline the same helpers
(`ListNode.of(1, 2, 3)`, `TreeNode.of(3, 9, 20, null, null, 15, 7)`). Go files are LeetCode-style functions plus a `main`
holding the tests; they share `go/lc.go` (`makeList(1, 2, 3)`, `makeTree(3, 9, 20, nil, nil, 15, 7)`, `byteGrid("1.1")` and the
`assertEq` / `assertUnordered` / `assertUnorderedDeep` checks), which the browser judge interprets too.

## Repo layout

```
content/                 <- source of truth (markdown, one file per topic)
  dsa/NN-topic.md            problems: idea + ```viz + ```python / ```java / ```go (each with a "# test" / "// test" part)
  system-design/*.md         articles (one "## " heading per article)
  ml/*.md                    articles
build.py                 <- content -> data/*.json (site, incl. starter stubs) + python/*.py + java/*.java + go/*.go
tools/import_companies.py <- public company-tag dataset -> data/companies.json
tools/import_lists.py    <- public study lists -> data/lists.json
index.html, assets/      <- static single-page site (no build tools; marked, highlight.js, CodeMirror, Pyodide from CDNs)
  assets/judge.py            the in-browser judge (also runs under CPython: `python assets/judge.py` self-checks every problem)
  assets/py-worker.js        Web Worker hosting Pyodide
  assets/go-worker.js        Web Worker hosting the Go judge (go-judge.wasm.gz + Go's wasm_exec.js)
tools/gojudge/           <- the Go judge: Yaegi + test harness; `sh build.sh` -> assets/go-judge.wasm.gz
data/, python/, java/, go/ <- generated by build.py (committed so GitHub Pages and local runs just work; go/lc.go is hand-written)
```

## Edit or add content

1. Edit `content/dsa/<topic>.md` (the problem format is documented at the top of `build.py`).
2. `python build.py --test` regenerates everything and runs **all Python, Java and Go test files** plus the 25 ML NumPy snippets
   (skipped if numpy isn't installed) in parallel. Java needs `java` on PATH or `JAVA_HOME`, Go needs `go` on PATH or `GOROOT`;
   a missing toolchain is skipped. Add problem numbers to test only those: `python build.py --test 1 146`.
3. Check the browser judges: every reference solution must pass its tests and every starter stub must fail.
   `python assets/judge.py` (Python) and `cd tools/gojudge && go run .` (Go; `go run . -only 1,146` for some problems).
   After changing `tools/gojudge` (its own tests: `go test`) or upgrading Yaegi (vendored with a bug fix in
   `tools/gojudge/third_party/yaegi`), rebuild the WebAssembly with `sh tools/gojudge/build.sh`.
4. Preview: `python -m http.server` and open http://localhost:8000.
5. Commit and push to `main`; GitHub Pages serves the repo root.

## License

Code (solutions, build scripts, the site and judges, code blocks in `content/`) is [MIT](LICENSE). The writing in `content/`
(explanations and articles) is [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Third-party pieces keep their own
licenses: the Go WebAssembly support (`assets/wasm_exec.js`, BSD 3-Clause), Yaegi inside `assets/go-judge.wasm.gz` (Apache 2.0)
and the data derived from the sources above. See [LICENSE](LICENSE) for details.

Not affiliated with or endorsed by LeetCode. Problem statements belong to LeetCode; this site links to them and contains
original explanations and solutions.
