# 137. Single Number II (Medium)
# https://leetcode.com/problems/single-number-ii/
# Run: python 0137_single_number_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for b in range(32):
            if sum((x >> b) & 1 for x in nums) % 3:
                res |= 1 << b
        return res - (1 << 32) if res >= 1 << 31 else res


if __name__ == "__main__":
    assert Solution().singleNumber([2, 2, 3, 2]) == 3
    assert Solution().singleNumber([0, 1, 0, 1, 0, 1, 99]) == 99
    assert Solution().singleNumber([-2, -2, 1, 1, 4, 1, 4, 4, -4, -2]) == -4
    print("OK")
