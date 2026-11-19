# 990. Satisfiability of Equality Equations (Medium)
# https://leetcode.com/problems/satisfiability-of-equality-equations/
# Run: python 0990_satisfiability_of_equality_equations.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        parent = list(range(26))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for e in equations:
            if e[1] == "=":
                parent[find(ord(e[0]) - 97)] = find(ord(e[3]) - 97)
        return all(find(ord(e[0]) - 97) != find(ord(e[3]) - 97) for e in equations if e[1] == "!")


if __name__ == "__main__":
    assert Solution().equationsPossible(["a==b", "b!=a"]) is False
    assert Solution().equationsPossible(["b==a", "a==b"]) is True
    assert Solution().equationsPossible(["a==b", "b==c", "a==c"]) is True
    print("OK")
