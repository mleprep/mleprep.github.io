# 36. Valid Sudoku (Medium)
# https://leetcode.com/problems/valid-sudoku/
# Run: python 0036_valid_sudoku.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]
        for r in range(9):
            for c in range(9):
                v = board[r][c]
                if v == ".":
                    continue
                b = (r // 3) * 3 + c // 3
                if v in rows[r] or v in cols[c] or v in boxes[b]:
                    return False
                rows[r].add(v); cols[c].add(v); boxes[b].add(v)
        return True


if __name__ == "__main__":
    B = ["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
         "7...2...6", ".6....28.", "...419..5", "....8..79"]
    board = [list(r) for r in B]
    assert Solution().isValidSudoku(board) is True
    board[0][0] = "8"  # column 0 already has an 8
    assert Solution().isValidSudoku(board) is False
    print("OK")
