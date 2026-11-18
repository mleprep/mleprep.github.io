# 162. Find Peak Element (Medium)
# https://leetcode.com/problems/find-peak-element/
# Run: python 0162_find_peak_element.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] < nums[mid + 1]:
                lo = mid + 1
            else:
                hi = mid
        return lo


if __name__ == "__main__":
    assert Solution().findPeakElement([1, 2, 3, 1]) == 2
    assert Solution().findPeakElement([1, 2, 1, 3, 5, 6, 4]) in (1, 5)
    print("OK")
