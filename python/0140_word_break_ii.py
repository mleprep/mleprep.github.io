# 140. Word Break II (Hard)
# https://leetcode.com/problems/word-break-ii/
# Run: python 0140_word_break_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        words = set(wordDict)

        @cache
        def go(i):
            if i == len(s):
                return [""]
            out = []
            for j in range(i + 1, len(s) + 1):
                if s[i:j] in words:
                    out += [s[i:j] + (" " + rest if rest else "") for rest in go(j)]
            return out

        return go(0)


if __name__ == "__main__":
    assert sorted(Solution().wordBreak("catsanddog", ["cat", "cats", "and", "sand", "dog"])) == ["cat sand dog", "cats and dog"]
    assert sorted(Solution().wordBreak("pineapplepenapple", ["apple", "pen", "applepen", "pine", "pineapple"])) == [
        "pine apple pen apple", "pine applepen apple", "pineapple pen apple"]
    assert Solution().wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]) == []
    print("OK")
