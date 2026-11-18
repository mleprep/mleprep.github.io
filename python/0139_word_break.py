# 139. Word Break (Medium)
# https://leetcode.com/problems/word-break/
# Run: python 0139_word_break.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        words = set(wordDict)
        L = max(map(len, words))
        dp = [True] + [False] * len(s)
        for i in range(1, len(s) + 1):
            for j in range(max(0, i - L), i):
                if dp[j] and s[j:i] in words:
                    dp[i] = True
                    break
        return dp[-1]


if __name__ == "__main__":
    assert Solution().wordBreak("leetcode", ["leet", "code"]) is True
    assert Solution().wordBreak("applepenapple", ["apple", "pen"]) is True
    assert Solution().wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]) is False
    print("OK")
