# 759. Employee Free Time (Hard)
# https://leetcode.com/problems/employee-free-time/
# Run: python 0759_employee_free_time.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Interval:  # provided by LeetCode
    def __init__(self, start: int = None, end: int = None):
        self.start = start
        self.end = end


class Solution:
    def employeeFreeTime(self, schedule: '[[Interval]]') -> '[Interval]':
        ivs = sorted((iv for emp in schedule for iv in emp), key=lambda iv: iv.start)
        res, end = [], ivs[0].end
        for iv in ivs[1:]:
            if iv.start > end:
                res.append(Interval(end, iv.start))
            end = max(end, iv.end)
        return res


if __name__ == "__main__":
    def free(s):
        return [[iv.start, iv.end] for iv in Solution().employeeFreeTime([[Interval(a, b) for a, b in emp] for emp in s])]
    assert free([[[1, 2], [5, 6]], [[1, 3]], [[4, 10]]]) == [[3, 4]]
    assert free([[[1, 3], [6, 7]], [[2, 4]], [[2, 5], [9, 12]]]) == [[5, 6], [7, 9]]
    assert free([[[1, 2]], [[2, 3]]]) == []
    assert free([[[1, 2], [8, 9]], [[3, 4]]]) == [[2, 3], [4, 8]]
    print("OK")
