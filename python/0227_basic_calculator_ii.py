# 227. Basic Calculator II (Medium)
# https://leetcode.com/problems/basic-calculator-ii/
# Run: python 0227_basic_calculator_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def calculate(self, s: str) -> int:
        st, num, op = [], 0, "+"
        for i, ch in enumerate(s):
            if ch.isdigit():
                num = num * 10 + int(ch)
            if ch in "+-*/" or i == len(s) - 1:
                if op == "+": st.append(num)
                elif op == "-": st.append(-num)
                elif op == "*": st.append(st.pop() * num)
                else: st.append(int(st.pop() / num))
                op, num = ch, 0
        return sum(st)


if __name__ == "__main__":
    assert Solution().calculate("3+2*2") == 7
    assert Solution().calculate(" 3/2 ") == 1
    assert Solution().calculate(" 3+5 / 2 ") == 5
    assert Solution().calculate("14-3/2") == 13
    print("OK")
