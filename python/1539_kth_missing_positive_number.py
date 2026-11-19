# 1539. Kth Missing Positive Number (Easy)
# https://leetcode.com/problems/kth-missing-positive-number/
# Run: python 1539_kth_missing_positive_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        lo, hi = 0, len(arr)
        while lo < hi:
            mid = (lo + hi) // 2
            if arr[mid] - (mid + 1) >= k:
                hi = mid
            else:
                lo = mid + 1
        return lo + k


if __name__ == "__main__":
    assert Solution().findKthPositive([2, 3, 4, 7, 11], 5) == 9
    assert Solution().findKthPositive([1, 2, 3, 4], 2) == 6
    print("OK")
