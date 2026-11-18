# 57. Insert Interval (Medium)
# https://leetcode.com/problems/insert-interval/
# Run: python 0057_insert_interval.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        res, i, n = [], 0, len(intervals)
        s, e = newInterval
        while i < n and intervals[i][1] < s:
            res.append(intervals[i]); i += 1
        while i < n and intervals[i][0] <= e:
            s, e = min(s, intervals[i][0]), max(e, intervals[i][1]); i += 1
        res.append([s, e])
        return res + intervals[i:]


if __name__ == "__main__":
    assert Solution().insert([[1, 3], [6, 9]], [2, 5]) == [[1, 5], [6, 9]]
    assert Solution().insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]) == [[1, 2], [3, 10], [12, 16]]
    assert Solution().insert([], [5, 7]) == [[5, 7]]
    print("OK")
