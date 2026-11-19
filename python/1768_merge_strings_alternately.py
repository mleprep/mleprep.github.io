# 1768. Merge Strings Alternately (Easy)
# https://leetcode.com/problems/merge-strings-alternately/
# Run: python 1768_merge_strings_alternately.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        return "".join(a + b for a, b in itertools.zip_longest(word1, word2, fillvalue=""))


if __name__ == "__main__":
    assert Solution().mergeAlternately("abc", "pqr") == "apbqcr"
    assert Solution().mergeAlternately("ab", "pqrs") == "apbqrs"
    assert Solution().mergeAlternately("abcd", "pq") == "apbqcd"
    print("OK")
