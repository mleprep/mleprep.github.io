# 1197. Minimum Knight Moves (Medium)
# https://leetcode.com/problems/minimum-knight-moves/
# Run: python 1197_minimum_knight_moves.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minKnightMoves(self, x: int, y: int) -> int:
        x, y = abs(x), abs(y)
        seen, q = {(0, 0)}, deque([(0, 0, 0)])
        while q:
            a, b, d = q.popleft()
            if (a, b) == (x, y):
                return d
            for da, db in ((1, 2), (2, 1), (2, -1), (1, -2), (-1, -2), (-2, -1), (-2, 1), (-1, 2)):
                na, nb = a + da, b + db
                if -2 <= na <= x + 2 and -2 <= nb <= y + 2 and (na, nb) not in seen:
                    seen.add((na, nb))
                    q.append((na, nb, d + 1))


if __name__ == "__main__":
    assert Solution().minKnightMoves(2, 1) == 1
    assert Solution().minKnightMoves(5, 5) == 4
    assert Solution().minKnightMoves(0, 0) == 0
    assert Solution().minKnightMoves(1, 1) == 2
    assert Solution().minKnightMoves(1, 0) == 3
    assert Solution().minKnightMoves(-1, -2) == 1
    assert Solution().minKnightMoves(300, 0) == 150
    print("OK")
