# 268. Missing Number (Easy)
# https://leetcode.com/problems/missing-number/
# Run: python 0268_missing_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        x = len(nums)
        for i, v in enumerate(nums):
            x ^= i ^ v
        return x


if __name__ == "__main__":
    assert Solution().missingNumber([3, 0, 1]) == 2
    assert Solution().missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]) == 8
    print("OK")
