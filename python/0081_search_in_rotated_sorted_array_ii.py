# 81. Search in Rotated Sorted Array II (Medium)
# https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
# Run: python 0081_search_in_rotated_sorted_array_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return True
            if nums[lo] == nums[mid] == nums[hi]:
                lo, hi = lo + 1, hi - 1
            elif nums[lo] <= nums[mid]:
                if nums[lo] <= target < nums[mid]:
                    hi = mid - 1
                else:
                    lo = mid + 1
            else:
                if nums[mid] < target <= nums[hi]:
                    lo = mid + 1
                else:
                    hi = mid - 1
        return False


if __name__ == "__main__":
    assert Solution().search([2, 5, 6, 0, 0, 1, 2], 0) is True
    assert Solution().search([2, 5, 6, 0, 0, 1, 2], 3) is False
    assert Solution().search([1, 0, 1, 1, 1], 0) is True
    print("OK")
