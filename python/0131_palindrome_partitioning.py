# 131. Palindrome Partitioning (Medium)
# https://leetcode.com/problems/palindrome-partitioning/
# Run: python 0131_palindrome_partitioning.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res, path = [], []

        def bt(i):
            if i == len(s):
                res.append(path[:])
                return
            for j in range(i, len(s)):
                piece = s[i:j + 1]
                if piece == piece[::-1]:
                    path.append(piece)
                    bt(j + 1)
                    path.pop()

        bt(0)
        return res


if __name__ == "__main__":
    assert sorted(Solution().partition("aab")) == sorted([["a", "a", "b"], ["aa", "b"]])
    assert Solution().partition("a") == [["a"]]
    assert sorted(Solution().partition("aba")) == [["a", "b", "a"], ["aba"]]
    print("OK")
