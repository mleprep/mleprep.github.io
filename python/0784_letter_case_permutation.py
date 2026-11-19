# 784. Letter Case Permutation (Medium)
# https://leetcode.com/problems/letter-case-permutation/
# Run: python 0784_letter_case_permutation.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def letterCasePermutation(self, s: str) -> List[str]:
        res = [""]
        for ch in s:
            if ch.isalpha():
                res = [p + c for p in res for c in (ch.lower(), ch.upper())]
            else:
                res = [p + ch for p in res]
        return res


if __name__ == "__main__":
    assert sorted(Solution().letterCasePermutation("a1b2")) == ["A1B2", "A1b2", "a1B2", "a1b2"]
    assert Solution().letterCasePermutation("3z4") == ["3z4", "3Z4"]
    print("OK")
