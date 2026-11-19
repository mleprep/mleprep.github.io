# 8. String to Integer (atoi) (Medium)
# https://leetcode.com/problems/string-to-integer-atoi/
# Run: python 0008_string_to_integer_atoi.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def myAtoi(self, s: str) -> int:
        i, n = 0, len(s)
        while i < n and s[i] == " ":
            i += 1
        sign = 1
        if i < n and s[i] in "+-":
            sign = -1 if s[i] == "-" else 1
            i += 1
        num = 0
        while i < n and s[i].isdigit():
            num = num * 10 + int(s[i])
            i += 1
        return max(-2 ** 31, min(2 ** 31 - 1, sign * num))


if __name__ == "__main__":
    assert Solution().myAtoi("42") == 42
    assert Solution().myAtoi("   -042") == -42
    assert Solution().myAtoi("1337c0d3") == 1337
    assert Solution().myAtoi("0-1") == 0
    assert Solution().myAtoi("-91283472332") == -2147483648
    print("OK")
