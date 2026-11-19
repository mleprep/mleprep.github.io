# 289. Game of Life (Medium)
# https://leetcode.com/problems/game-of-life/
# Run: python 0289_game_of_life.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        R, C = len(board), len(board[0])
        for r in range(R):
            for c in range(C):
                live = sum(board[i][j] & 1 for i in range(max(0, r - 1), min(R, r + 2))
                           for j in range(max(0, c - 1), min(C, c + 2))) - (board[r][c] & 1)
                if live == 3 or (live == 2 and board[r][c] & 1):
                    board[r][c] |= 2
        for r in range(R):
            for c in range(C):
                board[r][c] >>= 1


if __name__ == "__main__":
    b = [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]
    Solution().gameOfLife(b)
    assert b == [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]]
    b = [[1, 1], [1, 0]]; Solution().gameOfLife(b)
    assert b == [[1, 1], [1, 1]]
    print("OK")
