# 75. Sort Colors (Medium)
# https://leetcode.com/problems/sort-colors/
# Run: python 0075_sort_colors.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortColors(self, nums: List[int]) -> None:
        lo, i, hi = 0, 0, len(nums) - 1
        while i <= hi:
            if nums[i] == 0:
                nums[lo], nums[i] = nums[i], nums[lo]
                lo += 1; i += 1
            elif nums[i] == 2:
                nums[hi], nums[i] = nums[i], nums[hi]
                hi -= 1
            else:
                i += 1


if __name__ == "__main__":
    a = [2, 0, 2, 1, 1, 0]; Solution().sortColors(a)
    assert a == [0, 0, 1, 1, 2, 2]
    a = [2, 0, 1]; Solution().sortColors(a)
    assert a == [0, 1, 2]
    a = [1, 1, 0, 0, 2, 2, 1]; Solution().sortColors(a)
    assert a == [0, 0, 1, 1, 1, 2, 2]
    print("OK")
