# 54. Spiral Matrix (Medium)
# https://leetcode.com/problems/spiral-matrix/
# Run: python 0054_spiral_matrix.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        res = []
        top, bottom, left, right = 0, len(matrix) - 1, 0, len(matrix[0]) - 1
        while top <= bottom and left <= right:
            res += [matrix[top][c] for c in range(left, right + 1)]
            res += [matrix[r][right] for r in range(top + 1, bottom + 1)]
            if top < bottom and left < right:
                res += [matrix[bottom][c] for c in range(right - 1, left - 1, -1)]
                res += [matrix[r][left] for r in range(bottom - 1, top, -1)]
            top, bottom, left, right = top + 1, bottom - 1, left + 1, right - 1
        return res


if __name__ == "__main__":
    assert Solution().spiralOrder([[1, 2, 3], [4, 5, 6], [7, 8, 9]]) == [1, 2, 3, 6, 9, 8, 7, 4, 5]
    assert Solution().spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]) == [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
    print("OK")
