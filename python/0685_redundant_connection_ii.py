# 685. Redundant Connection II (Hard)
# https://leetcode.com/problems/redundant-connection-ii/
# Run: python 0685_redundant_connection_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findRedundantDirectedConnection(self, edges: List[List[int]]) -> List[int]:
        n = len(edges)
        parent_of, cand1, cand2 = {}, None, None
        for u, v in edges:
            if v in parent_of:
                cand1, cand2 = [parent_of[v], v], [u, v]
            else:
                parent_of[v] = u
        root = list(range(n + 1))

        def find(x):
            while root[x] != x:
                root[x] = root[root[x]]
                x = root[x]
            return x

        for u, v in edges:
            if [u, v] == cand2:
                continue
            ru, rv = find(u), find(v)
            if ru == rv:                       # cycle found
                return cand1 if cand1 else [u, v]
            root[rv] = ru
        return cand2


if __name__ == "__main__":
    assert Solution().findRedundantDirectedConnection([[1, 2], [1, 3], [2, 3]]) == [2, 3]
    assert Solution().findRedundantDirectedConnection([[1, 2], [2, 3], [3, 4], [4, 1], [1, 5]]) == [4, 1]
    assert Solution().findRedundantDirectedConnection([[2, 1], [3, 1], [4, 2], [1, 4]]) == [2, 1]
    print("OK")
