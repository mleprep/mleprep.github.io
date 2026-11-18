# 12. Integer to Roman (Medium)
# https://leetcode.com/problems/integer-to-roman/
# Run: python 0012_integer_to_roman.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def intToRoman(self, num: int) -> str:
        table = [(1000, "M"), (900, "CM"), (500, "D"), (400, "CD"), (100, "C"), (90, "XC"),
                 (50, "L"), (40, "XL"), (10, "X"), (9, "IX"), (5, "V"), (4, "IV"), (1, "I")]
        res = []
        for v, sym in table:
            k, num = divmod(num, v)
            res.append(sym * k)
        return "".join(res)


if __name__ == "__main__":
    assert Solution().intToRoman(3749) == "MMMDCCXLIX"
    assert Solution().intToRoman(58) == "LVIII"
    assert Solution().intToRoman(1994) == "MCMXCIV"
    print("OK")
