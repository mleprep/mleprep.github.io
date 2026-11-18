# 706. Design HashMap (Easy)
# https://leetcode.com/problems/design-hashmap/
# Run: python 0706_design_hashmap.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyHashMap:
    def __init__(self):
        self.size = 1009                       # prime number of buckets
        self.b = [[] for _ in range(self.size)]

    def put(self, key: int, value: int) -> None:
        bucket = self.b[key % self.size]
        for pair in bucket:
            if pair[0] == key:
                pair[1] = value
                return
        bucket.append([key, value])

    def get(self, key: int) -> int:
        for k, v in self.b[key % self.size]:
            if k == key:
                return v
        return -1

    def remove(self, key: int) -> None:
        bucket = self.b[key % self.size]
        for i, (k, _) in enumerate(bucket):
            if k == key:
                bucket.pop(i)
                return


if __name__ == "__main__":
    hm = MyHashMap()
    hm.put(1, 1); hm.put(2, 2)
    assert hm.get(1) == 1 and hm.get(3) == -1
    hm.put(2, 1); assert hm.get(2) == 1
    hm.remove(2); assert hm.get(2) == -1
    hm.put(1009 + 1, 7); assert hm.get(1) == 1 and hm.get(1010) == 7   # same bucket
    print("OK")
