# 852. Peak Index in a Mountain Array (Medium)
# https://leetcode.com/problems/peak-index-in-a-mountain-array/
# Run: python 0852_peak_index_in_a_mountain_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        lo, hi = 0, len(arr) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if arr[mid] > arr[mid + 1]:
                hi = mid
            else:
                lo = mid + 1
        return lo


if __name__ == "__main__":
    assert Solution().peakIndexInMountainArray([0, 1, 0]) == 1
    assert Solution().peakIndexInMountainArray([0, 10, 5, 2]) == 1
    assert Solution().peakIndexInMountainArray([3, 5, 3, 2, 0]) == 1
    print("OK")
