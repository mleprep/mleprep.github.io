# 371. Sum of Two Integers (Medium)
# https://leetcode.com/problems/sum-of-two-integers/
# Run: python 0371_sum_of_two_integers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getSum(self, a: int, b: int) -> int:
        mask = 0xFFFFFFFF
        while b & mask:
            a, b = (a ^ b) & mask, ((a & b) << 1) & mask
        a &= mask
        return a if a < 1 << 31 else a - (1 << 32)


if __name__ == "__main__":
    assert Solution().getSum(1, 2) == 3
    assert Solution().getSum(2, 3) == 5
    assert Solution().getSum(-1, 1) == 0
    assert Solution().getSum(-12, -8) == -20
    print("OK")
