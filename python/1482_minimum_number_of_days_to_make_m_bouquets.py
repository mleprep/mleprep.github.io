# 1482. Minimum Number of Days to Make m Bouquets (Medium)
# https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
# Run: python 1482_minimum_number_of_days_to_make_m_bouquets.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        if m * k > len(bloomDay):
            return -1

        def ok(day):
            made = run = 0
            for b in bloomDay:
                run = run + 1 if b <= day else 0
                if run == k:
                    made, run = made + 1, 0
            return made >= m

        lo, hi = min(bloomDay), max(bloomDay)
        while lo < hi:
            mid = (lo + hi) // 2
            if ok(mid):
                hi = mid
            else:
                lo = mid + 1
        return lo


if __name__ == "__main__":
    assert Solution().minDays([1, 10, 3, 10, 2], 3, 1) == 3
    assert Solution().minDays([1, 10, 3, 10, 2], 3, 2) == -1
    assert Solution().minDays([7, 7, 7, 7, 12, 7, 7], 2, 3) == 12
    print("OK")
