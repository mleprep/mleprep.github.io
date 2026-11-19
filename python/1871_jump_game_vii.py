# 1871. Jump Game VII (Medium)
# https://leetcode.com/problems/jump-game-vii/
# Run: python 1871_jump_game_vii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:
        n = len(s)
        ok = [False] * n
        ok[0] = True
        window = 0
        for i in range(1, n):
            if i >= minJump and ok[i - minJump]:
                window += 1
            if i > maxJump and ok[i - maxJump - 1]:
                window -= 1
            ok[i] = window > 0 and s[i] == "0"
        return ok[-1]


if __name__ == "__main__":
    assert Solution().canReach("011010", 2, 3) is True
    assert Solution().canReach("01101110", 2, 3) is False
    assert Solution().canReach("00", 1, 1) is True
    assert Solution().canReach("01", 1, 1) is False
    assert Solution().canReach("0000000000", 8, 8) is False
    print("OK")
