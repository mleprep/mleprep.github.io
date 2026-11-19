# 89. Gray Code (Medium)
# https://leetcode.com/problems/gray-code/
# Run: python 0089_gray_code.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def grayCode(self, n: int) -> List[int]:
        return [i ^ (i >> 1) for i in range(1 << n)]


if __name__ == "__main__":
    g = Solution().grayCode(3)
    assert g == [0, 1, 3, 2, 6, 7, 5, 4]
    assert all(bin(g[i] ^ g[(i + 1) % len(g)]).count("1") == 1 for i in range(len(g)))
    print("OK")
