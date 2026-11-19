# 2707. Extra Characters in a String (Medium)
# https://leetcode.com/problems/extra-characters-in-a-string/
# Run: python 2707_extra_characters_in_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        words, n = set(dictionary), len(s)
        dp = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            dp[i] = 1 + dp[i + 1]
            for j in range(i + 1, n + 1):
                if s[i:j] in words:
                    dp[i] = min(dp[i], dp[j])
        return dp[0]


if __name__ == "__main__":
    assert Solution().minExtraChar("leetscode", ["leet", "code", "leetcode"]) == 1
    assert Solution().minExtraChar("sayhelloworld", ["hello", "world"]) == 3
    assert Solution().minExtraChar("abc", ["d"]) == 3
    assert Solution().minExtraChar("abc", ["abc", "a"]) == 0
    print("OK")
