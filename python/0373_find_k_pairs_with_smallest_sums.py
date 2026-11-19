# 373. Find K Pairs with Smallest Sums (Medium)
# https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
# Run: python 0373_find_k_pairs_with_smallest_sums.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        h = [(nums1[i] + nums2[0], i, 0) for i in range(min(k, len(nums1)))]
        heapify(h)
        res = []
        while h and len(res) < k:
            _, i, j = heappop(h)
            res.append([nums1[i], nums2[j]])
            if j + 1 < len(nums2):
                heappush(h, (nums1[i] + nums2[j + 1], i, j + 1))
        return res


if __name__ == "__main__":
    assert sorted(Solution().kSmallestPairs([1, 7, 11], [2, 4, 6], 3)) == sorted([[1, 2], [1, 4], [1, 6]])
    assert sorted(Solution().kSmallestPairs([1, 1, 2], [1, 2, 3], 2)) == sorted([[1, 1], [1, 1]])
    print("OK")
