# 329. Longest Increasing Path in a Matrix (Hard)
# https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
# Run: python 0329_longest_increasing_path_in_a_matrix.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        R, C = len(matrix), len(matrix[0])

        @cache
        def best(r, c):
            v, res = matrix[r][c], 1
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and matrix[nr][nc] > v:
                    res = max(res, 1 + best(nr, nc))
            return res

        return max(best(r, c) for r in range(R) for c in range(C))


if __name__ == "__main__":
    assert Solution().longestIncreasingPath([[9, 9, 4], [6, 6, 8], [2, 1, 1]]) == 4
    assert Solution().longestIncreasingPath([[3, 4, 5], [3, 2, 6], [2, 2, 1]]) == 4
    print("OK")
