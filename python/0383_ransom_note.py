# 383. Ransom Note (Easy)
# https://leetcode.com/problems/ransom-note/
# Run: python 0383_ransom_note.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        have = Counter(magazine)
        return all(have[c] >= k for c, k in Counter(ransomNote).items())


if __name__ == "__main__":
    assert Solution().canConstruct("a", "b") is False
    assert Solution().canConstruct("aa", "ab") is False
    assert Solution().canConstruct("aa", "aab") is True
    assert Solution().canConstruct("", "x") is True
    print("OK")
