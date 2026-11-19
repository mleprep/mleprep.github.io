# 470. Implement Rand10() Using Rand7() (Medium)
# https://leetcode.com/problems/implement-rand10-using-rand7/
# Run: python 0470_implement_rand10_using_rand7.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rand10(self) -> int:
        while True:
            v = (rand7() - 1) * 7 + rand7()     # rand7 provided by LeetCode
            if v <= 40:
                return (v - 1) % 10 + 1


if __name__ == "__main__":
    random.seed(7)
    def rand7(): return random.randint(1, 7)
    cnt = Counter(Solution().rand10() for _ in range(20000))
    assert set(cnt) == set(range(1, 11)) and all(1800 < c < 2200 for c in cnt.values())
    print("OK")
