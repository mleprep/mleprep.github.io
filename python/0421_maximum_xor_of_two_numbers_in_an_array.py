# 421. Maximum XOR of Two Numbers in an Array (Medium)
# https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
# Run: python 0421_maximum_xor_of_two_numbers_in_an_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMaximumXOR(self, nums: List[int]) -> int:
        root, best = {}, 0
        for x in nums:
            node = root
            for b in range(30, -1, -1):
                node = node.setdefault((x >> b) & 1, {})
        for x in nums:
            node, cur = root, 0
            for b in range(30, -1, -1):
                bit = (x >> b) & 1
                if 1 - bit in node:
                    cur |= 1 << b
                    node = node[1 - bit]
                else:
                    node = node[bit]
            best = max(best, cur)
        return best


if __name__ == "__main__":
    assert Solution().findMaximumXOR([3, 10, 5, 25, 2, 8]) == 28
    assert Solution().findMaximumXOR([14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70]) == 127
    print("OK")
