# 448. Find All Numbers Disappeared in an Array (Easy)
# https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
# Run: python 0448_find_all_numbers_disappeared_in_an_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        for v in nums:
            i = abs(v) - 1
            nums[i] = -abs(nums[i])
        return [i + 1 for i, v in enumerate(nums) if v > 0]


if __name__ == "__main__":
    assert Solution().findDisappearedNumbers([4, 3, 2, 7, 8, 2, 3, 1]) == [5, 6]
    assert Solution().findDisappearedNumbers([1, 1]) == [2]
    print("OK")
