# 284. Peeking Iterator (Medium)
# https://leetcode.com/problems/peeking-iterator/
# Run: python 0284_peeking_iterator.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class PeekingIterator:
    def __init__(self, iterator):
        self.it = iterator
        self.nxt = iterator.next() if iterator.hasNext() else None
        self.has = self.nxt is not None or iterator.hasNext()

    def peek(self):
        return self.nxt

    def next(self):
        val = self.nxt
        self.has = self.it.hasNext()
        self.nxt = self.it.next() if self.has else None
        return val

    def hasNext(self):
        return self.has


if __name__ == "__main__":
    class Iterator:   # LeetCode's interface, minimal stand-in
        def __init__(self, nums): self.nums, self.i = nums, 0
        def hasNext(self): return self.i < len(self.nums)
        def next(self): self.i += 1; return self.nums[self.i - 1]
    pi = PeekingIterator(Iterator([1, 2, 3]))
    assert [pi.next(), pi.peek(), pi.next(), pi.next(), pi.hasNext()] == [1, 2, 2, 3, False]
    print("OK")
