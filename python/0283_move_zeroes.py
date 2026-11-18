# 283. Move Zeroes (Easy)
# https://leetcode.com/problems/move-zeroes/
# Run: python 0283_move_zeroes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        k = 0
        for i in range(len(nums)):
            if nums[i]:
                nums[k], nums[i] = nums[i], nums[k]
                k += 1


if __name__ == "__main__":
    a = [0, 1, 0, 3, 12]; Solution().moveZeroes(a)
    assert a == [1, 3, 12, 0, 0]
    print("OK")
