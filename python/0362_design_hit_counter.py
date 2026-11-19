# 362. Design Hit Counter (Medium)
# https://leetcode.com/problems/design-hit-counter/
# Run: python 0362_design_hit_counter.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class HitCounter:
    def __init__(self):
        self.q = deque()          # [timestamp, count], oldest first
        self.total = 0

    def hit(self, timestamp: int) -> None:
        if self.q and self.q[-1][0] == timestamp:
            self.q[-1][1] += 1
        else:
            self.q.append([timestamp, 1])
        self.total += 1

    def getHits(self, timestamp: int) -> int:
        while self.q and self.q[0][0] <= timestamp - 300:
            self.total -= self.q.popleft()[1]
        return self.total


if __name__ == "__main__":
    hc = HitCounter()
    hc.hit(1); hc.hit(2); hc.hit(3)
    assert hc.getHits(4) == 3
    hc.hit(300)
    assert hc.getHits(300) == 4 and hc.getHits(301) == 3
    hc2 = HitCounter()
    for _ in range(1000):
        hc2.hit(1)
    assert hc2.getHits(300) == 1000 and hc2.getHits(301) == 0
    print("OK")
