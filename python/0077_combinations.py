# 77. Combinations (Medium)
# https://leetcode.com/problems/combinations/
# Run: python 0077_combinations.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res, path = [], []

        def bt(start):
            if len(path) == k:
                res.append(path[:])
                return
            for i in range(start, n - (k - len(path)) + 2):
                path.append(i)
                bt(i + 1)
                path.pop()

        bt(1)
        return res


if __name__ == "__main__":
    assert Solution().combine(4, 2) == [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]
    assert Solution().combine(1, 1) == [[1]]
    print("OK")
