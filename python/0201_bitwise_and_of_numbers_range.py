# 201. Bitwise AND of Numbers Range (Medium)
# https://leetcode.com/problems/bitwise-and-of-numbers-range/
# Run: python 0201_bitwise_and_of_numbers_range.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        shift = 0
        while left != right:
            left, right, shift = left >> 1, right >> 1, shift + 1
        return left << shift


if __name__ == "__main__":
    assert Solution().rangeBitwiseAnd(5, 7) == 4
    assert Solution().rangeBitwiseAnd(0, 0) == 0
    assert Solution().rangeBitwiseAnd(1, 2147483647) == 0
    print("OK")
