# 52. N-Queens II (Hard)
# https://leetcode.com/problems/n-queens-ii/
# Run: python 0052_n_queens_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def totalNQueens(self, n: int) -> int:
        full = (1 << n) - 1

        def bt(cols, d1, d2):
            if cols == full:
                return 1
            count, free = 0, ~(cols | d1 | d2) & full
            while free:
                bit = free & -free
                free -= bit
                count += bt(cols | bit, ((d1 | bit) << 1) & full, (d2 | bit) >> 1)
            return count

        return bt(0, 0, 0)


if __name__ == "__main__":
    assert Solution().totalNQueens(4) == 2
    assert Solution().totalNQueens(1) == 1
    assert Solution().totalNQueens(8) == 92
    print("OK")
