# 1091. Shortest Path in Binary Matrix (Medium)
# https://leetcode.com/problems/shortest-path-in-binary-matrix/
# Run: python 1091_shortest_path_in_binary_matrix.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if grid[0][0] or grid[n - 1][n - 1]:
            return -1
        q = deque([(0, 0, 1)])
        grid[0][0] = 1
        while q:
            r, c, d = q.popleft()
            if r == c == n - 1:
                return d
            for dr in (-1, 0, 1):
                for dc in (-1, 0, 1):
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < n and 0 <= nc < n and grid[nr][nc] == 0:
                        grid[nr][nc] = 1
                        q.append((nr, nc, d + 1))
        return -1


if __name__ == "__main__":
    assert Solution().shortestPathBinaryMatrix([[0, 1], [1, 0]]) == 2
    assert Solution().shortestPathBinaryMatrix([[0, 0, 0], [1, 1, 0], [1, 1, 0]]) == 4
    assert Solution().shortestPathBinaryMatrix([[1, 0, 0], [1, 1, 0], [1, 1, 0]]) == -1
    print("OK")
