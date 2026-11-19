# 934. Shortest Bridge (Medium)
# https://leetcode.com/problems/shortest-bridge/
# Run: python 0934_shortest_bridge.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def shortestBridge(self, grid: List[List[int]]) -> int:
        n = len(grid)
        start = next((r, c) for r in range(n) for c in range(n) if grid[r][c])
        q, st = deque(), [start]
        grid[start[0]][start[1]] = 2
        while st:                                   # mark island A as 2
            r, c = st.pop()
            q.append((r, c))
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < n and 0 <= nc < n and grid[nr][nc] == 1:
                    grid[nr][nc] = 2
                    st.append((nr, nc))
        steps = 0
        while q:
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < n and 0 <= nc < n:
                        if grid[nr][nc] == 1:
                            return steps
                        if grid[nr][nc] == 0:
                            grid[nr][nc] = 2
                            q.append((nr, nc))
            steps += 1


if __name__ == "__main__":
    assert Solution().shortestBridge([[0, 1], [1, 0]]) == 1
    assert Solution().shortestBridge([[0, 1, 0], [0, 0, 0], [0, 0, 1]]) == 2
    assert Solution().shortestBridge([[1, 1, 1, 1, 1], [1, 0, 0, 0, 1], [1, 0, 1, 0, 1], [1, 0, 0, 0, 1], [1, 1, 1, 1, 1]]) == 1
    print("OK")
