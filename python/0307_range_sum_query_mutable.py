# 307. Range Sum Query - Mutable (Medium)
# https://leetcode.com/problems/range-sum-query-mutable/
# Run: python 0307_range_sum_query_mutable.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class NumArray:
    def __init__(self, nums: List[int]):
        self.n, self.nums = len(nums), nums[:]
        self.t = [0] * (self.n + 1)
        for i, x in enumerate(nums):
            self._add(i + 1, x)

    def _add(self, i, v):
        while i <= self.n:
            self.t[i] += v
            i += i & -i

    def _pre(self, i):
        s = 0
        while i:
            s += self.t[i]
            i -= i & -i
        return s

    def update(self, index: int, val: int) -> None:
        self._add(index + 1, val - self.nums[index])
        self.nums[index] = val

    def sumRange(self, left: int, right: int) -> int:
        return self._pre(right + 1) - self._pre(left)


if __name__ == "__main__":
    na = NumArray([1, 3, 5])
    assert na.sumRange(0, 2) == 9
    na.update(1, 2)
    assert na.sumRange(0, 2) == 8 and na.sumRange(1, 1) == 2
    print("OK")
