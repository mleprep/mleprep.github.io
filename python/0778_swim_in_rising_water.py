# 778. Swim in Rising Water (Hard)
# https://leetcode.com/problems/swim-in-rising-water/
# Run: python 0778_swim_in_rising_water.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        pq, seen = [(grid[0][0], 0, 0)], {(0, 0)}
        while pq:
            t, r, c = heappop(pq)
            if r == c == n - 1:
                return t
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in seen:
                    seen.add((nr, nc))
                    heappush(pq, (max(t, grid[nr][nc]), nr, nc))


if __name__ == "__main__":
    assert Solution().swimInWater([[0, 2], [1, 3]]) == 3
    G = [[0, 1, 2, 3, 4], [24, 23, 22, 21, 5], [12, 13, 14, 15, 16], [11, 17, 18, 19, 20], [10, 9, 8, 7, 6]]
    assert Solution().swimInWater(G) == 16
    print("OK")
