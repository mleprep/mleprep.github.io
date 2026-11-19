# 67. Add Binary (Easy)
# https://leetcode.com/problems/add-binary/
# Run: python 0067_add_binary.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def addBinary(self, a: str, b: str) -> str:
        i, j, carry, out = len(a) - 1, len(b) - 1, 0, []
        while i >= 0 or j >= 0 or carry:
            s = carry + (int(a[i]) if i >= 0 else 0) + (int(b[j]) if j >= 0 else 0)
            out.append(str(s % 2))
            carry = s // 2
            i, j = i - 1, j - 1
        return "".join(reversed(out))


if __name__ == "__main__":
    assert Solution().addBinary("11", "1") == "100"
    assert Solution().addBinary("1010", "1011") == "10101"
    print("OK")
