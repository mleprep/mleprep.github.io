# 217. Contains Duplicate (Easy)
# https://leetcode.com/problems/contains-duplicate/
# Run: python 0217_contains_duplicate.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        seen = set()
        for x in nums:
            if x in seen:
                return True
            seen.add(x)
        return False


if __name__ == "__main__":
    assert Solution().containsDuplicate([1, 2, 3, 1]) is True
    assert Solution().containsDuplicate([1, 2, 3, 4]) is False
    print("OK")
