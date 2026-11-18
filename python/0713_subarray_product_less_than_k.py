# 713. Subarray Product Less Than K (Medium)
# https://leetcode.com/problems/subarray-product-less-than-k/
# Run: python 0713_subarray_product_less_than_k.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        if k <= 1:
            return 0
        prod, l, count = 1, 0, 0
        for r, x in enumerate(nums):
            prod *= x
            while prod >= k:
                prod //= nums[l]
                l += 1
            count += r - l + 1
        return count


if __name__ == "__main__":
    assert Solution().numSubarrayProductLessThanK([10, 5, 2, 6], 100) == 8
    assert Solution().numSubarrayProductLessThanK([1, 2, 3], 0) == 0
    print("OK")
