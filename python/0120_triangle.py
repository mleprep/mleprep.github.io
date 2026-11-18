# 120. Triangle (Medium)
# https://leetcode.com/problems/triangle/
# Run: python 0120_triangle.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        best = triangle[-1][:]
        for r in range(len(triangle) - 2, -1, -1):
            for c in range(r + 1):
                best[c] = triangle[r][c] + min(best[c], best[c + 1])
        return best[0]


if __name__ == "__main__":
    assert Solution().minimumTotal([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]) == 11
    assert Solution().minimumTotal([[-10]]) == -10
    print("OK")
