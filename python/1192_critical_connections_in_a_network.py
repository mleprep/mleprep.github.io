# 1192. Critical Connections in a Network (Hard)
# https://leetcode.com/problems/critical-connections-in-a-network/
# Run: python 1192_critical_connections_in_a_network.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def criticalConnections(self, n: int, connections: List[List[int]]) -> List[List[int]]:
        g = defaultdict(list)
        for a, b in connections:
            g[a].append(b)
            g[b].append(a)
        tin, low, res, timer = [-1] * n, [0] * n, [], [0]

        def dfs(u, parent):
            tin[u] = low[u] = timer[0]
            timer[0] += 1
            for v in g[u]:
                if v == parent:
                    continue
                if tin[v] == -1:
                    dfs(v, u)
                    low[u] = min(low[u], low[v])
                    if low[v] > tin[u]:
                        res.append([u, v])
                else:
                    low[u] = min(low[u], tin[v])

        dfs(0, -1)
        return res


if __name__ == "__main__":
    assert [sorted(e) for e in Solution().criticalConnections(4, [[0, 1], [1, 2], [2, 0], [1, 3]])] == [[1, 3]]
    print("OK")
