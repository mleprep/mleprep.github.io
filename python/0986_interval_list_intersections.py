# 986. Interval List Intersections (Medium)
# https://leetcode.com/problems/interval-list-intersections/
# Run: python 0986_interval_list_intersections.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def intervalIntersection(self, firstList: List[List[int]], secondList: List[List[int]]) -> List[List[int]]:
        i = j = 0
        res = []
        while i < len(firstList) and j < len(secondList):
            lo = max(firstList[i][0], secondList[j][0])
            hi = min(firstList[i][1], secondList[j][1])
            if lo <= hi:
                res.append([lo, hi])
            if firstList[i][1] < secondList[j][1]:
                i += 1
            else:
                j += 1
        return res


if __name__ == "__main__":
    A = [[0, 2], [5, 10], [13, 23], [24, 25]]
    B = [[1, 5], [8, 12], [15, 24], [25, 26]]
    assert Solution().intervalIntersection(A, B) == [[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]]
    assert Solution().intervalIntersection([[1, 3], [5, 9]], []) == []
    assert Solution().intervalIntersection([[1, 7]], [[3, 10]]) == [[3, 7]]
    print("OK")
