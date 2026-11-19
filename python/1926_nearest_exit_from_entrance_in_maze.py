# 1926. Nearest Exit from Entrance in Maze (Medium)
# https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
# Run: python 1926_nearest_exit_from_entrance_in_maze.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        R, C = len(maze), len(maze[0])
        start = tuple(entrance)
        seen, q, steps = {start}, deque([start]), 0
        while q:
            steps += 1
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and maze[nr][nc] == "." and (nr, nc) not in seen:
                        if nr in (0, R - 1) or nc in (0, C - 1):
                            return steps
                        seen.add((nr, nc))
                        q.append((nr, nc))
        return -1


if __name__ == "__main__":
    G = lambda rows: [list(r) for r in rows]
    assert Solution().nearestExit(G(["++.+", "...+", "+++."]), [1, 2]) == 1
    assert Solution().nearestExit(G(["+++", "...", "+++"]), [1, 0]) == 2
    assert Solution().nearestExit(G([".+"]), [0, 0]) == -1
    print("OK")
