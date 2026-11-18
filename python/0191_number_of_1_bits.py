# 191. Number of 1 Bits (Easy)
# https://leetcode.com/problems/number-of-1-bits/
# Run: python 0191_number_of_1_bits.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        while n:
            n &= n - 1
            count += 1
        return count


if __name__ == "__main__":
    assert Solution().hammingWeight(11) == 3
    assert Solution().hammingWeight(2147483645) == 30
    print("OK")
