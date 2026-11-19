# 315. Count of Smaller Numbers After Self (Hard)
# https://leetcode.com/problems/count-of-smaller-numbers-after-self/
# Run: python 0315_count_of_smaller_numbers_after_self.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        rank = {v: i + 1 for i, v in enumerate(sorted(set(nums)))}
        t = [0] * (len(rank) + 1)
        res = []
        for x in reversed(nums):
            i, s = rank[x] - 1, 0
            while i:
                s += t[i]
                i -= i & -i
            res.append(s)
            i = rank[x]
            while i < len(t):
                t[i] += 1
                i += i & -i
        return res[::-1]


if __name__ == "__main__":
    assert Solution().countSmaller([5, 2, 6, 1]) == [2, 1, 1, 0]
    assert Solution().countSmaller([-1, -1]) == [0, 0]
    print("OK")
