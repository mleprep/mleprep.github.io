# 1405. Longest Happy String (Medium)
# https://leetcode.com/problems/longest-happy-string/
# Run: python 1405_longest_happy_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        heap = [(-n, ch) for n, ch in ((a, "a"), (b, "b"), (c, "c")) if n]
        heapify(heap)
        out = []
        while heap:
            n1, c1 = heappop(heap)
            if len(out) >= 2 and out[-1] == out[-2] == c1:
                if not heap:
                    break
                n2, c2 = heappop(heap)
                out.append(c2)
                if n2 + 1:
                    heappush(heap, (n2 + 1, c2))
                heappush(heap, (n1, c1))
            else:
                out.append(c1)
                if n1 + 1:
                    heappush(heap, (n1 + 1, c1))
        return "".join(out)


if __name__ == "__main__":
    def happy(s, a, b, c):  # any valid answer of maximum length is accepted
        x, y, z = sorted((a, b, c), reverse=True)
        return ("aaa" not in s and "bbb" not in s and "ccc" not in s and s.count("a") <= a and s.count("b") <= b
                and s.count("c") <= c and len(s) == y + z + min(x, 2 * (y + z + 1)))
    for a, b, c in ((1, 1, 7), (7, 1, 0), (2, 2, 1), (0, 0, 0), (0, 8, 11), (5, 5, 5), (0, 0, 9)):
        assert happy(Solution().longestDiverseString(a, b, c), a, b, c)
    print("OK")
