# 377. Combination Sum IV (Medium)
# https://leetcode.com/problems/combination-sum-iv/
# Run: python 0377_combination_sum_iv.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        dp = [1] + [0] * target
        for t in range(1, target + 1):
            for x in nums:
                if x <= t:
                    dp[t] += dp[t - x]
        return dp[target]


if __name__ == "__main__":
    assert Solution().combinationSum4([1, 2, 3], 4) == 7
    assert Solution().combinationSum4([9], 3) == 0
    print("OK")
