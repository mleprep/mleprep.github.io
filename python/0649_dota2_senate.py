# 649. Dota2 Senate (Medium)
# https://leetcode.com/problems/dota2-senate/
# Run: python 0649_dota2_senate.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        n = len(senate)
        r = deque(i for i, c in enumerate(senate) if c == "R")
        d = deque(i for i, c in enumerate(senate) if c == "D")
        while r and d:
            a, b = r.popleft(), d.popleft()
            if a < b:
                r.append(a + n)
            else:
                d.append(b + n)
        return "Radiant" if r else "Dire"


if __name__ == "__main__":
    assert Solution().predictPartyVictory("RD") == "Radiant"
    assert Solution().predictPartyVictory("RDD") == "Dire"
    print("OK")
