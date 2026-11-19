# 729. My Calendar I (Medium)
# https://leetcode.com/problems/my-calendar-i/
# Run: python 0729_my_calendar_i.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyCalendar:
    def __init__(self):
        self.starts, self.ends = [], []

    def book(self, start: int, end: int) -> bool:
        i = bisect.bisect_right(self.starts, start)
        if i > 0 and self.ends[i - 1] > start:          # previous overlaps
            return False
        if i < len(self.starts) and self.starts[i] < end:   # next overlaps
            return False
        self.starts.insert(i, start)
        self.ends.insert(i, end)
        return True


if __name__ == "__main__":
    c = MyCalendar()
    assert [c.book(10, 20), c.book(15, 25), c.book(20, 30)] == [True, False, True]
    print("OK")
