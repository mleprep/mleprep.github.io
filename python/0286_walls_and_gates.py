# 286. Walls and Gates (Medium)
# https://leetcode.com/problems/walls-and-gates/
# Run: python 0286_walls_and_gates.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def wallsAndGates(self, rooms: List[List[int]]) -> None:
        INF = 2147483647
        R, C = len(rooms), len(rooms[0])
        q = deque((r, c) for r in range(R) for c in range(C) if rooms[r][c] == 0)
        while q:
            r, c = q.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and rooms[nr][nc] == INF:
                    rooms[nr][nc] = rooms[r][c] + 1
                    q.append((nr, nc))


if __name__ == "__main__":
    I = 2147483647
    g = [[I, -1, 0, I], [I, I, I, -1], [I, -1, I, -1], [0, -1, I, I]]
    Solution().wallsAndGates(g)
    assert g == [[3, -1, 0, 1], [2, 2, 1, -1], [1, -1, 2, -1], [0, -1, 3, 4]]
    g = [[-1]]; Solution().wallsAndGates(g)
    assert g == [[-1]]
    g = [[I, 0]]; Solution().wallsAndGates(g)
    assert g == [[1, 0]]
    print("OK")
