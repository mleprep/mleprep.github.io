# 1976. Number of Ways to Arrive at Destination (Medium)
# https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
# Run: python 1976_number_of_ways_to_arrive_at_destination.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countPaths(self, n: int, roads: List[List[int]]) -> int:
        MOD = 10 ** 9 + 7
        g = defaultdict(list)
        for a, b, t in roads:
            g[a].append((b, t))
            g[b].append((a, t))
        dist, ways = [float("inf")] * n, [0] * n
        dist[0], ways[0] = 0, 1
        pq = [(0, 0)]
        while pq:
            d, u = heappop(pq)
            if d > dist[u]:
                continue
            for v, t in g[u]:
                if d + t < dist[v]:
                    dist[v], ways[v] = d + t, ways[u]
                    heappush(pq, (dist[v], v))
                elif d + t == dist[v]:
                    ways[v] = (ways[v] + ways[u]) % MOD
        return ways[n - 1]


if __name__ == "__main__":
    R = [[0, 6, 7], [0, 1, 2], [1, 2, 3], [1, 3, 3], [6, 3, 3], [3, 5, 1], [6, 5, 1], [2, 5, 1], [0, 4, 5], [4, 6, 2]]
    assert Solution().countPaths(7, R) == 4
    assert Solution().countPaths(2, [[1, 0, 10]]) == 1
    print("OK")
