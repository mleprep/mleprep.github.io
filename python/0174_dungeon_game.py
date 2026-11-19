# 174. Dungeon Game (Hard)
# https://leetcode.com/problems/dungeon-game/
# Run: python 0174_dungeon_game.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
        R, C = len(dungeon), len(dungeon[0])
        INF = float("inf")
        need = [INF] * (C + 1)
        need[C - 1] = 1                                   # virtual cell after the princess
        for r in range(R - 1, -1, -1):
            for c in range(C - 1, -1, -1):
                need[c] = max(1, min(need[c], need[c + 1]) - dungeon[r][c])
        return need[0]


if __name__ == "__main__":
    assert Solution().calculateMinimumHP([[-2, -3, 3], [-5, -10, 1], [10, 30, -5]]) == 7
    assert Solution().calculateMinimumHP([[0]]) == 1
    print("OK")
