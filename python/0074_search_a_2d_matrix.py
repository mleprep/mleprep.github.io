# 74. Search a 2D Matrix (Medium)
# https://leetcode.com/problems/search-a-2d-matrix/
# Run: python 0074_search_a_2d_matrix.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        lo, hi = 0, m * n - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            v = matrix[mid // n][mid % n]
            if v == target:
                return True
            if v < target:
                lo = mid + 1
            else:
                hi = mid - 1
        return False


if __name__ == "__main__":
    M = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
    assert Solution().searchMatrix(M, 3) is True
    assert Solution().searchMatrix(M, 13) is False
    print("OK")
