# 1318. Minimum Flips to Make a OR b Equal to c (Medium)
# https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
# Run: python 1318_minimum_flips_to_make_a_or_b_equal_to_c.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minFlips(self, a: int, b: int, c: int) -> int:
        flips = 0
        while a or b or c:
            x, y, z = a & 1, b & 1, c & 1
            flips += (x | y) ^ 1 if z else x + y
            a, b, c = a >> 1, b >> 1, c >> 1
        return flips


if __name__ == "__main__":
    assert Solution().minFlips(2, 6, 5) == 3
    assert Solution().minFlips(4, 2, 7) == 1
    assert Solution().minFlips(1, 2, 3) == 0
    assert Solution().minFlips(8, 3, 5) == 3
    assert all(Solution().minFlips(a, b, c) == bin((a | b) ^ c).count("1") + bin(a & b & ~c).count("1")
               for a in range(16) for b in range(16) for c in range(16))
    print("OK")
