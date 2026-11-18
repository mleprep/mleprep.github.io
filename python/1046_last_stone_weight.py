# 1046. Last Stone Weight (Easy)
# https://leetcode.com/problems/last-stone-weight/
# Run: python 1046_last_stone_weight.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        h = [-s for s in stones]
        heapify(h)
        while len(h) > 1:
            a, b = -heappop(h), -heappop(h)
            if a != b:
                heappush(h, -(a - b))
        return -h[0] if h else 0


if __name__ == "__main__":
    assert Solution().lastStoneWeight([2, 7, 4, 1, 8, 1]) == 1
    assert Solution().lastStoneWeight([1]) == 1
    print("OK")
