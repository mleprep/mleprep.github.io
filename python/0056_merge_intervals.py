# 56. Merge Intervals (Medium)
# https://leetcode.com/problems/merge-intervals/
# Run: python 0056_merge_intervals.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        res = []
        for s, e in sorted(intervals):
            if res and s <= res[-1][1]:
                res[-1][1] = max(res[-1][1], e)
            else:
                res.append([s, e])
        return res


if __name__ == "__main__":
    assert Solution().merge([[1, 3], [2, 6], [8, 10], [15, 18]]) == [[1, 6], [8, 10], [15, 18]]
    assert Solution().merge([[1, 4], [4, 5]]) == [[1, 5]]
    print("OK")
