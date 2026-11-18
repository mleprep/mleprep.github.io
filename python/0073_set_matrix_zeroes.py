# 73. Set Matrix Zeroes (Medium)
# https://leetcode.com/problems/set-matrix-zeroes/
# Run: python 0073_set_matrix_zeroes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        R, C = len(matrix), len(matrix[0])
        row0 = any(matrix[0][c] == 0 for c in range(C))
        col0 = any(matrix[r][0] == 0 for r in range(R))
        for r in range(1, R):
            for c in range(1, C):
                if matrix[r][c] == 0:
                    matrix[r][0] = matrix[0][c] = 0
        for r in range(1, R):
            for c in range(1, C):
                if matrix[r][0] == 0 or matrix[0][c] == 0:
                    matrix[r][c] = 0
        if row0:
            for c in range(C):
                matrix[0][c] = 0
        if col0:
            for r in range(R):
                matrix[r][0] = 0


if __name__ == "__main__":
    m = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]; Solution().setZeroes(m)
    assert m == [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]
    print("OK")
