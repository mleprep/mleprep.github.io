# 213. House Robber II (Medium)
# https://leetcode.com/problems/house-robber-ii/
# Run: python 0213_house_robber_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rob(self, nums: List[int]) -> int:
        def line(arr):
            prev = cur = 0
            for x in arr:
                prev, cur = cur, max(cur, prev + x)
            return cur

        if len(nums) == 1:
            return nums[0]
        return max(line(nums[1:]), line(nums[:-1]))


if __name__ == "__main__":
    assert Solution().rob([2, 3, 2]) == 3
    assert Solution().rob([1, 2, 3, 1]) == 4
    assert Solution().rob([1]) == 1
    print("OK")
