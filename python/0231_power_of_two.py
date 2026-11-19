# 231. Power of Two (Easy)
# https://leetcode.com/problems/power-of-two/
# Run: python 0231_power_of_two.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and n & (n - 1) == 0


if __name__ == "__main__":
    assert Solution().isPowerOfTwo(16) is True
    assert Solution().isPowerOfTwo(3) is False
    assert Solution().isPowerOfTwo(-16) is False
    print("OK")
