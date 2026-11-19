# 260. Single Number III (Medium)
# https://leetcode.com/problems/single-number-iii/
# Run: python 0260_single_number_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        x = reduce(lambda a, b: a ^ b, nums)
        diff = x & -x
        a = 0
        for v in nums:
            if v & diff:
                a ^= v
        return sorted([a, x ^ a])


if __name__ == "__main__":
    assert Solution().singleNumber([1, 2, 1, 3, 2, 5]) == [3, 5]
    assert Solution().singleNumber([-1, 0]) == [-1, 0]
    print("OK")
