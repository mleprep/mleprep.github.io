# 753. Cracking the Safe (Hard)
# https://leetcode.com/problems/cracking-the-safe/
# Run: python 0753_cracking_the_safe.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def crackSafe(self, n: int, k: int) -> str:
        seen, out = set(), []

        def dfs(node):
            for d in map(str, range(k)):
                edge = node + d
                if edge not in seen:
                    seen.add(edge)
                    dfs(edge[1:])
                    out.append(d)

        start = "0" * (n - 1)
        dfs(start)
        return "".join(out) + start


if __name__ == "__main__":
    def covers(s, n, k):
        return len({s[i:i + n] for i in range(len(s) - n + 1)}) == k ** n and len(s) == k ** n + n - 1
    assert covers(Solution().crackSafe(1, 2), 1, 2)
    assert covers(Solution().crackSafe(2, 2), 2, 2)
    assert covers(Solution().crackSafe(3, 3), 3, 3)
    print("OK")
