# 132. Palindrome Partitioning II (Hard)
# https://leetcode.com/problems/palindrome-partitioning-ii/
# Run: python 0132_palindrome_partitioning_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minCut(self, s: str) -> int:
        n = len(s)
        cuts = [i - 1 for i in range(n + 1)]       # worst case: cut between every char
        for c in range(n):
            for l, r in ((c, c), (c, c + 1)):
                while l >= 0 and r < n and s[l] == s[r]:
                    cuts[r + 1] = min(cuts[r + 1], cuts[l] + 1)
                    l, r = l - 1, r + 1
        return cuts[n]


if __name__ == "__main__":
    assert Solution().minCut("aab") == 1
    assert Solution().minCut("a") == 0
    assert Solution().minCut("abccbc") == 2
    print("OK")
