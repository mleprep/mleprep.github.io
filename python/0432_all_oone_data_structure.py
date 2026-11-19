# 432. All O`one Data Structure (Hard)
# https://leetcode.com/problems/all-oone-data-structure/
# Run: python 0432_all_oone_data_structure.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Bucket:
    def __init__(self, cnt):
        self.cnt, self.keys, self.prev, self.next = cnt, set(), None, None

class AllOne:
    def __init__(self):
        self.head, self.tail = Bucket(0), Bucket(0)
        self.head.next, self.tail.prev = self.tail, self.head
        self.where = {}

    def _after(self, node, cnt):
        b = Bucket(cnt)
        b.prev, b.next = node, node.next
        node.next.prev = b
        node.next = b
        return b

    def _drop_if_empty(self, b):
        if not b.keys:
            b.prev.next, b.next.prev = b.next, b.prev

    def inc(self, key: str) -> None:
        cur = self.where.get(key, self.head)
        nxt = cur.next
        if nxt is self.tail or nxt.cnt != cur.cnt + 1:
            nxt = self._after(cur, cur.cnt + 1)
        nxt.keys.add(key)
        self.where[key] = nxt
        if cur is not self.head:
            cur.keys.discard(key)
            self._drop_if_empty(cur)

    def dec(self, key: str) -> None:
        cur = self.where[key]
        if cur.cnt == 1:
            del self.where[key]
        else:
            prv = cur.prev
            if prv is self.head or prv.cnt != cur.cnt - 1:
                prv = self._after(cur.prev, cur.cnt - 1)
            prv.keys.add(key)
            self.where[key] = prv
        cur.keys.discard(key)
        self._drop_if_empty(cur)

    def getMaxKey(self) -> str:
        return next(iter(self.tail.prev.keys)) if self.tail.prev is not self.head else ""

    def getMinKey(self) -> str:
        return next(iter(self.head.next.keys)) if self.head.next is not self.tail else ""


if __name__ == "__main__":
    a = AllOne()
    a.inc("hello"); a.inc("hello")
    assert a.getMaxKey() == "hello" and a.getMinKey() == "hello"
    a.inc("leet")
    assert a.getMaxKey() == "hello" and a.getMinKey() == "leet"
    a.dec("hello"); a.dec("hello")
    assert a.getMaxKey() == "leet" and a.getMinKey() == "leet"
    a.dec("leet")
    assert a.getMaxKey() == "" and a.getMinKey() == ""
    print("OK")
