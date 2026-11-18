# 875. Koko Eating Bananas (Medium)
# https://leetcode.com/problems/koko-eating-bananas/
# Run: python 0875_koko_eating_bananas.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo, hi = 1, max(piles)
        while lo < hi:
            k = (lo + hi) // 2
            if sum((p + k - 1) // k for p in piles) <= h:
                hi = k
            else:
                lo = k + 1
        return lo


if __name__ == "__main__":
    assert Solution().minEatingSpeed([3, 6, 7, 11], 8) == 4
    assert Solution().minEatingSpeed([30, 11, 23, 4, 20], 5) == 30
    assert Solution().minEatingSpeed([30, 11, 23, 4, 20], 6) == 23
    print("OK")
