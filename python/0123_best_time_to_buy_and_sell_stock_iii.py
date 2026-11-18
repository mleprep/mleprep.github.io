# 123. Best Time to Buy and Sell Stock III (Hard)
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
# Run: python 0123_best_time_to_buy_and_sell_stock_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy1 = buy2 = float("-inf")
        sell1 = sell2 = 0
        for p in prices:
            buy1 = max(buy1, -p)
            sell1 = max(sell1, buy1 + p)
            buy2 = max(buy2, sell1 - p)
            sell2 = max(sell2, buy2 + p)
        return sell2


if __name__ == "__main__":
    assert Solution().maxProfit([3, 3, 5, 0, 0, 3, 1, 4]) == 6
    assert Solution().maxProfit([1, 2, 3, 4, 5]) == 4
    assert Solution().maxProfit([7, 6, 4, 3, 1]) == 0
    print("OK")
