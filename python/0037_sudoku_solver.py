# 37. Sudoku Solver (Hard)
# https://leetcode.com/problems/sudoku-solver/
# Run: python 0037_sudoku_solver.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]
        empty = []
        for r in range(9):
            for c in range(9):
                v = board[r][c]
                if v == ".":
                    empty.append((r, c))
                else:
                    rows[r].add(v); cols[c].add(v); boxes[r // 3 * 3 + c // 3].add(v)

        def bt(k):
            if k == len(empty):
                return True
            r, c = empty[k]
            b = r // 3 * 3 + c // 3
            for v in "123456789":
                if v in rows[r] or v in cols[c] or v in boxes[b]:
                    continue
                board[r][c] = v
                rows[r].add(v); cols[c].add(v); boxes[b].add(v)
                if bt(k + 1):
                    return True
                rows[r].remove(v); cols[c].remove(v); boxes[b].remove(v)
            board[r][c] = "."
            return False

        bt(0)


if __name__ == "__main__":
    B = ["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"]
    board = [list(r) for r in B]
    Solution().solveSudoku(board)
    assert ["".join(r) for r in board][0] == "534678912" and ["".join(r) for r in board][8] == "345286179"
    assert ["".join(r) for r in board] == ['534678912', '672195348', '198342567', '859761423', '426853791', '713924856', '961537284', '287419635', '345286179']
    print("OK")
