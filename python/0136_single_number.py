# 136. Single Number (Easy)
# https://leetcode.com/problems/single-number/
# Run: python 0136_single_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        return reduce(lambda a, b: a ^ b, nums)


if __name__ == "__main__":
    assert Solution().singleNumber([4, 1, 2, 1, 2]) == 4
    print("OK")
