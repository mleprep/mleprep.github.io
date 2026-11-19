# 983. Minimum Cost For Tickets (Medium)
# https://leetcode.com/problems/minimum-cost-for-tickets/
# Run: python 0983_minimum_cost_for_tickets.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mincostTickets(self, days: List[int], costs: List[int]) -> int:
        travel = set(days)
        dp = [0] * (days[-1] + 1)
        for d in range(1, len(dp)):
            if d not in travel:
                dp[d] = dp[d - 1]
            else:
                dp[d] = min(dp[d - 1] + costs[0], dp[max(0, d - 7)] + costs[1], dp[max(0, d - 30)] + costs[2])
        return dp[-1]


if __name__ == "__main__":
    assert Solution().mincostTickets([1, 4, 6, 7, 8, 20], [2, 7, 15]) == 11
    assert Solution().mincostTickets([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31], [2, 7, 15]) == 17
    print("OK")
