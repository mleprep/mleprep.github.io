# 118. Pascal's Triangle (Easy)
# https://leetcode.com/problems/pascals-triangle/
# Run: python 0118_pascals_triangle.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        rows = [[1]]
        for _ in range(numRows - 1):
            prev = rows[-1]
            rows.append([1] + [prev[i] + prev[i + 1] for i in range(len(prev) - 1)] + [1])
        return rows


if __name__ == "__main__":
    assert Solution().generate(5) == [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
    print("OK")
