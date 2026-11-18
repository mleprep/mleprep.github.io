# 45. Jump Game II (Medium)
# https://leetcode.com/problems/jump-game-ii/
# Run: python 0045_jump_game_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def jump(self, nums: List[int]) -> int:
        jumps = end = far = 0
        for i in range(len(nums) - 1):
            far = max(far, i + nums[i])
            if i == end:
                jumps += 1
                end = far
        return jumps


if __name__ == "__main__":
    assert Solution().jump([2, 3, 1, 1, 4]) == 2
    assert Solution().jump([2, 3, 0, 1, 4]) == 2
    assert Solution().jump([0]) == 0
    print("OK")
