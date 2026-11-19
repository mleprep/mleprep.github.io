# 130. Surrounded Regions (Medium)
# https://leetcode.com/problems/surrounded-regions/
# Run: python 0130_surrounded_regions.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def solve(self, board: List[List[str]]) -> None:
        R, C = len(board), len(board[0])

        def mark(r, c):
            if 0 <= r < R and 0 <= c < C and board[r][c] == "O":
                board[r][c] = "S"
                mark(r + 1, c); mark(r - 1, c); mark(r, c + 1); mark(r, c - 1)

        for r in range(R):
            mark(r, 0); mark(r, C - 1)
        for c in range(C):
            mark(0, c); mark(R - 1, c)
        for r in range(R):
            for c in range(C):
                board[r][c] = "O" if board[r][c] == "S" else "X"


if __name__ == "__main__":
    b = [list(r) for r in ["XXXX", "XOOX", "XXOX", "XOXX"]]
    Solution().solve(b)
    assert ["".join(r) for r in b] == ["XXXX", "XXXX", "XXXX", "XOXX"]
    b = [["X"]]; Solution().solve(b)
    assert b == [["X"]]
    b = [list(r) for r in ["XXX", "XOX", "XXX"]]; Solution().solve(b)
    assert ["".join(r) for r in b] == ["XXX", "XXX", "XXX"]
    b = [list(r) for r in ["OO", "OO"]]; Solution().solve(b)
    assert ["".join(r) for r in b] == ["OO", "OO"]
    print("OK")
