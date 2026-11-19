# 815. Bus Routes (Hard)
# https://leetcode.com/problems/bus-routes/
# Run: python 0815_bus_routes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numBusesToDestination(self, routes: List[List[int]], source: int, target: int) -> int:
        if source == target:
            return 0
        at = defaultdict(list)                        # stop -> routes through it
        for i, route in enumerate(routes):
            for s in route:
                at[s].append(i)
        used, seen, q, buses = set(), {source}, deque([source]), 0
        while q:
            buses += 1
            for _ in range(len(q)):
                for r in at[q.popleft()]:
                    if r in used:
                        continue
                    used.add(r)
                    for s in routes[r]:
                        if s == target:
                            return buses
                        if s not in seen:
                            seen.add(s)
                            q.append(s)
        return -1


if __name__ == "__main__":
    assert Solution().numBusesToDestination([[1, 2, 7], [3, 6, 7]], 1, 6) == 2
    assert Solution().numBusesToDestination([[7, 12], [4, 5, 15], [6], [15, 19], [9, 12, 13]], 15, 12) == -1
    assert Solution().numBusesToDestination([[1, 2]], 1, 1) == 0
    assert Solution().numBusesToDestination([[1, 2]], 3, 1) == -1
    assert Solution().numBusesToDestination([[1, 7], [3, 5]], 5, 5) == 0
    print("OK")
