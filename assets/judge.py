"""In-browser judge (runs inside Pyodide in assets/py-worker.js; also runs under plain CPython for testing).

judge(user_code, test_code) executes the user's code, then every top-level statement of the problem's
tests. Each `assert` (or loop containing asserts) is one test case; `assert got == expected` reports both
sides when it fails. Returns a JSON string:
  {"error": str|None, "cases": [{"src", "ok", "got"?, "expected"?, "error"?}], "stdout": str, "ms": float}
"""
import ast
import contextlib
import io
import json
import sys
import time
import traceback

PRELUDE = """
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *
"""
MAX_OUT = 20_000


def _short(v, n=300):
    s = repr(v)
    return s if len(s) <= n else s[:n] + " …"


def _where(e):
    """'line 7' of the user's code (or of the test) where exception e was raised."""
    frames = traceback.extract_tb(e.__traceback__)
    for fname, label in (("solution.py", "your code"), ("<test>", "test")):
        mine = [f for f in frames if f.filename == fname]
        if mine:
            return f" (line {mine[-1].lineno} of {label})"
    return ""


def _err(e):
    return f"{type(e).__name__}: {e}{_where(e)}"


def _has_assert(node):
    return any(isinstance(n, ast.Assert) for n in ast.walk(node))


def judge(user_code, test_code):
    out, cases, error = io.StringIO(), [], None
    g = {"__name__": "__solution__"}
    t0 = time.perf_counter()
    random_state = __import__("random").getstate()
    __import__("random").seed(0)
    with contextlib.redirect_stdout(out):
        try:
            exec(PRELUDE, g)
            exec(compile(user_code, "solution.py", "exec"), g)
        except SyntaxError as e:
            error = f"SyntaxError: {e.msg} (line {e.lineno} of your code)"
        except BaseException as e:  # noqa: BLE001 - report anything the user's code raises
            error = _err(e)
        if error is None:
            try:
                tree = ast.parse(test_code, "<test>")
            except SyntaxError as e:
                tree, error = None, f"test parse error: {e}"
            for node in tree.body if tree else []:
                src = ast.get_source_segment(test_code, node) or ""
                is_case = _has_assert(node)
                try:
                    if isinstance(node, ast.Assert) and isinstance(node.test, ast.Compare) \
                            and len(node.test.ops) == 1 and isinstance(node.test.ops[0], ast.Eq):
                        lhs = eval(compile(ast.Expression(node.test.left), "<test>", "eval"), g)
                        rhs = eval(compile(ast.Expression(node.test.comparators[0]), "<test>", "eval"), g)
                        ok = lhs == rhs
                        case = {"src": src, "ok": bool(ok)}
                        if not ok:
                            case.update(got=_short(lhs), expected=_short(rhs))
                        cases.append(case)
                    else:
                        exec(compile(ast.Module([node], []), "<test>", "exec"), g)
                        if is_case:
                            cases.append({"src": src, "ok": True})
                except AssertionError as e:
                    cases.append({"src": src, "ok": False, "error": str(e) or "assertion failed"})
                except BaseException as e:  # noqa: BLE001 - the user's code raised (runtime error)
                    cases.append({"src": src, "ok": False, "error": _err(e), "exc": True})
                    if not is_case:  # a setup line failed: later cases depend on it
                        break
    __import__("random").setstate(random_state)
    text = out.getvalue()
    if len(text) > MAX_OUT:
        text = text[:MAX_OUT] + f"\n… ({len(text) - MAX_OUT} more characters)"
    return json.dumps({"error": error, "cases": cases, "stdout": text, "ms": round((time.perf_counter() - t0) * 1000, 1)})


def run_script(code):
    """Run a standalone snippet (ML coding articles); returns JSON {"error", "stdout", "ms"}."""
    out, error = io.StringIO(), None
    t0 = time.perf_counter()
    with contextlib.redirect_stdout(out):
        try:
            exec(compile(code, "solution.py", "exec"), {"__name__": "__main__"})
        except BaseException as e:  # noqa: BLE001
            error = _err(e)
    return json.dumps({"error": error, "stdout": out.getvalue()[:MAX_OUT], "ms": round((time.perf_counter() - t0) * 1000, 1)})


if __name__ == "__main__":  # python assets/judge.py data/dsa/*.json  -> self-check against reference solutions
    import glob
    sys.path.insert(0, "python")
    sys.setrecursionlimit(10_000)
    bad = 0
    for f in sys.argv[1:] or glob.glob("data/dsa/*.json"):
        for it in json.load(open(f, encoding="utf-8"))["items"]:
            ref = json.loads(judge(it["py"], it["pyTest"]))
            stub = json.loads(judge(it["pyStub"], it["pyTest"]))
            ok_ref = ref["error"] is None and ref["cases"] and all(c["ok"] for c in ref["cases"])
            ok_stub = stub["error"] is None and stub["cases"] and not all(c["ok"] for c in stub["cases"])
            if not ok_ref or not ok_stub:
                bad += 1
                print(f"{it['id']} {it['title']}: reference {'ok' if ok_ref else ref}; stub {'ok' if ok_stub else stub}")
    print("judge self-check:", "all good" if not bad else f"{bad} problem(s) need attention")
    sys.exit(1 if bad else 0)
