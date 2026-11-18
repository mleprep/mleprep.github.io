# 287. Find the Duplicate Number (Medium)
# https://leetcode.com/problems/find-the-duplicate-number/
# Run: python 0287_find_the_duplicate_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow = fast = 0
        while True:
            slow, fast = nums[slow], nums[nums[fast]]
            if slow == fast:
                break
        p = 0
        while p != slow:
            p, slow = nums[p], nums[slow]
        return p


if __name__ == "__main__":
    assert Solution().findDuplicate([1, 3, 4, 2, 2]) == 2
    assert Solution().findDuplicate([3, 1, 3, 4, 2]) == 3
    print("OK")
