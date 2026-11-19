# 860. Lemonade Change (Easy)
# https://leetcode.com/problems/lemonade-change/
# Run: python 0860_lemonade_change.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        five = ten = 0
        for b in bills:
            if b == 5:
                five += 1
            elif b == 10:
                five, ten = five - 1, ten + 1
            elif ten:
                five, ten = five - 1, ten - 1
            else:
                five -= 3
            if five < 0:
                return False
        return True


if __name__ == "__main__":
    assert Solution().lemonadeChange([5, 5, 5, 10, 20]) is True
    assert Solution().lemonadeChange([5, 5, 10, 10, 20]) is False
    print("OK")
