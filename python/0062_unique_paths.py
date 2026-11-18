# 62. Unique Paths (Medium)
# https://leetcode.com/problems/unique-paths/
# Run: python 0062_unique_paths.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        row = [1] * n
        for _ in range(1, m):
            for c in range(1, n):
                row[c] += row[c - 1]
        return row[-1]


if __name__ == "__main__":
    assert Solution().uniquePaths(3, 7) == 28
    assert Solution().uniquePaths(3, 2) == 3
    print("OK")
