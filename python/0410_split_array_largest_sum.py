# 410. Split Array Largest Sum (Hard)
# https://leetcode.com/problems/split-array-largest-sum/
# Run: python 0410_split_array_largest_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def splitArray(self, nums: List[int], k: int) -> int:
        def parts(cap):
            cnt, s = 1, 0
            for x in nums:
                if s + x > cap:
                    cnt, s = cnt + 1, 0
                s += x
            return cnt

        lo, hi = max(nums), sum(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if parts(mid) <= k:
                hi = mid
            else:
                lo = mid + 1
        return lo


if __name__ == "__main__":
    assert Solution().splitArray([7, 2, 5, 10, 8], 2) == 18
    assert Solution().splitArray([1, 2, 3, 4, 5], 2) == 9
    print("OK")
