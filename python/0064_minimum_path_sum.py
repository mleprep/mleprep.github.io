# 64. Minimum Path Sum (Medium)
# https://leetcode.com/problems/minimum-path-sum/
# Run: python 0064_minimum_path_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        n = len(grid[0])
        row = [float("inf")] * n
        row[0] = 0
        for r in grid:
            row[0] += r[0]
            for c in range(1, n):
                row[c] = r[c] + min(row[c], row[c - 1])
        return row[-1]


if __name__ == "__main__":
    assert Solution().minPathSum([[1, 3, 1], [1, 5, 1], [4, 2, 1]]) == 7
    assert Solution().minPathSum([[1, 2, 3], [4, 5, 6]]) == 12
    print("OK")
