# 219. Contains Duplicate II (Easy)
# https://leetcode.com/problems/contains-duplicate-ii/
# Run: python 0219_contains_duplicate_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        window = set()
        for i, x in enumerate(nums):
            if x in window:
                return True
            window.add(x)
            if len(window) > k:
                window.remove(nums[i - k])
        return False


if __name__ == "__main__":
    assert Solution().containsNearbyDuplicate([1, 2, 3, 1], 3) is True
    assert Solution().containsNearbyDuplicate([1, 0, 1, 1], 1) is True
    assert Solution().containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2) is False
    assert Solution().containsNearbyDuplicate([1, 1], 0) is False
    print("OK")
