# 332. Reconstruct Itinerary (Hard)
# https://leetcode.com/problems/reconstruct-itinerary/
# Run: python 0332_reconstruct_itinerary.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findItinerary(self, tickets: List[List[str]]) -> List[str]:
        g = defaultdict(list)
        for a, b in sorted(tickets, reverse=True):
            g[a].append(b)               # reverse-sorted so pop() gives smallest
        route, st = [], ["JFK"]
        while st:
            while g[st[-1]]:
                st.append(g[st[-1]].pop())
            route.append(st.pop())
        return route[::-1]


if __name__ == "__main__":
    T = [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
    assert Solution().findItinerary(T) == ["JFK", "MUC", "LHR", "SFO", "SJC"]
    assert Solution().findItinerary([["JFK", "KUL"], ["JFK", "NRT"], ["NRT", "JFK"]]) == ["JFK", "NRT", "JFK", "KUL"]
    print("OK")
