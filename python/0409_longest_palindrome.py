# 409. Longest Palindrome (Easy)
# https://leetcode.com/problems/longest-palindrome/
# Run: python 0409_longest_palindrome.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestPalindrome(self, s: str) -> int:
        pairs = sum(v // 2 * 2 for v in Counter(s).values())
        return pairs + 1 if pairs < len(s) else pairs


if __name__ == "__main__":
    assert Solution().longestPalindrome("abccccdd") == 7
    assert Solution().longestPalindrome("a") == 1
    assert Solution().longestPalindrome("Aa") == 1
    assert Solution().longestPalindrome("bb") == 2
    print("OK")
