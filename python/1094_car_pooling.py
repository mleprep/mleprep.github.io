# 1094. Car Pooling (Medium)
# https://leetcode.com/problems/car-pooling/
# Run: python 1094_car_pooling.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        d = [0] * 1001
        for n, a, b in trips:
            d[a] += n
            d[b] -= n
        load = 0
        for x in d:
            load += x
            if load > capacity:
                return False
        return True


if __name__ == "__main__":
    assert Solution().carPooling([[2, 1, 5], [3, 3, 7]], 4) is False
    assert Solution().carPooling([[2, 1, 5], [3, 3, 7]], 5) is True
    assert Solution().carPooling([[2, 1, 5], [3, 5, 7]], 3) is True
    print("OK")
