# 6. Zigzag Conversion (Medium)
# https://leetcode.com/problems/zigzag-conversion/
# Run: python 0006_zigzag_conversion.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or numRows >= len(s):
            return s
        rows = [[] for _ in range(numRows)]
        r, step = 0, 1
        for ch in s:
            rows[r].append(ch)
            if r == 0:
                step = 1
            elif r == numRows - 1:
                step = -1
            r += step
        return "".join("".join(row) for row in rows)


if __name__ == "__main__":
    assert Solution().convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR"
    assert Solution().convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI"
    assert Solution().convert("A", 1) == "A"
    assert Solution().convert("AB", 1) == "AB"
    assert Solution().convert("ABC", 2) == "ACB"
    print("OK")
