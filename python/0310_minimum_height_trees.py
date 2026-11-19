# 310. Minimum Height Trees (Medium)
# https://leetcode.com/problems/minimum-height-trees/
# Run: python 0310_minimum_height_trees.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if n <= 2:
            return list(range(n))
        g = defaultdict(set)
        for a, b in edges:
            g[a].add(b)
            g[b].add(a)
        leaves = [i for i in range(n) if len(g[i]) == 1]
        remaining = n
        while remaining > 2:
            remaining -= len(leaves)
            nxt = []
            for leaf in leaves:
                nb = g[leaf].pop()
                g[nb].remove(leaf)
                if len(g[nb]) == 1:
                    nxt.append(nb)
            leaves = nxt
        return sorted(leaves)


if __name__ == "__main__":
    assert sorted(Solution().findMinHeightTrees(4, [[1, 0], [1, 2], [1, 3]])) == sorted([1])
    assert sorted(Solution().findMinHeightTrees(6, [[3, 0], [3, 1], [3, 2], [3, 4], [5, 4]])) == sorted([3, 4])
    print("OK")
