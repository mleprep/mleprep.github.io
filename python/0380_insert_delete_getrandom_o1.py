# 380. Insert Delete GetRandom O(1) (Medium)
# https://leetcode.com/problems/insert-delete-getrandom-o1/
# Run: python 0380_insert_delete_getrandom_o1.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class RandomizedSet:
    def __init__(self):
        self.vals, self.idx = [], {}

    def insert(self, val: int) -> bool:
        if val in self.idx:
            return False
        self.idx[val] = len(self.vals)
        self.vals.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.idx:
            return False
        i, last = self.idx.pop(val), self.vals.pop()
        if i < len(self.vals):            # removed element wasn't the last one
            self.vals[i] = last
            self.idx[last] = i
        return True

    def getRandom(self) -> int:
        return random.choice(self.vals)


if __name__ == "__main__":
    rs = RandomizedSet()
    assert rs.insert(1) and not rs.remove(2) and rs.insert(2)
    assert rs.getRandom() in (1, 2)
    assert rs.remove(1) and not rs.insert(2) and rs.getRandom() == 2
    print("OK")
