# 1631. Path With Minimum Effort (Medium)
# https://leetcode.com/problems/path-with-minimum-effort/
# Run: python 1631_path_with_minimum_effort.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        R, C = len(heights), len(heights[0])
        best = [[float("inf")] * C for _ in range(R)]
        best[0][0] = 0
        pq = [(0, 0, 0)]
        while pq:
            e, r, c = heappop(pq)
            if (r, c) == (R - 1, C - 1):
                return e
            if e > best[r][c]:
                continue
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C:
                    ne = max(e, abs(heights[nr][nc] - heights[r][c]))
                    if ne < best[nr][nc]:
                        best[nr][nc] = ne
                        heappush(pq, (ne, nr, nc))


if __name__ == "__main__":
    assert Solution().minimumEffortPath([[1, 2, 2], [3, 8, 2], [5, 3, 5]]) == 2
    assert Solution().minimumEffortPath([[1, 2, 1, 1, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 1, 1, 2, 1]]) == 0
    print("OK")
