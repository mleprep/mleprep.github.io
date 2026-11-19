# 51. N-Queens (Hard)
# https://leetcode.com/problems/n-queens/
# Run: python 0051_n_queens.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        res, cols, d1, d2 = [], set(), set(), set()
        board = [["."] * n for _ in range(n)]

        def bt(r):
            if r == n:
                res.append(["".join(row) for row in board])
                return
            for c in range(n):
                if c in cols or r - c in d1 or r + c in d2:
                    continue
                cols.add(c); d1.add(r - c); d2.add(r + c)
                board[r][c] = "Q"
                bt(r + 1)
                board[r][c] = "."
                cols.remove(c); d1.remove(r - c); d2.remove(r + c)

        bt(0)
        return res


if __name__ == "__main__":
    assert sorted(Solution().solveNQueens(4)) == sorted([[".Q..", "...Q", "Q...", "..Q."], ["..Q.", "Q...", "...Q", ".Q.."]])
    assert len(Solution().solveNQueens(8)) == 92
    print("OK")
