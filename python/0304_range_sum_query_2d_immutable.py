# 304. Range Sum Query 2D - Immutable (Medium)
# https://leetcode.com/problems/range-sum-query-2d-immutable/
# Run: python 0304_range_sum_query_2d_immutable.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class NumMatrix:
    def __init__(self, matrix: List[List[int]]):
        R, C = len(matrix), len(matrix[0])
        self.P = [[0] * (C + 1) for _ in range(R + 1)]
        for r in range(R):
            for c in range(C):
                self.P[r + 1][c + 1] = matrix[r][c] + self.P[r][c + 1] + self.P[r + 1][c] - self.P[r][c]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        P = self.P
        return P[row2 + 1][col2 + 1] - P[row1][col2 + 1] - P[row2 + 1][col1] + P[row1][col1]


if __name__ == "__main__":
    nm = NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]])
    assert [nm.sumRegion(2, 1, 4, 3), nm.sumRegion(1, 1, 2, 2), nm.sumRegion(1, 2, 2, 4)] == [8, 11, 12]
    print("OK")
