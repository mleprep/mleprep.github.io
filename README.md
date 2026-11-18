# leetcode_review

A concise, browsable interview-prep site: **https://mleprep.github.io/**

- **Algorithms**: ~250 hand-picked problems (Blind 75, NeetCode 150, Grind 169, LeetCode Top Interview 150), grouped by pattern. Each one has the key idea, an ASCII sketch, complexity, **Python + Java** solutions with local tests, and links to related problems.
- **System Design**: fundamentals and classic designs, with architecture diagrams.
- **ML / Deep Learning**: core interview Q&A, NumPy implementations from scratch, and ML system design.

Progress checkboxes are saved in your browser (localStorage). Press `/` to search.

## Run solutions locally

Every problem is also generated as a standalone runnable file:

```bash
cd python && python 0001_two_sum.py          # prints OK when all asserts pass
cd java   && java P0001TwoSum.java            # Java 11+ single-file launch, prints OK
```

The class inside each file (`class Solution` / `class Solution {...}`) is exactly what you paste into LeetCode.
Python files import `python/lc.py` (ListNode / TreeNode + `make_list`, `make_tree`, ...). Java files inline the same helpers.

## Repo layout

```
content/            <- source of truth (markdown, one file per topic)
  dsa/NN-topic.md       problems: idea + viz + ```python / ```java (with "# test" / "// test" sections)
  system-design/*.md    articles
  ml/*.md               articles
build.py            <- content -> data/*.json (site) + python/*.py + java/*.java (runnable)
index.html, assets/ <- static single-page site (no build tools; marked + highlight.js from CDN)
```

## Edit / add a problem

1. Edit `content/dsa/<topic>.md` (format is documented at the top of `build.py`).
2. `python build.py --test` regenerates everything and runs all Python + Java tests (Java needs `java` on PATH or `JAVA_HOME`).
3. Preview: `python -m http.server` then open http://localhost:8000.
4. Commit and push; GitHub Pages serves `main` (root).
