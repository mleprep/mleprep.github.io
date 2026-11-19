# 1312. Minimum Insertion Steps to Make a String Palindrome (Hard)
# https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
# Run: python 1312_minimum_insertion_steps_to_make_a_string_palindrome.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minInsertions(self, s: str) -> int:
        n = len(s)
        dp = [0] * n
        for l in range(n - 1, -1, -1):
            dp[l], diag = 1, 0
            for r in range(l + 1, n):
                tmp = dp[r]
                dp[r] = diag + 2 if s[l] == s[r] else max(dp[r], dp[r - 1])
                diag = tmp
        return n - dp[n - 1]


if __name__ == "__main__":
    assert Solution().minInsertions("zzazz") == 0
    assert Solution().minInsertions("mbadm") == 2
    assert Solution().minInsertions("leetcode") == 5
    print("OK")
