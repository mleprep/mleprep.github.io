# 334. Increasing Triplet Subsequence (Medium)
# https://leetcode.com/problems/increasing-triplet-subsequence/
# Run: python 0334_increasing_triplet_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        first = second = float("inf")
        for x in nums:
            if x <= first:
                first = x
            elif x <= second:
                second = x
            else:
                return True
        return False


if __name__ == "__main__":
    assert Solution().increasingTriplet([1, 2, 3, 4, 5]) is True
    assert Solution().increasingTriplet([5, 4, 3, 2, 1]) is False
    assert Solution().increasingTriplet([2, 1, 5, 0, 4, 6]) is True
    assert Solution().increasingTriplet([20, 100, 10, 12, 5, 13]) is True
    assert Solution().increasingTriplet([1, 1, 1]) is False
    assert Solution().increasingTriplet([1, 5, 0, 4, 1, 3]) is True
    print("OK")
