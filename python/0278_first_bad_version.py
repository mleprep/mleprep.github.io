# 278. First Bad Version (Easy)
# https://leetcode.com/problems/first-bad-version/
# Run: python 0278_first_bad_version.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def firstBadVersion(self, n: int) -> int:
        lo, hi = 1, n
        while lo < hi:
            mid = (lo + hi) // 2
            if isBadVersion(mid):      # provided by LeetCode
                hi = mid
            else:
                lo = mid + 1
        return lo


if __name__ == "__main__":
    BAD = 4
    def isBadVersion(v): return v >= BAD
    assert Solution().firstBadVersion(5) == 4
    BAD = 1
    assert Solution().firstBadVersion(1) == 1
    print("OK")
