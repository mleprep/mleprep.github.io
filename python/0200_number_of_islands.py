# 200. Number of Islands (Medium)
# https://leetcode.com/problems/number-of-islands/
# Run: python 0200_number_of_islands.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        R, C = len(grid), len(grid[0])

        def sink(r, c):
            st = [(r, c)]
            grid[r][c] = "0"
            while st:
                r, c = st.pop()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and grid[nr][nc] == "1":
                        grid[nr][nc] = "0"
                        st.append((nr, nc))

        count = 0
        for r in range(R):
            for c in range(C):
                if grid[r][c] == "1":
                    count += 1
                    sink(r, c)
        return count


if __name__ == "__main__":
    g = [list(r) for r in ["11000", "11000", "00100", "00011"]]
    assert Solution().numIslands(g) == 3
    print("OK")
