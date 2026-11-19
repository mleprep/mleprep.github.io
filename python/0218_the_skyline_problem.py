# 218. The Skyline Problem (Hard)
# https://leetcode.com/problems/the-skyline-problem/
# Run: python 0218_the_skyline_problem.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        xs = sorted({x for l, r, _ in buildings for x in (l, r)})
        res, h, i = [], [(0, float("inf"))], 0        # heap of (-height, right); ground never ends
        for x in xs:
            while i < len(buildings) and buildings[i][0] <= x:
                l, r, ht = buildings[i]
                heappush(h, (-ht, r))
                i += 1
            while h[0][1] <= x:
                heappop(h)
            top = -h[0][0]
            if not res or res[-1][1] != top:
                res.append([x, top])
        return res


if __name__ == "__main__":
    B = [[2, 9, 10], [3, 7, 15], [5, 12, 12], [15, 20, 10], [19, 24, 8]]
    assert Solution().getSkyline(B) == [[2, 10], [3, 15], [7, 12], [12, 0], [15, 10], [20, 8], [24, 0]]
    assert Solution().getSkyline([[0, 2, 3], [2, 5, 3]]) == [[0, 3], [5, 0]]
    print("OK")
