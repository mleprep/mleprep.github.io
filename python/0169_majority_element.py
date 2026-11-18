# 169. Majority Element (Easy)
# https://leetcode.com/problems/majority-element/
# Run: python 0169_majority_element.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cand, cnt = None, 0
        for x in nums:
            if cnt == 0:
                cand = x
            cnt += 1 if x == cand else -1
        return cand


if __name__ == "__main__":
    assert Solution().majorityElement([3, 2, 3]) == 3
    assert Solution().majorityElement([2, 2, 1, 1, 1, 2, 2]) == 2
    print("OK")
