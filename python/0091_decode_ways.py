# 91. Decode Ways (Medium)
# https://leetcode.com/problems/decode-ways/
# Run: python 0091_decode_ways.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numDecodings(self, s: str) -> int:
        a, b = 1, 1 if s[0] != "0" else 0      # dp[i-2], dp[i-1]
        for i in range(2, len(s) + 1):
            cur = b if s[i - 1] != "0" else 0
            if 10 <= int(s[i - 2:i]) <= 26:
                cur += a
            a, b = b, cur
        return b


if __name__ == "__main__":
    assert Solution().numDecodings("12") == 2
    assert Solution().numDecodings("226") == 3
    assert Solution().numDecodings("06") == 0
    assert Solution().numDecodings("10") == 1
    print("OK")
