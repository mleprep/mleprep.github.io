# 877. Stone Game (Medium)
# https://leetcode.com/problems/stone-game/
# Run: python 0877_stone_game.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def stoneGame(self, piles: List[int]) -> bool:
        n = len(piles)
        dp = piles[:]                               # dp[i] for intervals of length 1
        for length in range(2, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1
                dp[i] = max(piles[i] - dp[i + 1], piles[j] - dp[i])
        return dp[0] > 0


if __name__ == "__main__":
    assert Solution().stoneGame([5, 3, 4, 5]) is True
    assert Solution().stoneGame([3, 7, 2, 3]) is True
    assert Solution().stoneGame([1, 100, 2, 4]) is True
    print("OK")
