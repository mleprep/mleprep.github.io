# 59. Spiral Matrix II (Medium)
# https://leetcode.com/problems/spiral-matrix-ii/
# Run: python 0059_spiral_matrix_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        m = [[0] * n for _ in range(n)]
        top, bottom, left, right, k = 0, n - 1, 0, n - 1, 1
        while top <= bottom:
            for c in range(left, right + 1):
                m[top][c] = k; k += 1
            for r in range(top + 1, bottom + 1):
                m[r][right] = k; k += 1
            for c in range(right - 1, left - 1, -1):
                if top < bottom:
                    m[bottom][c] = k; k += 1
            for r in range(bottom - 1, top, -1):
                if left < right:
                    m[r][left] = k; k += 1
            top, bottom, left, right = top + 1, bottom - 1, left + 1, right - 1
        return m


if __name__ == "__main__":
    assert Solution().generateMatrix(3) == [[1, 2, 3], [8, 9, 4], [7, 6, 5]]
    assert Solution().generateMatrix(1) == [[1]]
    print("OK")
