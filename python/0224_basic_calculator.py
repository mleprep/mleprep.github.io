# 224. Basic Calculator (Hard)
# https://leetcode.com/problems/basic-calculator/
# Run: python 0224_basic_calculator.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def calculate(self, s: str) -> int:
        res, num, sign, st = 0, 0, 1, []
        for ch in s:
            if ch.isdigit():
                num = num * 10 + int(ch)
            elif ch in "+-":
                res += sign * num
                num, sign = 0, (1 if ch == "+" else -1)
            elif ch == "(":
                st.append((res, sign))
                res, sign = 0, 1
            elif ch == ")":
                res += sign * num
                num = 0
                prev, psign = st.pop()
                res = prev + psign * res
        return res + sign * num


if __name__ == "__main__":
    assert Solution().calculate("1 + 1") == 2
    assert Solution().calculate(" 2-1 + 2 ") == 3
    assert Solution().calculate("(1+(4+5+2)-3)+(6+8)") == 23
    assert Solution().calculate("-(2+3)") == -5
    print("OK")
