# 1326. Minimum Number of Taps to Open to Water a Garden (Hard)
# https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
# Run: python 1326_minimum_number_of_taps_to_open_to_water_a_garden.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minTaps(self, n: int, ranges: List[int]) -> int:
        reach = [0] * (n + 1)
        for i, r in enumerate(ranges):
            lo, hi = max(0, i - r), min(n, i + r)
            reach[lo] = max(reach[lo], hi)
        taps = end = far = 0
        for i in range(n):
            far = max(far, reach[i])
            if i == end:
                if far <= i:
                    return -1
                taps += 1
                end = far
        return taps


if __name__ == "__main__":
    assert Solution().minTaps(5, [3, 4, 1, 1, 0, 0]) == 1
    assert Solution().minTaps(3, [0, 0, 0, 0]) == -1
    assert Solution().minTaps(7, [1, 2, 1, 0, 2, 1, 0, 1]) == 3
    print("OK")
