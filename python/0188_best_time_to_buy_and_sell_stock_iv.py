# 188. Best Time to Buy and Sell Stock IV (Hard)
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
# Run: python 0188_best_time_to_buy_and_sell_stock_iv.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        if k >= len(prices) // 2:
            return sum(max(0, b - a) for a, b in zip(prices, prices[1:]))
        buy = [float("-inf")] * (k + 1)
        sell = [0] * (k + 1)
        for p in prices:
            for j in range(1, k + 1):
                buy[j] = max(buy[j], sell[j - 1] - p)
                sell[j] = max(sell[j], buy[j] + p)
        return sell[k]


if __name__ == "__main__":
    assert Solution().maxProfit(2, [2, 4, 1]) == 2
    assert Solution().maxProfit(2, [3, 2, 6, 5, 0, 3]) == 7
    print("OK")
