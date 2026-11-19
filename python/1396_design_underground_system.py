# 1396. Design Underground System (Medium)
# https://leetcode.com/problems/design-underground-system/
# Run: python 1396_design_underground_system.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class UndergroundSystem:
    def __init__(self):
        self.active, self.stats = {}, defaultdict(lambda: [0, 0])

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.active[id] = (stationName, t)

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        start, t0 = self.active.pop(id)
        s = self.stats[(start, stationName)]
        s[0] += t - t0
        s[1] += 1

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        total, n = self.stats[(startStation, endStation)]
        return total / n


if __name__ == "__main__":
    u = UndergroundSystem()
    u.checkIn(45, "Leyton", 3); u.checkIn(32, "Paradise", 8); u.checkIn(27, "Leyton", 10)
    u.checkOut(45, "Waterloo", 15); u.checkOut(27, "Waterloo", 20); u.checkOut(32, "Cambridge", 22)
    assert u.getAverageTime("Paradise", "Cambridge") == 14.0 and u.getAverageTime("Leyton", "Waterloo") == 11.0
    u2 = UndergroundSystem()
    u2.checkIn(10, "A", 1); u2.checkOut(10, "B", 5); u2.checkIn(11, "A", 2); u2.checkOut(11, "B", 4)
    assert u2.getAverageTime("A", "B") == 3.0
    print("OK")
