# 785. Is Graph Bipartite? (Medium)
# https://leetcode.com/problems/is-graph-bipartite/
# Run: python 0785_is_graph_bipartite.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isBipartite(self, graph: List[List[int]]) -> bool:
        color = [-1] * len(graph)
        for s in range(len(graph)):
            if color[s] != -1:
                continue
            color[s] = 0
            q = deque([s])
            while q:
                u = q.popleft()
                for v in graph[u]:
                    if color[v] == -1:
                        color[v] = 1 - color[u]
                        q.append(v)
                    elif color[v] == color[u]:
                        return False
        return True


if __name__ == "__main__":
    assert Solution().isBipartite([[1, 2, 3], [0, 2], [0, 1, 3], [0, 2]]) is False
    assert Solution().isBipartite([[1, 3], [0, 2], [1, 3], [0, 2]]) is True
    print("OK")
