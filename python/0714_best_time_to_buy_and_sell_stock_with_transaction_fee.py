# 714. Best Time to Buy and Sell Stock with Transaction Fee (Medium)
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
# Run: python 0714_best_time_to_buy_and_sell_stock_with_transaction_fee.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProfit(self, prices: List[int], fee: int) -> int:
        cash, hold = 0, -prices[0]
        for p in prices[1:]:
            cash = max(cash, hold + p - fee)
            hold = max(hold, cash - p)
        return cash


if __name__ == "__main__":
    assert Solution().maxProfit([1, 3, 2, 8, 4, 9], 2) == 8
    assert Solution().maxProfit([1, 3, 7, 5, 10, 3], 3) == 6
    print("OK")
