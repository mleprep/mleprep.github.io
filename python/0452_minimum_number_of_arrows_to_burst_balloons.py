# 452. Minimum Number of Arrows to Burst Balloons (Medium)
# https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
# Run: python 0452_minimum_number_of_arrows_to_burst_balloons.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        arrows, pos = 0, float("-inf")
        for s, e in sorted(points, key=lambda p: p[1]):
            if s > pos:
                arrows += 1
                pos = e
        return arrows


if __name__ == "__main__":
    assert Solution().findMinArrowShots([[10, 16], [2, 8], [1, 6], [7, 12]]) == 2
    assert Solution().findMinArrowShots([[1, 2], [3, 4], [5, 6], [7, 8]]) == 4
    print("OK")
