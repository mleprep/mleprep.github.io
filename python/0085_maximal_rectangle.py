# 85. Maximal Rectangle (Hard)
# https://leetcode.com/problems/maximal-rectangle/
# Run: python 0085_maximal_rectangle.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix:
            return 0
        h = [0] * len(matrix[0])
        best = 0
        for row in matrix:
            for c, v in enumerate(row):
                h[c] = h[c] + 1 if v == "1" else 0
            st = []
            for i, x in enumerate(h + [0]):
                while st and h[st[-1]] >= x:
                    top = h[st.pop()]
                    best = max(best, top * (i - (st[-1] if st else -1) - 1))
                st.append(i)
        return best


if __name__ == "__main__":
    M = ["10100", "10111", "11111", "10010"]
    assert Solution().maximalRectangle([list(r) for r in M]) == 6
    assert Solution().maximalRectangle([["0"]]) == 0
    assert Solution().maximalRectangle([["1"]]) == 1
    assert Solution().maximalRectangle([list(r) for r in ["0110", "1111", "1110"]]) == 6
    print("OK")
