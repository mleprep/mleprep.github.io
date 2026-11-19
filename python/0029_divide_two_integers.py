# 29. Divide Two Integers (Medium)
# https://leetcode.com/problems/divide-two-integers/
# Run: python 0029_divide_two_integers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == -2 ** 31 and divisor == -1:
            return 2 ** 31 - 1
        neg = (dividend < 0) != (divisor < 0)
        a, b, q = abs(dividend), abs(divisor), 0
        while a >= b:
            k = 0
            while a >= b << (k + 1):
                k += 1
            a -= b << k
            q += 1 << k
        return -q if neg else q


if __name__ == "__main__":
    assert Solution().divide(10, 3) == 3
    assert Solution().divide(7, -3) == -2
    assert Solution().divide(-2147483648, -1) == 2147483647
    assert Solution().divide(-2147483648, 2) == -1073741824
    print("OK")
