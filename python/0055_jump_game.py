# 55. Jump Game (Medium)
# https://leetcode.com/problems/jump-game/
# Run: python 0055_jump_game.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        reach = 0
        for i, x in enumerate(nums):
            if i > reach:
                return False
            reach = max(reach, i + x)
        return True


if __name__ == "__main__":
    assert Solution().canJump([2, 3, 1, 1, 4]) is True
    assert Solution().canJump([3, 2, 1, 0, 4]) is False
    print("OK")
