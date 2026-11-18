# 133. Clone Graph (Medium)
# https://leetcode.com/problems/clone-graph/
# Run: python 0133_clone_graph.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Node:  # provided by LeetCode
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        clones = {}

        def dfs(n):
            if n in clones:
                return clones[n]
            c = clones[n] = Node(n.val)
            c.neighbors = [dfs(nb) for nb in n.neighbors]
            return c

        return dfs(node) if node else None


if __name__ == "__main__":
    ns = [Node(i) for i in range(1, 5)]
    for i, nbs in enumerate([[2, 4], [1, 3], [2, 4], [1, 3]]):
        ns[i].neighbors = [ns[j - 1] for j in nbs]
    c = Solution().cloneGraph(ns[0])
    assert c is not ns[0] and [n.val for n in c.neighbors] == [2, 4]
    assert c.neighbors[0].neighbors[0] is c        # cycle preserved
    print("OK")
