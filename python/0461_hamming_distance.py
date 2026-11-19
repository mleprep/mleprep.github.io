# 461. Hamming Distance (Easy)
# https://leetcode.com/problems/hamming-distance/
# Run: python 0461_hamming_distance.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        return bin(x ^ y).count("1")


if __name__ == "__main__":
    assert Solution().hammingDistance(1, 4) == 2
    assert Solution().hammingDistance(3, 1) == 1
    print("OK")
