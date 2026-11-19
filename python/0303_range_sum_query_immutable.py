# 303. Range Sum Query - Immutable (Easy)
# https://leetcode.com/problems/range-sum-query-immutable/
# Run: python 0303_range_sum_query_immutable.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class NumArray:
    def __init__(self, nums: List[int]):
        self.pre = [0]
        for x in nums:
            self.pre.append(self.pre[-1] + x)

    def sumRange(self, left: int, right: int) -> int:
        return self.pre[right + 1] - self.pre[left]


if __name__ == "__main__":
    na = NumArray([-2, 0, 3, -5, 2, -1])
    assert [na.sumRange(0, 2), na.sumRange(2, 5), na.sumRange(0, 5)] == [1, -1, -3]
    print("OK")
