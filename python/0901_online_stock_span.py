# 901. Online Stock Span (Medium)
# https://leetcode.com/problems/online-stock-span/
# Run: python 0901_online_stock_span.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class StockSpanner:
    def __init__(self):
        self.st = []  # (price, span)

    def next(self, price: int) -> int:
        span = 1
        while self.st and self.st[-1][0] <= price:
            span += self.st.pop()[1]
        self.st.append((price, span))
        return span


if __name__ == "__main__":
    s = StockSpanner()
    assert [s.next(p) for p in [100, 80, 60, 70, 60, 75, 85]] == [1, 1, 1, 2, 1, 4, 6]
    s2 = StockSpanner()
    assert [s2.next(p) for p in [31, 41, 48, 59, 79]] == [1, 2, 3, 4, 5]
    print("OK")
