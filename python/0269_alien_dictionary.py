# 269. Alien Dictionary (Hard)
# https://leetcode.com/problems/alien-dictionary/
# Run: python 0269_alien_dictionary.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def alienOrder(self, words: List[str]) -> str:
        g = {ch: set() for w in words for ch in w}
        indeg = {ch: 0 for ch in g}
        for a, b in zip(words, words[1:]):
            for x, y in zip(a, b):
                if x != y:
                    if y not in g[x]:
                        g[x].add(y)
                        indeg[y] += 1
                    break
            else:
                if len(a) > len(b):
                    return ""
        q = deque(sorted(ch for ch in indeg if indeg[ch] == 0))
        out = []
        while q:
            u = q.popleft()
            out.append(u)
            for v in sorted(g[u]):
                indeg[v] -= 1
                if indeg[v] == 0:
                    q.append(v)
        return "".join(out) if len(out) == len(g) else ""


if __name__ == "__main__":
    assert Solution().alienOrder(["wrt", "wrf", "er", "ett", "rftt"]) == "wertf"
    assert Solution().alienOrder(["z", "x", "z"]) == ""
    assert Solution().alienOrder(["abc", "ab"]) == ""
    print("OK")
