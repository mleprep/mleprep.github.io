#!/usr/bin/env python3
"""Generate the site data and runnable solution files from content/**/*.md.

    python build.py            # write data/, python/, java/
    python build.py --test     # ...and run every Python + Java test (Java needs `java` 11+ on PATH or JAVA_HOME)
    python build.py --test 146 # ...only run problems whose number is listed

Problem format (content/dsa/*.md):

    ## 1. Two Sum | Easy            (optional: "| custom-slug" when the LeetCode slug differs)
    related: 15, 167
    tags: blind75, hash-map

    Markdown explanation, ```viz blocks for ASCII diagrams, **Complexity** line...

    ```python
    class Solution: ...
    # test
    assert Solution().twoSum([2, 7, 11, 15], 9) == [0, 1]
    ```

    ```java
    class Solution { ... }
    // test
    check(new Solution().twoSum(new int[]{2, 7, 11, 15}, 9), new int[]{0, 1});
    ```
"""
import json
import os
import re
import shutil
import subprocess
import sys
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

ROOT = Path(__file__).resolve().parent
CONTENT = ROOT / "content"
SECTIONS = [("dsa", "Algorithms"), ("system-design", "System Design"), ("ml", "ML / Deep Learning")]
FENCE = re.compile(r"^```([\w-]*)[ \t]*\n(.*?)^```[ \t]*$", re.M | re.S)
HEAD = re.compile(r"(\d+)\.\s+(.+?)\s*\|\s*(Easy|Medium|Hard)\s*(?:\|\s*(\S+))?$")


def slugify(s):
    s = re.sub(r"[^a-z0-9 \-]", "", s.lower())
    return re.sub(r"[\s\-]+", "-", s).strip("-")


def camel(slug):
    return "".join(w[:1].upper() + w[1:] for w in slug.split("-"))


def split_test(code, marker):
    parts = re.split(rf"^{re.escape(marker)}\s*$", code, maxsplit=1, flags=re.M)
    return parts[0].rstrip() + "\n", (parts[1].strip("\n") + "\n") if len(parts) > 1 else ""


def parse_item(section, topic, head, body, src):
    lines = body.strip("\n").split("\n")
    meta = {}
    while lines and re.match(r"^(related|tags):", lines[0]):
        k, v = lines.pop(0).split(":", 1)
        meta[k] = [x.strip() for x in v.split(",") if x.strip()]
    body = "\n".join(lines)
    item = {"topic": topic, "tags": meta.get("tags", [])}
    if section == "dsa":
        m = HEAD.match(head)
        if not m:
            sys.exit(f"{src}: bad problem header: {head!r}")
        pid, title, diff, slug = m.groups()
        item.update(id=int(pid), title=title, difficulty=diff, slug=slug or slugify(title),
                    related=[int(x) for x in meta.get("related", [])])
        code = {}

        def grab(mm):
            lang = mm.group(1)
            if lang in ("python", "java") and lang not in code:
                code[lang] = mm.group(2)
                return ""
            return mm.group(0)

        body = FENCE.sub(grab, body)
        for lang in ("python", "java"):
            if lang not in code:
                sys.exit(f"{src}: problem {pid} missing ```{lang} block")
        item["py"], item["pyTest"] = split_test(code["python"], "# test")
        item["java"], item["javaTest"] = split_test(code["java"], "// test")
    else:
        item.update(title=head, slug=slugify(head), related=meta.get("related", []))
    item["md"] = re.sub(r"\n{3,}", "\n\n", body).strip()
    return item


def parse_file(section, path):
    text = path.read_text(encoding="utf-8").replace("\r\n", "\n")
    m = re.match(r"#\s+(.+)\n", text)
    if not m:
        sys.exit(f"{path}: must start with '# Topic Title'")
    topic = re.sub(r"^\d+-", "", path.stem)
    # split on "## " headings that are not inside code fences
    chunks, buf, in_fence = [], [], False
    for line in text[m.end():].split("\n"):
        if line.startswith("```"):
            in_fence = not in_fence
        if not in_fence and line.startswith("## "):
            chunks.append(("\n".join(buf), line[3:].strip()))
            buf = []
        else:
            buf.append(line)
    chunks.append(("\n".join(buf), None))
    intro = chunks[0][0].strip()
    items = [parse_item(section, topic, chunks[i][1], chunks[i + 1][0], path) for i in range(len(chunks) - 1)]
    return {"id": topic, "title": m.group(1).strip(), "intro": intro, "items": items}


