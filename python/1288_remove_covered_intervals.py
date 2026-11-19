# 1288. Remove Covered Intervals (Medium)
# https://leetcode.com/problems/remove-covered-intervals/
# Run: python 1288_remove_covered_intervals.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        keep, max_end = 0, -1
        for s, e in sorted(intervals, key=lambda x: (x[0], -x[1])):
            if e > max_end:
                keep += 1
                max_end = e
        return keep


if __name__ == "__main__":
    assert Solution().removeCoveredIntervals([[1, 4], [3, 6], [2, 8]]) == 2
    assert Solution().removeCoveredIntervals([[1, 4], [2, 3]]) == 1
    print("OK")
