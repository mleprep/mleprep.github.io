# 309. Best Time to Buy and Sell Stock with Cooldown (Medium)
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
# Run: python 0309_best_time_to_buy_and_sell_stock_with_cooldown.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        hold, sold, rest = float("-inf"), 0, 0
        for p in prices:
            hold, sold, rest = max(hold, rest - p), hold + p, max(rest, sold)
        return max(sold, rest)


if __name__ == "__main__":
    assert Solution().maxProfit([1, 2, 3, 0, 2]) == 3
    assert Solution().maxProfit([1]) == 0
    print("OK")
