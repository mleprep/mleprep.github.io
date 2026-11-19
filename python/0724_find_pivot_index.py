# 724. Find Pivot Index (Easy)
# https://leetcode.com/problems/find-pivot-index/
# Run: python 0724_find_pivot_index.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        total, left = sum(nums), 0
        for i, x in enumerate(nums):
            if left == total - left - x:
                return i
            left += x
        return -1


if __name__ == "__main__":
    assert Solution().pivotIndex([1, 7, 3, 6, 5, 6]) == 3
    assert Solution().pivotIndex([1, 2, 3]) == -1
    assert Solution().pivotIndex([2, 1, -1]) == 0
    assert Solution().pivotIndex([0]) == 0
    assert Solution().pivotIndex([-1, -1, 0, 1, 1, 0]) == 5
    print("OK")
