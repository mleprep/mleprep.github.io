# 525. Contiguous Array (Medium)
# https://leetcode.com/problems/contiguous-array/
# Run: python 0525_contiguous_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        first = {0: -1}
        pre = best = 0
        for i, x in enumerate(nums):
            pre += 1 if x else -1
            if pre in first:
                best = max(best, i - first[pre])
            else:
                first[pre] = i
        return best


if __name__ == "__main__":
    assert Solution().findMaxLength([0, 1]) == 2
    assert Solution().findMaxLength([0, 1, 1, 0, 1, 1]) == 4
    print("OK")
