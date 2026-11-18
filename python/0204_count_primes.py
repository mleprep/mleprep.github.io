# 204. Count Primes (Medium)
# https://leetcode.com/problems/count-primes/
# Run: python 0204_count_primes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0
        is_p = bytearray([1]) * n
        is_p[0] = is_p[1] = 0
        for p in range(2, int(n ** 0.5) + 1):
            if is_p[p]:
                is_p[p * p::p] = bytearray(len(range(p * p, n, p)))
        return sum(is_p)


if __name__ == "__main__":
    assert Solution().countPrimes(10) == 4
    assert Solution().countPrimes(0) == 0
    assert Solution().countPrimes(100) == 25
    print("OK")
