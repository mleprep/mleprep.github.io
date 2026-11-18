# 242. Valid Anagram (Easy)
# https://leetcode.com/problems/valid-anagram/
# Run: python 0242_valid_anagram.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        cnt = [0] * 26
        for a, b in zip(s, t):
            cnt[ord(a) - 97] += 1
            cnt[ord(b) - 97] -= 1
        return all(c == 0 for c in cnt)


if __name__ == "__main__":
    assert Solution().isAnagram("anagram", "nagaram") is True
    assert Solution().isAnagram("rat", "car") is False
    print("OK")
