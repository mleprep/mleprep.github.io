# 368. Largest Divisible Subset (Medium)
# https://leetcode.com/problems/largest-divisible-subset/
# Run: python 0368_largest_divisible_subset.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        nums.sort()
        n = len(nums)
        size, parent = [1] * n, [-1] * n
        for i in range(n):
            for j in range(i):
                if nums[i] % nums[j] == 0 and size[j] + 1 > size[i]:
                    size[i], parent[i] = size[j] + 1, j
        i = max(range(n), key=size.__getitem__)
        res = []
        while i != -1:
            res.append(nums[i])
            i = parent[i]
        return res[::-1]


if __name__ == "__main__":
    assert Solution().largestDivisibleSubset([1, 2, 3]) in ([1, 2], [1, 3])
    assert Solution().largestDivisibleSubset([1, 2, 4, 8]) == [1, 2, 4, 8]
    print("OK")
