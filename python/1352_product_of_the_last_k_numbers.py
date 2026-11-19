# 1352. Product of the Last K Numbers (Medium)
# https://leetcode.com/problems/product-of-the-last-k-numbers/
# Run: python 1352_product_of_the_last_k_numbers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class ProductOfNumbers:
    def __init__(self):
        self.pre = [1]

    def add(self, num: int) -> None:
        if num == 0:
            self.pre = [1]
        else:
            self.pre.append(self.pre[-1] * num)

    def getProduct(self, k: int) -> int:
        if k >= len(self.pre):
            return 0
        return self.pre[-1] // self.pre[-1 - k]


if __name__ == "__main__":
    p = ProductOfNumbers()
    for x in (3, 0, 2, 5, 4):
        p.add(x)
    assert [p.getProduct(2), p.getProduct(3), p.getProduct(4)] == [20, 40, 0]
    p.add(8)
    assert p.getProduct(2) == 32
    print("OK")