# ---------------------------------------------------------------- runnable files
PY_HEADER = '''# {id}. {title} ({difficulty})
# https://leetcode.com/problems/{slug}/
# Run: python {fname}
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

'''

JAVA_MAIN = '''// {id}. {title} ({difficulty})
// https://leetcode.com/problems/{slug}/
// Run (Java 11+): java {fname}
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class {cls} {{
    public static void main(String[] args) {{
{tests}
        if (fails > 0) {{ System.out.println(fails + " test(s) FAILED"); System.exit(1); }}
        System.out.println("OK");
    }}

    static int fails = 0;

    static void check(Object got, Object expected) {{
        String g = str(got), e = str(expected);
        if (!g.equals(e)) {{ fails++; System.out.println("FAIL: got " + g + ", expected " + e); }}
    }}

    static String str(Object o) {{
        if (o == null) return "null";
        if (o instanceof int[]) return Arrays.toString((int[]) o);
        if (o instanceof long[]) return Arrays.toString((long[]) o);
        if (o instanceof double[]) return Arrays.toString((double[]) o);
        if (o instanceof char[]) return Arrays.toString((char[]) o);
        if (o instanceof boolean[]) return Arrays.toString((boolean[]) o);
        if (o instanceof Object[]) return Arrays.deepToString((Object[]) o);
        return String.valueOf(o);
    }}
}}

'''

JAVA_LIST = '''
// ---- LeetCode's ListNode + local test helpers ----
class ListNode {
    int val; ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    static ListNode of(int... vals) {
        ListNode dummy = new ListNode(), t = dummy;
        for (int v : vals) t = t.next = new ListNode(v);
        return dummy.next;
    }
    public String toString() {
        StringBuilder sb = new StringBuilder("[");
        int guard = 0;
        for (ListNode n = this; n != null && guard++ < 10000; n = n.next) sb.append(n.val).append(n.next != null ? ", " : "");
        return sb.append("]").toString();
    }
}
'''

JAVA_TREE = '''
// ---- LeetCode's TreeNode + local test helpers ----
class TreeNode {
    int val; TreeNode left, right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) { this.val = val; this.left = left; this.right = right; }
    /** Build from LeetCode level-order, e.g. TreeNode.of(3, 9, 20, null, null, 15, 7). */
    static TreeNode of(Integer... v) {
        if (v.length == 0 || v[0] == null) return null;
        TreeNode root = new TreeNode(v[0]);
        Deque<TreeNode> q = new ArrayDeque<>(List.of(root));
        int i = 1;
        while (!q.isEmpty() && i < v.length) {
            TreeNode n = q.poll();
            if (i < v.length && v[i] != null) q.add(n.left = new TreeNode(v[i])); i++;
            if (i < v.length && v[i] != null) q.add(n.right = new TreeNode(v[i])); i++;
        }
        return root;
    }
    static TreeNode find(TreeNode r, int x) {
        if (r == null || r.val == x) return r;
        TreeNode l = find(r.left, x);
        return l != null ? l : find(r.right, x);
    }
    public String toString() {
        List<String> out = new ArrayList<>();
        Deque<TreeNode> q = new LinkedList<>(); q.add(this);
        while (!q.isEmpty()) {
            TreeNode n = q.poll();
            if (n == null) { out.add("null"); continue; }
            out.add(String.valueOf(n.val)); q.add(n.left); q.add(n.right);
        }
        while (!out.isEmpty() && out.get(out.size() - 1).equals("null")) out.remove(out.size() - 1);
        return "[" + String.join(", ", out) + "]";
    }
}
'''


def indent(s, n):
    return "".join((" " * n + ln) if ln.strip() else ln for ln in s.splitlines(True))


