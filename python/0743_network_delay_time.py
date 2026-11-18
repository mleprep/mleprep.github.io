# 743. Network Delay Time (Medium)
# https://leetcode.com/problems/network-delay-time/
# Run: python 0743_network_delay_time.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        g = defaultdict(list)
        for u, v, w in times:
            g[u].append((v, w))
        dist = {}
        pq = [(0, k)]
        while pq:
            d, u = heappop(pq)
            if u in dist:
                continue
            dist[u] = d
            for v, w in g[u]:
                if v not in dist:
                    heappush(pq, (d + w, v))
        return max(dist.values()) if len(dist) == n else -1


if __name__ == "__main__":
    assert Solution().networkDelayTime([[2, 1, 1], [2, 3, 1], [3, 4, 1]], 4, 2) == 2
    assert Solution().networkDelayTime([[1, 2, 1]], 2, 2) == -1
    print("OK")
