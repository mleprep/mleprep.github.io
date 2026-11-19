# 1730. Shortest Path to Get Food (Medium)
# https://leetcode.com/problems/shortest-path-to-get-food/
# Run: python 1730_shortest_path_to_get_food.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getFood(self, grid: List[List[str]]) -> int:
        R, C = len(grid), len(grid[0])
        start = next((r, c) for r in range(R) for c in range(C) if grid[r][c] == "*")
        seen, q, steps = {start}, deque([start]), 0
        while q:
            steps += 1
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and grid[nr][nc] != "X" and (nr, nc) not in seen:
                        if grid[nr][nc] == "#":
                            return steps
                        seen.add((nr, nc))
                        q.append((nr, nc))
        return -1


if __name__ == "__main__":
    G = lambda rows: [list(r) for r in rows]
    assert Solution().getFood(G(["XXXXXX", "X*OOOX", "XOO#OX", "XXXXXX"])) == 3
    assert Solution().getFood(G(["XXXXX", "X*XOX", "XOX#X", "XXXXX"])) == -1
    assert Solution().getFood(G(["XXXXXXXX", "X*OXO#OX", "XOOXOOXX", "XOOOO#OX", "XXXXXXXX"])) == 6
    assert Solution().getFood(G(["*#"])) == 1
    print("OK")
