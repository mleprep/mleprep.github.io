# 473. Matchsticks to Square (Medium)
# https://leetcode.com/problems/matchsticks-to-square/
# Run: python 0473_matchsticks_to_square.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def makesquare(self, matchsticks: List[int]) -> bool:
        total = sum(matchsticks)
        if total % 4:
            return False
        side = total // 4
        sticks = sorted(matchsticks, reverse=True)
        if sticks[0] > side:
            return False
        sides = [0] * 4

        def place(i):
            if i == len(sticks):
                return True                    # every side <= side and the total is 4·side -> all equal
            for k in range(4):
                if sides[k] + sticks[i] <= side:
                    sides[k] += sticks[i]
                    if place(i + 1):
                        return True
                    sides[k] -= sticks[i]
                if sides[k] == 0:
                    break
            return False

        return place(0)


if __name__ == "__main__":
    assert Solution().makesquare([1, 1, 2, 2, 2]) is True
    assert Solution().makesquare([3, 3, 3, 3, 4]) is False
    assert Solution().makesquare([5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3]) is True
    assert Solution().makesquare([10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2]) is True
    assert Solution().makesquare([1, 1, 1]) is False
    print("OK")
