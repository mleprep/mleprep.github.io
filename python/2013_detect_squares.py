# 2013. Detect Squares (Medium)
# https://leetcode.com/problems/detect-squares/
# Run: python 2013_detect_squares.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class DetectSquares:
    def __init__(self):
        self.cnt = Counter()

    def add(self, point: List[int]) -> None:
        self.cnt[tuple(point)] += 1

    def count(self, point: List[int]) -> int:
        x, y = point
        total = 0
        for (px, py), c in list(self.cnt.items()):
            if abs(px - x) != abs(py - y) or px == x:
                continue
            total += c * self.cnt[(x, py)] * self.cnt[(px, y)]
        return total


if __name__ == "__main__":
    ds = DetectSquares()
    for p in ([3, 10], [11, 2], [3, 2]):
        ds.add(p)
    assert ds.count([11, 10]) == 1
    assert ds.count([14, 8]) == 0
    ds.add([11, 2])
    assert ds.count([11, 10]) == 2
    print("OK")
