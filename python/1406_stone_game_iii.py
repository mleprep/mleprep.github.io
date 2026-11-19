# 1406. Stone Game III (Hard)
# https://leetcode.com/problems/stone-game-iii/
# Run: python 1406_stone_game_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def stoneGameIII(self, stoneValue: List[int]) -> str:
        n = len(stoneValue)
        dp = [0] * (n + 3)
        for i in range(n - 1, -1, -1):
            take, best = 0, float("-inf")
            for k in range(3):
                if i + k < n:
                    take += stoneValue[i + k]
                    best = max(best, take - dp[i + k + 1])
            dp[i] = best
        return "Alice" if dp[0] > 0 else "Bob" if dp[0] < 0 else "Tie"


if __name__ == "__main__":
    assert Solution().stoneGameIII([1, 2, 3, 7]) == "Bob"
    assert Solution().stoneGameIII([1, 2, 3, -9]) == "Alice"
    assert Solution().stoneGameIII([1, 2, 3, 6]) == "Tie"
    assert Solution().stoneGameIII([-1, -2, -3]) == "Tie"
    print("OK")
