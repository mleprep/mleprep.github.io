# 1493. Longest Subarray of 1's After Deleting One Element (Medium)
# https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
# Run: python 1493_longest_subarray_of_1s_after_deleting_one_element.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        lo = zeros = best = 0
        for hi, x in enumerate(nums):
            zeros += x == 0
            while zeros > 1:
                zeros -= nums[lo] == 0
                lo += 1
            best = max(best, hi - lo)          # (hi - lo + 1) - 1 deleted element
        return best


if __name__ == "__main__":
    assert Solution().longestSubarray([1, 1, 0, 1]) == 3
    assert Solution().longestSubarray([0, 1, 1, 1, 0, 1, 1, 0, 1]) == 5
    assert Solution().longestSubarray([1, 1, 1]) == 2
    assert Solution().longestSubarray([0, 0, 0]) == 0
    print("OK")
