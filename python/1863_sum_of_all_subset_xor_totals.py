# 1863. Sum of All Subset XOR Totals (Easy)
# https://leetcode.com/problems/sum-of-all-subset-xor-totals/
# Run: python 1863_sum_of_all_subset_xor_totals.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subsetXORSum(self, nums: List[int]) -> int:
        return functools.reduce(lambda x, y: x | y, nums) << (len(nums) - 1)


if __name__ == "__main__":
    assert Solution().subsetXORSum([1, 3]) == 6
    assert Solution().subsetXORSum([5, 1, 6]) == 28
    assert Solution().subsetXORSum([3, 4, 5, 6, 7, 8]) == 480
    nums = [9, 2, 7, 12, 1]
    brute = sum(functools.reduce(lambda x, y: x ^ y, c, 0) for r in range(len(nums) + 1) for c in itertools.combinations(nums, r))
    assert Solution().subsetXORSum(nums) == brute
    print("OK")
