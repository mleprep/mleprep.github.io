# 463. Island Perimeter (Easy)
# https://leetcode.com/problems/island-perimeter/
# Run: python 0463_island_perimeter.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        land = shared = 0
        for r, row in enumerate(grid):
            for c, v in enumerate(row):
                if v:
                    land += 1
                    shared += (r + 1 < len(grid) and grid[r + 1][c]) + (c + 1 < len(row) and row[c + 1])
        return 4 * land - 2 * shared


if __name__ == "__main__":
    assert Solution().islandPerimeter([[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]]) == 16
    assert Solution().islandPerimeter([[1]]) == 4
    assert Solution().islandPerimeter([[1, 0]]) == 4
    assert Solution().islandPerimeter([[1, 1], [1, 1]]) == 8
    print("OK")
