# 680. Valid Palindrome II (Easy)
# https://leetcode.com/problems/valid-palindrome-ii/
# Run: python 0680_valid_palindrome_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def validPalindrome(self, s: str) -> bool:
        def pal(l, r):
            while l < r:
                if s[l] != s[r]:
                    return False
                l, r = l + 1, r - 1
            return True

        l, r = 0, len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return pal(l + 1, r) or pal(l, r - 1)
            l, r = l + 1, r - 1
        return True


if __name__ == "__main__":
    assert Solution().validPalindrome("aba") is True
    assert Solution().validPalindrome("abca") is True
    assert Solution().validPalindrome("abc") is False
    print("OK")
