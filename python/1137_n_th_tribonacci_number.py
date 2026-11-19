# 1137. N-th Tribonacci Number (Easy)
# https://leetcode.com/problems/n-th-tribonacci-number/
# Run: python 1137_n_th_tribonacci_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def tribonacci(self, n: int) -> int:
        a, b, c = 0, 1, 1
        for _ in range(n):
            a, b, c = b, c, a + b + c
        return a


if __name__ == "__main__":
    assert Solution().tribonacci(4) == 4
    assert Solution().tribonacci(25) == 1389537
    assert [Solution().tribonacci(i) for i in range(4)] == [0, 1, 1, 2]
    assert Solution().tribonacci(37) == 2082876103
    print("OK")
