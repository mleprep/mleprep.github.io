# 190. Reverse Bits (Easy)
# https://leetcode.com/problems/reverse-bits/
# Run: python 0190_reverse_bits.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        for _ in range(32):
            res = (res << 1) | (n & 1)
            n >>= 1
        return res


if __name__ == "__main__":
    assert Solution().reverseBits(43261596) == 964176192
    assert Solution().reverseBits(4294967293) == 3221225471
    print("OK")
