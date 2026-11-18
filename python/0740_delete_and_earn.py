# 740. Delete and Earn (Medium)
# https://leetcode.com/problems/delete-and-earn/
# Run: python 0740_delete_and_earn.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def deleteAndEarn(self, nums: List[int]) -> int:
        gain = [0] * (max(nums) + 1)
        for x in nums:
            gain[x] += x
        prev = cur = 0
        for g in gain:
            prev, cur = cur, max(cur, prev + g)
        return cur


if __name__ == "__main__":
    assert Solution().deleteAndEarn([3, 4, 2]) == 6
    assert Solution().deleteAndEarn([2, 2, 3, 3, 3, 4]) == 9
    print("OK")
