# 1456. Maximum Number of Vowels in a Substring of Given Length (Medium)
# https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
# Run: python 1456_maximum_number_of_vowels_in_a_substring_of_given_length.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = set("aeiou")
        cur = best = sum(c in vowels for c in s[:k])
        for i in range(k, len(s)):
            cur += (s[i] in vowels) - (s[i - k] in vowels)
            best = max(best, cur)
        return best


if __name__ == "__main__":
    assert Solution().maxVowels("abciiidef", 3) == 3
    assert Solution().maxVowels("aeiou", 2) == 2
    assert Solution().maxVowels("leetcode", 3) == 2
    assert Solution().maxVowels("tryhard", 4) == 1
    print("OK")
