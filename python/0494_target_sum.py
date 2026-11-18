# 494. Target Sum (Medium)
# https://leetcode.com/problems/target-sum/
# Run: python 0494_target_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        total = sum(nums)
        if abs(target) > total or (total + target) % 2:
            return 0
        P = (total + target) // 2
        dp = [1] + [0] * P
        for x in nums:
            for s in range(P, x - 1, -1):
                dp[s] += dp[s - x]
        return dp[P]


if __name__ == "__main__":
    assert Solution().findTargetSumWays([1, 1, 1, 1, 1], 3) == 5
    assert Solution().findTargetSumWays([1], 1) == 1
    assert Solution().findTargetSumWays([0, 0, 1], 1) == 4
    print("OK")
