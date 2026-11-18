# 853. Car Fleet (Medium)
# https://leetcode.com/problems/car-fleet/
# Run: python 0853_car_fleet.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        fleets, slowest = 0, 0.0
        for p, s in sorted(zip(position, speed), reverse=True):
            t = (target - p) / s
            if t > slowest:          # can't catch the fleet ahead
                fleets += 1
                slowest = t
        return fleets


if __name__ == "__main__":
    assert Solution().carFleet(12, [10, 8, 0, 5, 3], [2, 4, 1, 1, 3]) == 3
    assert Solution().carFleet(100, [0, 2, 4], [4, 2, 1]) == 1
    print("OK")
