# 96. Unique Binary Search Trees (Medium)
# https://leetcode.com/problems/unique-binary-search-trees/
# Run: python 0096_unique_binary_search_trees.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numTrees(self, n: int) -> int:
        G = [1] + [0] * n
        for m in range(1, n + 1):
            G[m] = sum(G[r - 1] * G[m - r] for r in range(1, m + 1))
        return G[n]


if __name__ == "__main__":
    assert Solution().numTrees(3) == 5
    assert Solution().numTrees(1) == 1
    assert Solution().numTrees(19) == 1767263190
    print("OK")
