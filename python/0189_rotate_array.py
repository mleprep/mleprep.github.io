# 189. Rotate Array (Medium)
# https://leetcode.com/problems/rotate-array/
# Run: python 0189_rotate_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        def rev(i, j):
            while i < j:
                nums[i], nums[j] = nums[j], nums[i]
                i, j = i + 1, j - 1
        n = len(nums)
        k %= n
        rev(0, n - 1); rev(0, k - 1); rev(k, n - 1)


if __name__ == "__main__":
    a = [1, 2, 3, 4, 5, 6, 7]; Solution().rotate(a, 3)
    assert a == [5, 6, 7, 1, 2, 3, 4]
    b = [-1, -100, 3, 99]; Solution().rotate(b, 2)
    assert b == [3, 99, -1, -100]
    print("OK")
