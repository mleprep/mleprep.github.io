# 424. Longest Repeating Character Replacement (Medium)
# https://leetcode.com/problems/longest-repeating-character-replacement/
# Run: python 0424_longest_repeating_character_replacement.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        cnt = Counter()
        l = maxf = 0
        for r, ch in enumerate(s):
            cnt[ch] += 1
            maxf = max(maxf, cnt[ch])
            if r - l + 1 - maxf > k:
                cnt[s[l]] -= 1
                l += 1
        return len(s) - l


if __name__ == "__main__":
    assert Solution().characterReplacement("ABAB", 2) == 4
    assert Solution().characterReplacement("AABABBA", 1) == 4
    print("OK")
