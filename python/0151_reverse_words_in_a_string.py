# 151. Reverse Words in a String (Medium)
# https://leetcode.com/problems/reverse-words-in-a-string/
# Run: python 0151_reverse_words_in_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(reversed(s.split()))


if __name__ == "__main__":
    assert Solution().reverseWords("the sky is blue") == "blue is sky the"
    assert Solution().reverseWords("  hello world  ") == "world hello"
    assert Solution().reverseWords("a good   example") == "example good a"
    print("OK")
