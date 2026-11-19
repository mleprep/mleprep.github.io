# 695. Max Area of Island (Medium)
# https://leetcode.com/problems/max-area-of-island/
# Run: python 0695_max_area_of_island.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        R, C = len(grid), len(grid[0])

        def area(r, c):
            if not (0 <= r < R and 0 <= c < C) or grid[r][c] != 1:
                return 0
            grid[r][c] = 0
            return 1 + area(r + 1, c) + area(r - 1, c) + area(r, c + 1) + area(r, c - 1)

        return max(area(r, c) for r in range(R) for c in range(C))


if __name__ == "__main__":
    g = [[0, 0, 1, 0, 0], [1, 1, 1, 0, 0], [0, 0, 0, 1, 1], [0, 0, 0, 1, 0]]
    assert Solution().maxAreaOfIsland(g) == 4
    assert Solution().maxAreaOfIsland([[0, 0, 0, 0, 0]]) == 0
    assert Solution().maxAreaOfIsland([[1]]) == 1
    print("OK")
