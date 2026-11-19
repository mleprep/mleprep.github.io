# 630. Course Schedule III (Hard)
# https://leetcode.com/problems/course-schedule-iii/
# Run: python 0630_course_schedule_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def scheduleCourse(self, courses: List[List[int]]) -> int:
        h, t = [], 0
        for dur, last in sorted(courses, key=lambda c: c[1]):
            heappush(h, -dur)
            t += dur
            if t > last:
                t += heappop(h)          # remove the longest (stored negative)
        return len(h)


if __name__ == "__main__":
    assert Solution().scheduleCourse([[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]) == 3
    assert Solution().scheduleCourse([[1, 2]]) == 1
    assert Solution().scheduleCourse([[3, 2], [4, 3]]) == 0
    print("OK")
