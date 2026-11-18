# 1584. Min Cost to Connect All Points (Medium)
# https://leetcode.com/problems/min-cost-to-connect-all-points/
# Run: python 1584_min_cost_to_connect_all_points.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        INF = float("inf")
        dist = [INF] * n
        dist[0] = 0
        used = [False] * n
        total = 0
        for _ in range(n):
            u = min((d, i) for i, d in enumerate(dist) if not used[i])[1]
            used[u] = True
            total += dist[u]
            ux, uy = points[u]
            for v in range(n):
                if not used[v]:
                    dist[v] = min(dist[v], abs(points[v][0] - ux) + abs(points[v][1] - uy))
        return total


if __name__ == "__main__":
    assert Solution().minCostConnectPoints([[0, 0], [2, 2], [3, 10], [5, 2], [7, 0]]) == 20
    assert Solution().minCostConnectPoints([[3, 12], [-2, 5], [-4, 1]]) == 18
    print("OK")
