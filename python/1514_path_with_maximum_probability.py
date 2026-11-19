# 1514. Path with Maximum Probability (Medium)
# https://leetcode.com/problems/path-with-maximum-probability/
# Run: python 1514_path_with_maximum_probability.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProbability(self, n: int, edges: List[List[int]], succProb: List[float], start_node: int, end_node: int) -> float:
        g = defaultdict(list)
        for (a, b), p in zip(edges, succProb):
            g[a].append((b, p))
            g[b].append((a, p))
        best = [0.0] * n
        best[start_node] = 1.0
        pq = [(-1.0, start_node)]
        while pq:
            p, u = heappop(pq)
            p = -p
            if u == end_node:
                return p
            if p < best[u]:
                continue
            for v, w in g[u]:
                if p * w > best[v]:
                    best[v] = p * w
                    heappush(pq, (-best[v], v))
        return 0.0


if __name__ == "__main__":
    assert abs(Solution().maxProbability(3, [[0, 1], [1, 2], [0, 2]], [0.5, 0.5, 0.2], 0, 2) - 0.25) < 1e-9
    assert Solution().maxProbability(3, [[0, 1]], [0.5], 0, 2) == 0.0
    print("OK")
