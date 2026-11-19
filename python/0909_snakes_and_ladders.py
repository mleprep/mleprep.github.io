# 909. Snakes and Ladders (Medium)
# https://leetcode.com/problems/snakes-and-ladders/
# Run: python 0909_snakes_and_ladders.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def snakesAndLadders(self, board: List[List[int]]) -> int:
        n = len(board)

        def cell(s):
            r, c = divmod(s - 1, n)
            if r % 2:
                c = n - 1 - c
            return board[n - 1 - r][c]

        q, seen, moves = deque([1]), {1}, 0
        while q:
            for _ in range(len(q)):
                s = q.popleft()
                if s == n * n:
                    return moves
                for nxt in range(s + 1, min(s + 6, n * n) + 1):
                    dest = cell(nxt) if cell(nxt) != -1 else nxt
                    if dest not in seen:
                        seen.add(dest)
                        q.append(dest)
            moves += 1
        return -1


if __name__ == "__main__":
    B = [[-1, -1, -1, -1, -1, -1], [-1, -1, -1, -1, -1, -1], [-1, -1, -1, -1, -1, -1],
         [-1, 35, -1, -1, 13, -1], [-1, -1, -1, -1, -1, -1], [-1, 15, -1, -1, -1, -1]]
    assert Solution().snakesAndLadders(B) == 4
    assert Solution().snakesAndLadders([[-1, -1], [-1, 3]]) == 1
    print("OK")
