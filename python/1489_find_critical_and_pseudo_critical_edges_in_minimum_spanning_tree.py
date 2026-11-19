# 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree (Hard)
# https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
# Run: python 1489_find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findCriticalAndPseudoCriticalEdges(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        order = sorted(range(len(edges)), key=lambda i: edges[i][2])

        def mst(skip=-1, force=-1):
            parent = list(range(n))

            def find(x):
                while parent[x] != x:
                    parent[x] = parent[parent[x]]
                    x = parent[x]
                return x

            weight = used = 0
            if force >= 0:
                a, b, w = edges[force]
                parent[find(a)] = find(b)
                weight, used = w, 1
            for i in order:
                a, b, w = edges[i]
                if i != skip and find(a) != find(b):
                    parent[find(a)] = find(b)
                    weight, used = weight + w, used + 1
            return weight if used == n - 1 else float("inf")

        base = mst()
        critical, pseudo = [], []
        for i in range(len(edges)):
            if mst(skip=i) > base:
                critical.append(i)
            elif mst(force=i) == base:
                pseudo.append(i)
        return [critical, pseudo]


if __name__ == "__main__":
    r = Solution().findCriticalAndPseudoCriticalEdges(5, [[0, 1, 1], [1, 2, 1], [2, 3, 2], [0, 3, 2], [0, 4, 3], [3, 4, 3], [1, 4, 6]])
    assert [sorted(r[0]), sorted(r[1])] == [[0, 1], [2, 3, 4, 5]]
    r = Solution().findCriticalAndPseudoCriticalEdges(4, [[0, 1, 1], [1, 2, 1], [2, 3, 1], [0, 3, 1]])
    assert [sorted(r[0]), sorted(r[1])] == [[], [0, 1, 2, 3]]
    r = Solution().findCriticalAndPseudoCriticalEdges(3, [[0, 1, 1], [1, 2, 2], [0, 2, 3]])
    assert [sorted(r[0]), sorted(r[1])] == [[0, 1], []]
    print("OK")
