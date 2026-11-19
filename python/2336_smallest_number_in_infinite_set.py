# 2336. Smallest Number in Infinite Set (Medium)
# https://leetcode.com/problems/smallest-number-in-infinite-set/
# Run: python 2336_smallest_number_in_infinite_set.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class SmallestInfiniteSet:
    def __init__(self):
        self.nxt = 1
        self.back, self.in_back = [], set()

    def popSmallest(self) -> int:
        if self.back:
            x = heappop(self.back)
            self.in_back.remove(x)
            return x
        self.nxt += 1
        return self.nxt - 1

    def addBack(self, num: int) -> None:
        if num < self.nxt and num not in self.in_back:
            self.in_back.add(num)
            heappush(self.back, num)


if __name__ == "__main__":
    s = SmallestInfiniteSet()
    s.addBack(2)
    assert [s.popSmallest(), s.popSmallest(), s.popSmallest()] == [1, 2, 3]
    s.addBack(1)
    assert [s.popSmallest(), s.popSmallest(), s.popSmallest()] == [1, 4, 5]
    s2 = SmallestInfiniteSet()
    s2.popSmallest(); s2.popSmallest()
    s2.addBack(1); s2.addBack(1)
    assert [s2.popSmallest(), s2.popSmallest()] == [1, 3]
    print("OK")
