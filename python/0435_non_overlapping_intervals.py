# 435. Non-overlapping Intervals (Medium)
# https://leetcode.com/problems/non-overlapping-intervals/
# Run: python 0435_non_overlapping_intervals.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        keep, end = 0, float("-inf")
        for s, e in sorted(intervals, key=lambda x: x[1]):
            if s >= end:
                keep += 1
                end = e
        return len(intervals) - keep


if __name__ == "__main__":
    assert Solution().eraseOverlapIntervals([[1, 2], [2, 3], [3, 4], [1, 3]]) == 1
    assert Solution().eraseOverlapIntervals([[1, 2], [1, 2], [1, 2]]) == 2
    assert Solution().eraseOverlapIntervals([[1, 2], [2, 3]]) == 0
    print("OK")
