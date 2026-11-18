# 438. Find All Anagrams in a String (Medium)
# https://leetcode.com/problems/find-all-anagrams-in-a-string/
# Run: python 0438_find_all_anagrams_in_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        k, res = len(p), []
        need, win = [0] * 26, [0] * 26
        for ch in p:
            need[ord(ch) - 97] += 1
        for r, ch in enumerate(s):
            win[ord(ch) - 97] += 1
            if r >= k:
                win[ord(s[r - k]) - 97] -= 1
            if win == need:
                res.append(r - k + 1)
        return res


if __name__ == "__main__":
    assert Solution().findAnagrams("cbaebabacd", "abc") == [0, 6]
    assert Solution().findAnagrams("abab", "ab") == [0, 1, 2]
    print("OK")
