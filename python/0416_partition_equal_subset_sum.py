# 416. Partition Equal Subset Sum (Medium)
# https://leetcode.com/problems/partition-equal-subset-sum/
# Run: python 0416_partition_equal_subset_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total % 2:
            return False
        target = total // 2
        dp = [True] + [False] * target
        for x in nums:
            for s in range(target, x - 1, -1):
                dp[s] = dp[s] or dp[s - x]
        return dp[target]


if __name__ == "__main__":
    assert Solution().canPartition([1, 5, 11, 5]) is True
    assert Solution().canPartition([1, 2, 3, 5]) is False
    print("OK")
