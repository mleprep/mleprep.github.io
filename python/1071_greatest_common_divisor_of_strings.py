# 1071. Greatest Common Divisor of Strings (Easy)
# https://leetcode.com/problems/greatest-common-divisor-of-strings/
# Run: python 1071_greatest_common_divisor_of_strings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        if str1 + str2 != str2 + str1:
            return ""
        return str1[:math.gcd(len(str1), len(str2))]


if __name__ == "__main__":
    assert Solution().gcdOfStrings("ABCABC", "ABC") == "ABC"
    assert Solution().gcdOfStrings("ABABAB", "ABAB") == "AB"
    assert Solution().gcdOfStrings("LEET", "CODE") == ""
    assert Solution().gcdOfStrings("ABABABAB", "ABAB") == "ABAB"
    assert Solution().gcdOfStrings("AAAA", "AA") == "AA"
    print("OK")
