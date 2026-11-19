# 583. Delete Operation for Two Strings (Medium)
# https://leetcode.com/problems/delete-operation-for-two-strings/
# Run: python 0583_delete_operation_for_two_strings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        n = len(word2)
        dp = [0] * (n + 1)
        for a in word1:
            diag = 0
            for j in range(1, n + 1):
                up = dp[j]
                dp[j] = diag + 1 if a == word2[j - 1] else max(dp[j], dp[j - 1])
                diag = up
        return len(word1) + n - 2 * dp[n]


if __name__ == "__main__":
    assert Solution().minDistance("sea", "eat") == 2
    assert Solution().minDistance("leetcode", "etco") == 4
    print("OK")
