# 1109. Corporate Flight Bookings (Medium)
# https://leetcode.com/problems/corporate-flight-bookings/
# Run: python 1109_corporate_flight_bookings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        d = [0] * (n + 1)
        for first, last, seats in bookings:
            d[first - 1] += seats
            d[last] -= seats
        return list(itertools.accumulate(d[:n]))


if __name__ == "__main__":
    assert Solution().corpFlightBookings([[1, 2, 10], [2, 3, 20], [2, 5, 25]], 5) == [10, 55, 45, 25, 25]
    assert Solution().corpFlightBookings([[1, 2, 10], [2, 2, 15]], 2) == [10, 25]
    print("OK")
