# 338. Counting Bits (Easy)
# https://leetcode.com/problems/counting-bits/
# Run: python 0338_counting_bits.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countBits(self, n: int) -> List[int]:
        bits = [0] * (n + 1)
        for i in range(1, n + 1):
            bits[i] = bits[i >> 1] + (i & 1)
        return bits


if __name__ == "__main__":
    assert Solution().countBits(5) == [0, 1, 1, 2, 1, 2]
    assert Solution().countBits(0) == [0]
    assert Solution().countBits(2) == [0, 1, 1]
    print("OK")
