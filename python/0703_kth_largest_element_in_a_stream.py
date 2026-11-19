# 703. Kth Largest Element in a Stream (Easy)
# https://leetcode.com/problems/kth-largest-element-in-a-stream/
# Run: python 0703_kth_largest_element_in_a_stream.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k, self.h = k, []
        for x in nums:
            self.add(x)

    def add(self, val: int) -> int:
        heappush(self.h, val)
        if len(self.h) > self.k:
            heappop(self.h)
        return self.h[0]


if __name__ == "__main__":
    kl = KthLargest(3, [4, 5, 8, 2])
    assert [kl.add(v) for v in (3, 5, 10, 9, 4)] == [4, 5, 5, 8, 8]
    kl = KthLargest(1, [])
    assert [kl.add(v) for v in (-3, -2, -4, 0, 4)] == [-3, -2, -2, 0, 4]
    print("OK")
