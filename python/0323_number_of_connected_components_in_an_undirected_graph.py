# 323. Number of Connected Components in an Undirected Graph (Medium)
# https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
# Run: python 0323_number_of_connected_components_in_an_undirected_graph.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        comps = n
        for a, b in edges:
            ra, rb = find(a), find(b)
            if ra != rb:
                parent[ra] = rb
                comps -= 1
        return comps


if __name__ == "__main__":
    assert Solution().countComponents(5, [[0, 1], [1, 2], [3, 4]]) == 2
    assert Solution().countComponents(5, [[0, 1], [1, 2], [2, 3], [3, 4]]) == 1
    print("OK")
