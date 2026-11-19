# 443. String Compression (Medium)
# https://leetcode.com/problems/string-compression/
# Run: python 0443_string_compression.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def compress(self, chars: List[str]) -> int:
        w = i = 0
        while i < len(chars):
            j = i
            while j < len(chars) and chars[j] == chars[i]:
                j += 1
            chars[w] = chars[i]
            w += 1
            if j - i > 1:
                for d in str(j - i):
                    chars[w] = d
                    w += 1
            i = j
        return w


if __name__ == "__main__":
    a = ["a", "a", "b", "b", "c", "c", "c"]
    k = Solution().compress(a)
    assert k == 6 and a[:k] == ["a", "2", "b", "2", "c", "3"]
    a = ["a"]
    k = Solution().compress(a)
    assert k == 1 and a[:k] == ["a"]
    a = ["a"] + ["b"] * 12
    k = Solution().compress(a)
    assert k == 4 and a[:k] == ["a", "b", "1", "2"]
    print("OK")
