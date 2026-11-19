# 354. Russian Doll Envelopes (Hard)
# https://leetcode.com/problems/russian-doll-envelopes/
# Run: python 0354_russian_doll_envelopes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:
        envelopes.sort(key=lambda e: (e[0], -e[1]))
        tails = []
        for _, h in envelopes:
            i = bisect.bisect_left(tails, h)
            if i == len(tails):
                tails.append(h)
            else:
                tails[i] = h
        return len(tails)


if __name__ == "__main__":
    assert Solution().maxEnvelopes([[5, 4], [6, 4], [6, 7], [2, 3]]) == 3
    assert Solution().maxEnvelopes([[1, 1], [1, 1], [1, 1]]) == 1
    print("OK")
