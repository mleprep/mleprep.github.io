# 978. Longest Turbulent Subarray (Medium)
# https://leetcode.com/problems/longest-turbulent-subarray/
# Run: python 0978_longest_turbulent_subarray.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        best = up = down = 1
        for i in range(1, len(arr)):
            if arr[i] > arr[i - 1]:
                up, down = down + 1, 1
            elif arr[i] < arr[i - 1]:
                up, down = 1, up + 1
            else:
                up = down = 1
            best = max(best, up, down)
        return best


if __name__ == "__main__":
    assert Solution().maxTurbulenceSize([9, 4, 2, 10, 7, 8, 8, 1, 9]) == 5
    assert Solution().maxTurbulenceSize([4, 8, 12, 16]) == 2
    assert Solution().maxTurbulenceSize([100]) == 1
    assert Solution().maxTurbulenceSize([9, 9]) == 1
    print("OK")
