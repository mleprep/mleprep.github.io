# 931. Minimum Falling Path Sum (Medium)
# https://leetcode.com/problems/minimum-falling-path-sum/
# Run: python 0931_minimum_falling_path_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        prev = matrix[0][:]
        n = len(prev)
        for row in matrix[1:]:
            prev = [row[c] + min(prev[max(0, c - 1):c + 2]) for c in range(n)]
        return min(prev)


if __name__ == "__main__":
    assert Solution().minFallingPathSum([[2, 1, 3], [6, 5, 4], [7, 8, 9]]) == 13
    assert Solution().minFallingPathSum([[-19, 57], [-40, -5]]) == -59
    print("OK")
