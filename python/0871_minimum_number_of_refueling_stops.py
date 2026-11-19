# 871. Minimum Number of Refueling Stops (Hard)
# https://leetcode.com/problems/minimum-number-of-refueling-stops/
# Run: python 0871_minimum_number_of_refueling_stops.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minRefuelStops(self, target: int, startFuel: int, stations: List[List[int]]) -> int:
        h, reach, i, stops = [], startFuel, 0, 0
        while reach < target:
            while i < len(stations) and stations[i][0] <= reach:
                heappush(h, -stations[i][1])
                i += 1
            if not h:
                return -1
            reach -= heappop(h)
            stops += 1
        return stops


if __name__ == "__main__":
    assert Solution().minRefuelStops(1, 1, []) == 0
    assert Solution().minRefuelStops(100, 1, [[10, 100]]) == -1
    assert Solution().minRefuelStops(100, 10, [[10, 60], [20, 30], [30, 30], [60, 40]]) == 2
    print("OK")
