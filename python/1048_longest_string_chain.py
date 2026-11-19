# 1048. Longest String Chain (Medium)
# https://leetcode.com/problems/longest-string-chain/
# Run: python 1048_longest_string_chain.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        dp = {}
        for w in sorted(words, key=len):
            dp[w] = 1 + max((dp.get(w[:i] + w[i + 1:], 0) for i in range(len(w))), default=0)
        return max(dp.values())


if __name__ == "__main__":
    assert Solution().longestStrChain(["a", "b", "ba", "bca", "bda", "bdca"]) == 4
    assert Solution().longestStrChain(["xbc", "pcxbcf", "xb", "cxbc", "pcxbc"]) == 5
    assert Solution().longestStrChain(["abcd", "dbqca"]) == 1
    print("OK")
