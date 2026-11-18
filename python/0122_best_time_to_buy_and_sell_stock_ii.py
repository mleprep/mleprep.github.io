# 122. Best Time to Buy and Sell Stock II (Medium)
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
# Run: python 0122_best_time_to_buy_and_sell_stock_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        return sum(max(0, b - a) for a, b in zip(prices, prices[1:]))


if __name__ == "__main__":
    assert Solution().maxProfit([7, 1, 5, 3, 6, 4]) == 7
    assert Solution().maxProfit([1, 2, 3, 4, 5]) == 4
    print("OK")
