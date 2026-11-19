# 1049. Last Stone Weight II (Medium)
# https://leetcode.com/problems/last-stone-weight-ii/
# Run: python 1049_last_stone_weight_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lastStoneWeightII(self, stones: List[int]) -> int:
        total, reach = sum(stones), 1               # bit s set <=> some subset sums to s
        for x in stones:
            reach |= reach << x
        best = max(s for s in range(total // 2 + 1) if reach >> s & 1)
        return total - 2 * best


if __name__ == "__main__":
    assert Solution().lastStoneWeightII([2, 7, 4, 1, 8, 1]) == 1
    assert Solution().lastStoneWeightII([31, 26, 33, 21, 40]) == 5
    assert Solution().lastStoneWeightII([1]) == 1
    assert Solution().lastStoneWeightII([1, 1]) == 0
    print("OK")
