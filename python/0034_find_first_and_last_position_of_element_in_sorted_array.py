# 34. Find First and Last Position of Element in Sorted Array (Medium)
# https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
# Run: python 0034_find_first_and_last_position_of_element_in_sorted_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def lower(t):
            lo, hi = 0, len(nums)
            while lo < hi:
                mid = (lo + hi) // 2
                if nums[mid] >= t:
                    hi = mid
                else:
                    lo = mid + 1
            return lo

        a = lower(target)
        if a == len(nums) or nums[a] != target:
            return [-1, -1]
        return [a, lower(target + 1) - 1]


if __name__ == "__main__":
    assert Solution().searchRange([5, 7, 7, 8, 8, 10], 8) == [3, 4]
    assert Solution().searchRange([5, 7, 7, 8, 8, 10], 6) == [-1, -1]
    assert Solution().searchRange([], 0) == [-1, -1]
    print("OK")
