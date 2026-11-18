# 746. Min Cost Climbing Stairs (Easy)
# https://leetcode.com/problems/min-cost-climbing-stairs/
# Run: python 0746_min_cost_climbing_stairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        a = b = 0            # dp[i-2], dp[i-1]
        for i in range(2, len(cost) + 1):
            a, b = b, min(b + cost[i - 1], a + cost[i - 2])
        return b


if __name__ == "__main__":
    assert Solution().minCostClimbingStairs([10, 15, 20]) == 15
    assert Solution().minCostClimbingStairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1]) == 6
    print("OK")
