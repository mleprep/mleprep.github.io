# 58. Length of Last Word (Easy)
# https://leetcode.com/problems/length-of-last-word/
# Run: python 0058_length_of_last_word.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        i = len(s) - 1
        while i >= 0 and s[i] == " ":
            i -= 1
        j = i
        while j >= 0 and s[j] != " ":
            j -= 1
        return i - j


if __name__ == "__main__":
    assert Solution().lengthOfLastWord("Hello World") == 5
    assert Solution().lengthOfLastWord("   fly me   to   the moon  ") == 4
    assert Solution().lengthOfLastWord("luffy is still joyboy") == 6
    assert Solution().lengthOfLastWord("a") == 1
    print("OK")
