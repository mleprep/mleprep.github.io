# 209. Minimum Size Subarray Sum (Medium)
# https://leetcode.com/problems/minimum-size-subarray-sum/
# Run: python 0209_minimum_size_subarray_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = s = 0
        best = float("inf")
        for r, x in enumerate(nums):
            s += x
            while s >= target:
                best = min(best, r - l + 1)
                s -= nums[l]
                l += 1
        return 0 if best == float("inf") else best


if __name__ == "__main__":
    assert Solution().minSubArrayLen(7, [2, 3, 1, 2, 4, 3]) == 2
    assert Solution().minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]) == 0
    print("OK")
