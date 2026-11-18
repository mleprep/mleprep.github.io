# 33. Search in Rotated Sorted Array (Medium)
# https://leetcode.com/problems/search-in-rotated-sorted-array/
# Run: python 0033_search_in_rotated_sorted_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid
            if nums[lo] <= nums[mid]:                    # left half sorted
                if nums[lo] <= target < nums[mid]:
                    hi = mid - 1
                else:
                    lo = mid + 1
            else:                                        # right half sorted
                if nums[mid] < target <= nums[hi]:
                    lo = mid + 1
                else:
                    hi = mid - 1
        return -1


if __name__ == "__main__":
    assert Solution().search([4, 5, 6, 7, 0, 1, 2], 0) == 4
    assert Solution().search([4, 5, 6, 7, 0, 1, 2], 3) == -1
    assert Solution().search([1], 0) == -1
    print("OK")
