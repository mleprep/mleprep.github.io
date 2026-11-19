# 886. Possible Bipartition (Medium)
# https://leetcode.com/problems/possible-bipartition/
# Run: python 0886_possible_bipartition.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        g = defaultdict(list)
        for a, b in dislikes:
            g[a].append(b)
            g[b].append(a)
        color = {}
        for s in range(1, n + 1):
            if s in color:
                continue
            color[s] = 0
            q = deque([s])
            while q:
                u = q.popleft()
                for v in g[u]:
                    if v not in color:
                        color[v] = 1 - color[u]
                        q.append(v)
                    elif color[v] == color[u]:
                        return False
        return True


if __name__ == "__main__":
    assert Solution().possibleBipartition(4, [[1, 2], [1, 3], [2, 4]]) is True
    assert Solution().possibleBipartition(3, [[1, 2], [1, 3], [2, 3]]) is False
    print("OK")
