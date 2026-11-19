# 2542. Maximum Subsequence Score (Medium)
# https://leetcode.com/problems/maximum-subsequence-score/
# Run: python 2542_maximum_subsequence_score.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        heap, total, best = [], 0, 0
        for b, a in sorted(zip(nums2, nums1), reverse=True):
            heappush(heap, a)
            total += a
            if len(heap) > k:
                total -= heappop(heap)
            if len(heap) == k:
                best = max(best, total * b)
        return best


if __name__ == "__main__":
    assert Solution().maxScore([1, 3, 3, 2], [2, 1, 3, 4], 3) == 12
    assert Solution().maxScore([4, 2, 3, 1, 1], [7, 5, 10, 9, 6], 1) == 30
    assert Solution().maxScore([2, 1, 14, 12], [11, 7, 13, 6], 3) == 168
    print("OK")
