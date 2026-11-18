# 146. LRU Cache (Medium)
# https://leetcode.com/problems/lru-cache/
# Run: python 0146_lru_cache.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.od = OrderedDict()      # hash map + doubly linked list in one

    def get(self, key: int) -> int:
        if key not in self.od:
            return -1
        self.od.move_to_end(key)     # most recent at the end
        return self.od[key]

    def put(self, key: int, value: int) -> None:
        self.od[key] = value
        self.od.move_to_end(key)
        if len(self.od) > self.cap:
            self.od.popitem(last=False)   # evict least recent


if __name__ == "__main__":
    c = LRUCache(2)
    c.put(1, 1); c.put(2, 2)
    assert c.get(1) == 1
    c.put(3, 3)                  # evicts 2
    assert c.get(2) == -1
    c.put(4, 4)                  # evicts 1
    assert c.get(1) == -1 and c.get(3) == 3 and c.get(4) == 4
    print("OK")
