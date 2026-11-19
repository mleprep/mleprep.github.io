# 2352. Equal Row and Column Pairs (Medium)
# https://leetcode.com/problems/equal-row-and-column-pairs/
# Run: python 2352_equal_row_and_column_pairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def equalPairs(self, grid: List[List[int]]) -> int:
        rows = Counter(tuple(r) for r in grid)
        return sum(rows[col] for col in zip(*grid))


if __name__ == "__main__":
    assert Solution().equalPairs([[3, 2, 1], [1, 7, 6], [2, 7, 7]]) == 1
    assert Solution().equalPairs([[3, 1, 2, 2], [1, 4, 4, 5], [2, 4, 2, 2], [2, 4, 2, 2]]) == 3
    assert Solution().equalPairs([[1]]) == 1
    print("OK")
