# 63. Unique Paths II (Medium)
# https://leetcode.com/problems/unique-paths-ii/
# Run: python 0063_unique_paths_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        n = len(obstacleGrid[0])
        row = [0] * n
        row[0] = 1
        for r in obstacleGrid:
            for c in range(n):
                if r[c]:
                    row[c] = 0
                elif c:
                    row[c] += row[c - 1]
        return row[-1]


if __name__ == "__main__":
    assert Solution().uniquePathsWithObstacles([[0, 0, 0], [0, 1, 0], [0, 0, 0]]) == 2
    assert Solution().uniquePathsWithObstacles([[0, 1], [0, 0]]) == 1
    print("OK")
