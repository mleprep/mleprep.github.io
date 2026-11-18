# 53. Maximum Subarray (Medium)
# https://leetcode.com/problems/maximum-subarray/
# Run: python 0053_maximum_subarray.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        best = cur = nums[0]
        for x in nums[1:]:
            cur = max(x, cur + x)
            best = max(best, cur)
        return best


if __name__ == "__main__":
    assert Solution().maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6
    assert Solution().maxSubArray([-1]) == -1
    print("OK")
