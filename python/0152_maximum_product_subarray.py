# 152. Maximum Product Subarray (Medium)
# https://leetcode.com/problems/maximum-product-subarray/
# Run: python 0152_maximum_product_subarray.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        best = hi = lo = nums[0]
        for x in nums[1:]:
            cands = (x, hi * x, lo * x)
            hi, lo = max(cands), min(cands)
            best = max(best, hi)
        return best


if __name__ == "__main__":
    assert Solution().maxProduct([2, 3, -2, 4]) == 6
    assert Solution().maxProduct([-2, 0, -1]) == 0
    assert Solution().maxProduct([-2, 3, -4]) == 24
    print("OK")
