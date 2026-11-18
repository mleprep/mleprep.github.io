# 43. Multiply Strings (Medium)
# https://leetcode.com/problems/multiply-strings/
# Run: python 0043_multiply_strings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        m, n = len(num1), len(num2)
        pos = [0] * (m + n)
        for i in range(m - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                s = int(num1[i]) * int(num2[j]) + pos[i + j + 1]
                pos[i + j + 1] = s % 10
                pos[i + j] += s // 10
        res = "".join(map(str, pos)).lstrip("0")
        return res or "0"


if __name__ == "__main__":
    assert Solution().multiply("2", "3") == "6"
    assert Solution().multiply("123", "456") == "56088"
    assert Solution().multiply("0", "52") == "0"
    print("OK")
