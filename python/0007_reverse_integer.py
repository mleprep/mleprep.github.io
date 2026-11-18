# 7. Reverse Integer (Medium)
# https://leetcode.com/problems/reverse-integer/
# Run: python 0007_reverse_integer.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverse(self, x: int) -> int:
        sign = -1 if x < 0 else 1
        x, res = abs(x), 0
        while x:
            x, d = divmod(x, 10)
            res = res * 10 + d
        res *= sign
        return res if -2 ** 31 <= res <= 2 ** 31 - 1 else 0


if __name__ == "__main__":
    assert Solution().reverse(123) == 321
    assert Solution().reverse(-123) == -321
    assert Solution().reverse(120) == 21
    assert Solution().reverse(1534236469) == 0
    print("OK")
