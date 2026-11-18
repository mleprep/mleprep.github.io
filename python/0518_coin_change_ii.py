# 518. Coin Change II (Medium)
# https://leetcode.com/problems/coin-change-ii/
# Run: python 0518_coin_change_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        dp = [1] + [0] * amount
        for c in coins:
            for a in range(c, amount + 1):
                dp[a] += dp[a - c]
        return dp[amount]


if __name__ == "__main__":
    assert Solution().change(5, [1, 2, 5]) == 4
    assert Solution().change(3, [2]) == 0
    assert Solution().change(10, [10]) == 1
    print("OK")
