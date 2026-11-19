# 274. H-Index (Medium)
# https://leetcode.com/problems/h-index/
# Run: python 0274_h_index.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def hIndex(self, citations: List[int]) -> int:
        n = len(citations)
        cnt = [0] * (n + 1)
        for c in citations:
            cnt[min(c, n)] += 1
        papers = 0
        for h in range(n, -1, -1):
            papers += cnt[h]
            if papers >= h:
                return h
        return 0


if __name__ == "__main__":
    assert Solution().hIndex([3, 0, 6, 1, 5]) == 3
    assert Solution().hIndex([1, 3, 1]) == 1
    print("OK")
