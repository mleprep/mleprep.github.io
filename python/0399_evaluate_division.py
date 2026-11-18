# 399. Evaluate Division (Medium)
# https://leetcode.com/problems/evaluate-division/
# Run: python 0399_evaluate_division.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        g = defaultdict(dict)
        for (a, b), v in zip(equations, values):
            g[a][b] = v
            g[b][a] = 1 / v

        def dfs(x, y, seen):
            if x == y:
                return 1.0
            seen.add(x)
            for nb, w in g[x].items():
                if nb not in seen:
                    r = dfs(nb, y, seen)
                    if r != -1:
                        return w * r
            return -1.0

        return [dfs(x, y, set()) if x in g and y in g else -1.0 for x, y in queries]


if __name__ == "__main__":
    eq, val = [["a", "b"], ["b", "c"]], [2.0, 3.0]
    qs = [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]
    assert Solution().calcEquation(eq, val, qs) == [6.0, 0.5, -1.0, 1.0, -1.0]
    print("OK")
