# 198. House Robber (Medium)
# https://leetcode.com/problems/house-robber/
# Run: python 0198_house_robber.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rob(self, nums: List[int]) -> int:
        prev, cur = 0, 0
        for x in nums:
            prev, cur = cur, max(cur, prev + x)
        return cur


if __name__ == "__main__":
    assert Solution().rob([1, 2, 3, 1]) == 4
    assert Solution().rob([2, 7, 9, 3, 1]) == 12
    print("OK")