def write_runnables(problems):
    for d in ("python", "java"):
        out = ROOT / d
        out.mkdir(exist_ok=True)
        for f in out.glob("[0-9P]*"):
            f.unlink()
    files = []
    for p in problems:
        base = f"{p['id']:04d}_{p['slug'].replace('-', '_')}"
        py_name = f"{base}.py"
        tests = indent(p["pyTest"], 4) if p["pyTest"].strip() else "    pass\n"
        py = (PY_HEADER.format(fname=py_name, **p) + p["py"] + '\n\nif __name__ == "__main__":\n'
              + tests + '    print("OK")\n')
        (ROOT / "python" / py_name).write_text(py, encoding="utf-8", newline="\n")

        cls = f"P{p['id']:04d}{camel(p['slug'])}"
        java_name = f"{cls}.java"
        src = p["java"] + p["javaTest"]
        java = JAVA_MAIN.format(fname=java_name, cls=cls, tests=indent(p["javaTest"], 8).rstrip("\n"), **p) + p["java"]
        if re.search(r"\bListNode\b", src) and not re.search(r"\bclass ListNode\b", src):
            java += JAVA_LIST
        if re.search(r"\bTreeNode\b", src) and not re.search(r"\bclass TreeNode\b", src):
            java += JAVA_TREE
        (ROOT / "java" / java_name).write_text(java, encoding="utf-8", newline="\n")
        files.append((p["id"], py_name, java_name))
    return files


# ---------------------------------------------------------------- tests
def find_java():
    home = os.environ.get("JAVA_HOME")
    if home and (Path(home) / "bin").exists():
        for exe in ("java.exe", "java"):
            if (Path(home) / "bin" / exe).exists():
                return str(Path(home) / "bin" / exe)
    return shutil.which("java")


def run(cmd, cwd):
    try:
        r = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=120)
        return r.returncode == 0 and r.stdout.strip().endswith("OK"), (r.stdout + r.stderr).strip()
    except subprocess.TimeoutExpired:
        return False, "TIMEOUT"


def run_tests(files, only):
    java = find_java()
    if not java:
        print("! java not found (set JAVA_HOME or PATH) - only running Python tests")
    jobs = []
    for pid, py, jv in files:
        if only and pid not in only:
            continue
        jobs.append((f"{py}", [sys.executable, py], ROOT / "python"))
        if java:
            jobs.append((f"{jv}", [java, jv], ROOT / "java"))
    failed = []
    with ThreadPoolExecutor(max_workers=os.cpu_count() or 4) as ex:
        for (name, _, _), (ok, out) in zip(jobs, ex.map(lambda j: run(j[1], j[2]), jobs)):
            if not ok:
                failed.append(name)
                print(f"FAIL {name}\n{indent(out[-1500:], 4)}")
    print(f"\n{len(jobs) - len(failed)}/{len(jobs)} runnable files passed")
    return not failed


# ---------------------------------------------------------------- main
def main():
    manifest, problems = {"sections": []}, []
    for sec, sec_title in SECTIONS:
        sec_dir = CONTENT / sec
        topics = [parse_file(sec, f) for f in sorted(sec_dir.glob("*.md"))] if sec_dir.exists() else []
        manifest["sections"].append({"id": sec, "title": sec_title, "topics": [
            {"id": t["id"], "title": t["title"], "items": [
                {k: it[k] for k in ("id", "title", "slug", "difficulty", "tags") if k in it} for it in t["items"]]}
            for t in topics]})
        if sec == "dsa":
            problems = [it for t in topics for it in t["items"]]
        out_dir = ROOT / "data" / sec
        if out_dir.exists():
            shutil.rmtree(out_dir)
        out_dir.mkdir(parents=True)
        for t in topics:
            (out_dir / f"{t['id']}.json").write_text(json.dumps(t, ensure_ascii=False, separators=(",", ":")),
                                                     encoding="utf-8")

    by_id = {}
    for p in problems:
        if p["id"] in by_id:
            sys.exit(f"duplicate problem {p['id']}")
        by_id[p["id"]] = p
    missing = sorted({r for p in problems for r in p["related"] if r not in by_id})
    if missing:
        print(f"! related problems not in the set (not linked): {missing}")
    manifest["counts"] = {d: sum(p["difficulty"] == d for p in problems) for d in ("Easy", "Medium", "Hard")}
    (ROOT / "data" / "manifest.json").write_text(json.dumps(manifest, ensure_ascii=False, separators=(",", ":")),
                                                 encoding="utf-8")
    files = write_runnables(problems)
    print(f"built {len(problems)} problems ({manifest['counts']}), "
          + ", ".join(f"{s['title']}: {sum(len(t['items']) for t in s['topics'])}" for s in manifest["sections"]))

    if "--test" in sys.argv:
        only = {int(a) for a in sys.argv[sys.argv.index("--test") + 1:] if a.isdigit()}
        sys.exit(0 if run_tests(files, only) else 1)


if __name__ == "__main__":
    main()
