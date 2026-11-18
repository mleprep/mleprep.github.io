# 1011. Capacity To Ship Packages Within D Days (Medium)
# https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
# Run: python 1011_capacity_to_ship_packages_within_d_days.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        def need(cap):
            d, load = 1, 0
            for w in weights:
                if load + w > cap:
                    d, load = d + 1, 0
                load += w
            return d

        lo, hi = max(weights), sum(weights)
        while lo < hi:
            mid = (lo + hi) // 2
            if need(mid) <= days:
                hi = mid
            else:
                lo = mid + 1
        return lo


if __name__ == "__main__":
    assert Solution().shipWithinDays([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5) == 15
    assert Solution().shipWithinDays([3, 2, 2, 4, 1, 4], 3) == 6
    print("OK")
