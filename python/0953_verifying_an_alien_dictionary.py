# 953. Verifying an Alien Dictionary (Easy)
# https://leetcode.com/problems/verifying-an-alien-dictionary/
# Run: python 0953_verifying_an_alien_dictionary.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        rank = {ch: i for i, ch in enumerate(order)}
        keys = [[rank[ch] for ch in w] for w in words]
        return all(a <= b for a, b in zip(keys, keys[1:]))   # list comparison = lexicographic


if __name__ == "__main__":
    assert Solution().isAlienSorted(["hello", "leetcode"], "hlabcdefgijkmnopqrstuvwxyz") is True
    assert Solution().isAlienSorted(["word", "world", "row"], "worldabcefghijkmnpqstuvxyz") is False
    assert Solution().isAlienSorted(["apple", "app"], "abcdefghijklmnopqrstuvwxyz") is False
    print("OK")
