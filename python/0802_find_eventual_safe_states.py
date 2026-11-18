# 802. Find Eventual Safe States (Medium)
# https://leetcode.com/problems/find-eventual-safe-states/
# Run: python 0802_find_eventual_safe_states.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        color = [0] * len(graph)

        def safe(u):
            if color[u]:
                return color[u] == 2
            color[u] = 1
            if all(safe(v) for v in graph[u]):
                color[u] = 2
                return True
            return False          # stays 1 -> unsafe

        return [u for u in range(len(graph)) if safe(u)]


if __name__ == "__main__":
    assert Solution().eventualSafeNodes([[1, 2], [2, 3], [5], [0], [5], [], []]) == [2, 4, 5, 6]
    assert Solution().eventualSafeNodes([[1, 2, 3, 4], [1, 2], [3, 4], [0, 4], []]) == [4]
    print("OK")
