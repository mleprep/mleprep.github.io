# 1029. Two City Scheduling (Medium)
# https://leetcode.com/problems/two-city-scheduling/
# Run: python 1029_two_city_scheduling.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def twoCitySchedCost(self, costs: List[List[int]]) -> int:
        costs.sort(key=lambda c: c[0] - c[1])
        n = len(costs) // 2
        return sum(a for a, _ in costs[:n]) + sum(b for _, b in costs[n:])


if __name__ == "__main__":
    assert Solution().twoCitySchedCost([[10, 20], [30, 200], [400, 50], [30, 20]]) == 110
    assert Solution().twoCitySchedCost([[259, 770], [448, 54], [926, 667], [184, 139], [840, 118], [577, 469]]) == 1859
    print("OK")
