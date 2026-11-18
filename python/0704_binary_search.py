# 704. Binary Search (Easy)
# https://leetcode.com/problems/binary-search/
# Run: python 0704_binary_search.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                lo = mid + 1
            else:
                hi = mid - 1
        return -1


if __name__ == "__main__":
    assert Solution().search([-1, 0, 3, 5, 9, 12], 9) == 4
    assert Solution().search([-1, 0, 3, 5, 9, 12], 2) == -1
    print("OK")
