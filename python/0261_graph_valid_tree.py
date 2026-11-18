# 261. Graph Valid Tree (Medium)
# https://leetcode.com/problems/graph-valid-tree/
# Run: python 0261_graph_valid_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        if len(edges) != n - 1:
            return False
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for a, b in edges:
            ra, rb = find(a), find(b)
            if ra == rb:
                return False
            parent[ra] = rb
        return True


if __name__ == "__main__":
    assert Solution().validTree(5, [[0, 1], [0, 2], [0, 3], [1, 4]]) is True
    assert Solution().validTree(5, [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]) is False
    print("OK")
