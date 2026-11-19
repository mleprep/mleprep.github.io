# 977. Squares of a Sorted Array (Easy)
# https://leetcode.com/problems/squares-of-a-sorted-array/
# Run: python 0977_squares_of_a_sorted_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res = [0] * n
        l, r = 0, n - 1
        for k in range(n - 1, -1, -1):
            if abs(nums[l]) > abs(nums[r]):
                res[k] = nums[l] ** 2; l += 1
            else:
                res[k] = nums[r] ** 2; r -= 1
        return res


if __name__ == "__main__":
    assert Solution().sortedSquares([-4, -1, 0, 3, 10]) == [0, 1, 9, 16, 100]
    print("OK")
