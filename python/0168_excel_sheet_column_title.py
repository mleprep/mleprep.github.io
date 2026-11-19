# 168. Excel Sheet Column Title (Easy)
# https://leetcode.com/problems/excel-sheet-column-title/
# Run: python 0168_excel_sheet_column_title.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        out = []
        while columnNumber:
            columnNumber -= 1
            out.append(chr(ord("A") + columnNumber % 26))
            columnNumber //= 26
        return "".join(reversed(out))


if __name__ == "__main__":
    assert Solution().convertToTitle(1) == "A"
    assert Solution().convertToTitle(28) == "AB"
    assert Solution().convertToTitle(701) == "ZY"
    assert Solution().convertToTitle(26) == "Z"
    assert Solution().convertToTitle(52) == "AZ"
    assert Solution().convertToTitle(2147483647) == "FXSHRXW"
    print("OK")
