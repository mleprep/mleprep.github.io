# 705. Design HashSet (Easy)
# https://leetcode.com/problems/design-hashset/
# Run: python 0705_design_hashset.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyHashSet:
    def __init__(self):
        self.size = 1009
        self.buckets = [[] for _ in range(self.size)]

    def add(self, key: int) -> None:
        b = self.buckets[key % self.size]
        if key not in b:
            b.append(key)

    def remove(self, key: int) -> None:
        b = self.buckets[key % self.size]
        if key in b:
            b.remove(key)

    def contains(self, key: int) -> bool:
        return key in self.buckets[key % self.size]


if __name__ == "__main__":
    hs = MyHashSet()
    hs.add(1); hs.add(2)
    assert hs.contains(1) and not hs.contains(3)
    hs.add(2)
    assert hs.contains(2)
    hs.remove(2)
    assert not hs.contains(2)
    hs.add(0); hs.add(1009); hs.add(2018)                 # same bucket
    hs.remove(1009)
    assert hs.contains(0) and not hs.contains(1009) and hs.contains(2018)
    print("OK")
