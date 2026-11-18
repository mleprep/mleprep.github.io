# 973. K Closest Points to Origin (Medium)
# https://leetcode.com/problems/k-closest-points-to-origin/
# Run: python 0973_k_closest_points_to_origin.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        h = []
        for x, y in points:
            heappush(h, (-(x * x + y * y), x, y))
            if len(h) > k:
                heappop(h)
        return [[x, y] for _, x, y in h]


if __name__ == "__main__":
    assert Solution().kClosest([[1, 3], [-2, 2]], 1) == [[-2, 2]]
    assert sorted(Solution().kClosest([[3, 3], [5, -1], [-2, 4]], 2)) == [[-2, 4], [3, 3]]
    print("OK")
