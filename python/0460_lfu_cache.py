# 460. LFU Cache (Hard)
# https://leetcode.com/problems/lfu-cache/
# Run: python 0460_lfu_cache.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class LFUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.kv = {}                                 # key -> [value, freq]
        self.buckets = defaultdict(OrderedDict)      # freq -> keys in LRU order
        self.min_f = 0

    def _touch(self, key):
        val, f = self.kv[key]
        del self.buckets[f][key]
        if not self.buckets[f]:
            del self.buckets[f]
            if self.min_f == f:
                self.min_f += 1
        self.kv[key][1] = f + 1
        self.buckets[f + 1][key] = None

    def get(self, key: int) -> int:
        if key not in self.kv:
            return -1
        self._touch(key)
        return self.kv[key][0]

    def put(self, key: int, value: int) -> None:
        if self.cap == 0:
            return
        if key in self.kv:
            self.kv[key][0] = value
            self._touch(key)
            return
        if len(self.kv) == self.cap:
            old, _ = self.buckets[self.min_f].popitem(last=False)
            del self.kv[old]
        self.kv[key] = [value, 1]
        self.buckets[1][key] = None
        self.min_f = 1


if __name__ == "__main__":
    c = LFUCache(2)
    c.put(1, 1); c.put(2, 2)
    assert c.get(1) == 1
    c.put(3, 3)                       # evicts 2 (freq 1)
    assert c.get(2) == -1 and c.get(3) == 3
    c.put(4, 4)                       # 1 and 3 both freq 2 -> evict LRU among them = 1
    assert c.get(1) == -1 and c.get(3) == 3 and c.get(4) == 4
    print("OK")
