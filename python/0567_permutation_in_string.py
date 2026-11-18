# 567. Permutation in String (Medium)
# https://leetcode.com/problems/permutation-in-string/
# Run: python 0567_permutation_in_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        k = len(s1)
        if k > len(s2):
            return False
        need, win = Counter(s1), Counter(s2[:k])
        if need == win:
            return True
        for r in range(k, len(s2)):
            win[s2[r]] += 1
            win[s2[r - k]] -= 1
            if win[s2[r - k]] == 0:
                del win[s2[r - k]]
            if win == need:
                return True
        return False


if __name__ == "__main__":
    assert Solution().checkInclusion("ab", "eidbaooo") is True
    assert Solution().checkInclusion("ab", "eidboaoo") is False
    print("OK")
