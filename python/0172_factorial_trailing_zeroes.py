# 172. Factorial Trailing Zeroes (Medium)
# https://leetcode.com/problems/factorial-trailing-zeroes/
# Run: python 0172_factorial_trailing_zeroes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def trailingZeroes(self, n: int) -> int:
        count = 0
        while n:
            n //= 5
            count += n
        return count


if __name__ == "__main__":
    assert Solution().trailingZeroes(3) == 0
    assert Solution().trailingZeroes(5) == 1
    assert Solution().trailingZeroes(100) == 24
    print("OK")
