# 1095. Find in Mountain Array (Hard)
# https://leetcode.com/problems/find-in-mountain-array/
# Run: python 1095_find_in_mountain_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findInMountainArray(self, target: int, mountainArr: 'MountainArray') -> int:
        n = mountainArr.length()
        lo, hi = 0, n - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if mountainArr.get(mid) < mountainArr.get(mid + 1):
                lo = mid + 1
            else:
                hi = mid
        peak = lo

        def search(lo, hi, ascending):
            while lo <= hi:
                mid = (lo + hi) // 2
                v = mountainArr.get(mid)
                if v == target:
                    return mid
                if (v < target) == ascending:
                    lo = mid + 1
                else:
                    hi = mid - 1
            return -1

        i = search(0, peak, True)
        return i if i != -1 else search(peak + 1, n - 1, False)


if __name__ == "__main__":
    class MountainArray:  # LeetCode's interface (counts calls; the limit is 100)
        def __init__(self, a): self.a, self.calls = a, 0
        def get(self, i): self.calls += 1; return self.a[i]
        def length(self): return len(self.a)
    m = MountainArray([1, 2, 3, 4, 5, 3, 1])
    assert Solution().findInMountainArray(3, m) == 2 and m.calls <= 100
    assert Solution().findInMountainArray(3, MountainArray([0, 1, 2, 4, 2, 1])) == -1
    assert Solution().findInMountainArray(1, MountainArray([0, 1, 2, 4, 2, 1])) == 1
    big = MountainArray(list(range(5000)) + list(range(4998, -1, -1)))
    assert Solution().findInMountainArray(4000, big) == 4000 and big.calls <= 100
    big2 = MountainArray(list(range(5000)) + list(range(4998, 2000, -1)))
    assert Solution().findInMountainArray(2001, big2) == 2001
    assert Solution().findInMountainArray(2000, MountainArray(list(range(10, 3000)) + [5])) == 1990
    print("OK")
