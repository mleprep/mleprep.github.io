# 721. Accounts Merge (Medium)
# https://leetcode.com/problems/accounts-merge/
# Run: python 0721_accounts_merge.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def accountsMerge(self, accounts: List[List[str]]) -> List[List[str]]:
        parent, owner = {}, {}

        def find(x):
            parent.setdefault(x, x)
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for name, *emails in accounts:
            for e in emails:
                owner[e] = name
                parent[find(e)] = find(emails[0])
        groups = defaultdict(list)
        for e in owner:
            groups[find(e)].append(e)
        return [[owner[root]] + sorted(es) for root, es in groups.items()]


if __name__ == "__main__":
    acc = [["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["John", "johnsmith@mail.com", "john00@mail.com"],
           ["Mary", "mary@mail.com"], ["John", "johnnybravo@mail.com"]]
    assert sorted(Solution().accountsMerge(acc)) == [
        ["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"],
        ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
    print("OK")
