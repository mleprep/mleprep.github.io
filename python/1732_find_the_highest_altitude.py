# 1732. Find the Highest Altitude (Easy)
# https://leetcode.com/problems/find-the-highest-altitude/
# Run: python 1732_find_the_highest_altitude.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        return max(0, max(itertools.accumulate(gain)))


if __name__ == "__main__":
    assert Solution().largestAltitude([-5, 1, 5, 0, -7]) == 1
    assert Solution().largestAltitude([-4, -3, -2, -1, 4, 3, 2]) == 0
    assert Solution().largestAltitude([3]) == 3
    print("OK")
