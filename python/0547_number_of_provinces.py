# 547. Number of Provinces (Medium)
# https://leetcode.com/problems/number-of-provinces/
# Run: python 0547_number_of_provinces.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        n, seen = len(isConnected), set()

        def dfs(u):
            for v in range(n):
                if isConnected[u][v] and v not in seen:
                    seen.add(v)
                    dfs(v)

        count = 0
        for u in range(n):
            if u not in seen:
                seen.add(u)
                dfs(u)
                count += 1
        return count


if __name__ == "__main__":
    assert Solution().findCircleNum([[1, 1, 0], [1, 1, 0], [0, 0, 1]]) == 2
    assert Solution().findCircleNum([[1, 0, 0], [0, 1, 0], [0, 0, 1]]) == 3
    print("OK")
