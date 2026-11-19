# 345. Reverse Vowels of a String (Easy)
# https://leetcode.com/problems/reverse-vowels-of-a-string/
# Run: python 0345_reverse_vowels_of_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = set("aeiouAEIOU")
        a, i, j = list(s), 0, len(s) - 1
        while i < j:
            if a[i] not in vowels:
                i += 1
            elif a[j] not in vowels:
                j -= 1
            else:
                a[i], a[j] = a[j], a[i]
                i, j = i + 1, j - 1
        return "".join(a)


if __name__ == "__main__":
    assert Solution().reverseVowels("IceCreAm") == "AceCreIm"
    assert Solution().reverseVowels("leetcode") == "leotcede"
    assert Solution().reverseVowels("bcd") == "bcd"
    assert Solution().reverseVowels("aA") == "Aa"
    print("OK")
