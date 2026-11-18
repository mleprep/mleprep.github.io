# 673. Number of Longest Increasing Subsequence (Medium)
# https://leetcode.com/problems/number-of-longest-increasing-subsequence/
# Run: python 0673_number_of_longest_increasing_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findNumberOfLIS(self, nums: List[int]) -> int:
        n = len(nums)
        length, cnt = [1] * n, [1] * n
        for i in range(n):
            for j in range(i):
                if nums[j] < nums[i]:
                    if length[j] + 1 > length[i]:
                        length[i], cnt[i] = length[j] + 1, cnt[j]
                    elif length[j] + 1 == length[i]:
                        cnt[i] += cnt[j]
        best = max(length)
        return sum(c for l, c in zip(length, cnt) if l == best)


if __name__ == "__main__":
    assert Solution().findNumberOfLIS([1, 3, 5, 4, 7]) == 2
    assert Solution().findNumberOfLIS([2, 2, 2, 2, 2]) == 5
    print("OK")
