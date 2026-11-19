# 28. Find the Index of the First Occurrence in a String (Easy)
# https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
# Run: python 0028_find_the_index_of_the_first_occurrence_in_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        lps, k = [0] * len(needle), 0
        for i in range(1, len(needle)):
            while k and needle[i] != needle[k]:
                k = lps[k - 1]
            if needle[i] == needle[k]:
                k += 1
            lps[i] = k
        k = 0
        for i, ch in enumerate(haystack):
            while k and ch != needle[k]:
                k = lps[k - 1]
            if ch == needle[k]:
                k += 1
                if k == len(needle):
                    return i - k + 1
        return -1


if __name__ == "__main__":
    assert Solution().strStr("sadbutsad", "sad") == 0
    assert Solution().strStr("leetcode", "leeto") == -1
    assert Solution().strStr("aaaaab", "aab") == 3
    assert Solution().strStr("mississippi", "issip") == 4
    assert Solution().strStr("a", "a") == 0
    assert Solution().strStr("abc", "c") == 2
    print("OK")
