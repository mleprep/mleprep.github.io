# 1679. Max Number of K-Sum Pairs (Medium)
# https://leetcode.com/problems/max-number-of-k-sum-pairs/
# Run: python 1679_max_number_of_k_sum_pairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        waiting, ops = Counter(), 0
        for x in nums:
            if waiting[k - x]:
                waiting[k - x] -= 1
                ops += 1
            else:
                waiting[x] += 1
        return ops


if __name__ == "__main__":
    assert Solution().maxOperations([1, 2, 3, 4], 5) == 2
    assert Solution().maxOperations([3, 1, 3, 4, 3], 6) == 1
    assert Solution().maxOperations([2, 2, 2, 2], 4) == 2
    assert Solution().maxOperations([1], 2) == 0
    print("OK")
