# 202. Happy Number (Easy)
# https://leetcode.com/problems/happy-number/
# Run: python 0202_happy_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isHappy(self, n: int) -> bool:
        def nxt(x):
            return sum(int(d) ** 2 for d in str(x))

        slow, fast = n, nxt(n)
        while fast != 1 and slow != fast:
            slow, fast = nxt(slow), nxt(nxt(fast))
        return fast == 1


if __name__ == "__main__":
    assert Solution().isHappy(19) is True
    assert Solution().isHappy(2) is False
    print("OK")
