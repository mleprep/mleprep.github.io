# 336. Palindrome Pairs (Hard)
# https://leetcode.com/problems/palindrome-pairs/
# Run: python 0336_palindrome_pairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        rev = {w[::-1]: i for i, w in enumerate(words)}
        res = []
        for i, w in enumerate(words):
            for j in range(len(w) + 1):
                pre, suf = w[:j], w[j:]
                if pre in rev and rev[pre] != i and suf == suf[::-1]:
                    res.append([i, rev[pre]])
                if j and suf in rev and rev[suf] != i and pre == pre[::-1]:
                    res.append([rev[suf], i])
        return res


if __name__ == "__main__":
    assert sorted(Solution().palindromePairs(["abcd", "dcba", "lls", "s", "sssll"])) == [[0, 1], [1, 0], [2, 4], [3, 2]]
    assert sorted(Solution().palindromePairs(["bat", "tab", "cat"])) == [[0, 1], [1, 0]]
    assert sorted(Solution().palindromePairs(["a", ""])) == [[0, 1], [1, 0]]
    assert sorted(Solution().palindromePairs(["a", "b", "c"])) == []
    assert sorted(Solution().palindromePairs(["ab", "ba", "aba"])) == [[0, 1], [0, 2], [1, 0], [2, 1]]  # "ababa" twice
    print("OK")
