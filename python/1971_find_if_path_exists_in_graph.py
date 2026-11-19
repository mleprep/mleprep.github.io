# 1971. Find if Path Exists in Graph (Easy)
# https://leetcode.com/problems/find-if-path-exists-in-graph/
# Run: python 1971_find_if_path_exists_in_graph.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for a, b in edges:
            parent[find(a)] = find(b)
        return find(source) == find(destination)


if __name__ == "__main__":
    assert Solution().validPath(3, [[0, 1], [1, 2], [2, 0]], 0, 2) is True
    assert Solution().validPath(6, [[0, 1], [0, 2], [3, 5], [5, 4], [4, 3]], 0, 5) is False
    print("OK")
