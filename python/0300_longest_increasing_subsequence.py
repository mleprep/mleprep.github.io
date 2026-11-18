# 300. Longest Increasing Subsequence (Medium)
# https://leetcode.com/problems/longest-increasing-subsequence/
# Run: python 0300_longest_increasing_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        tails = []
        for x in nums:
            i = bisect.bisect_left(tails, x)
            if i == len(tails):
                tails.append(x)
            else:
                tails[i] = x
        return len(tails)


if __name__ == "__main__":
    assert Solution().lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]) == 4
    assert Solution().lengthOfLIS([0, 1, 0, 3, 2, 3]) == 4
    assert Solution().lengthOfLIS([7, 7, 7, 7]) == 1
    print("OK")
