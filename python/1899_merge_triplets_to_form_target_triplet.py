# 1899. Merge Triplets to Form Target Triplet (Medium)
# https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
# Run: python 1899_merge_triplets_to_form_target_triplet.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mergeTriplets(self, triplets: List[List[int]], target: List[int]) -> bool:
        good = set()
        for t in triplets:
            if all(t[i] <= target[i] for i in range(3)):
                good |= {i for i in range(3) if t[i] == target[i]}
        return len(good) == 3


if __name__ == "__main__":
    assert Solution().mergeTriplets([[2, 5, 3], [1, 8, 4], [1, 7, 5]], [2, 7, 5]) is True
    assert Solution().mergeTriplets([[3, 4, 5], [4, 5, 6]], [3, 2, 5]) is False
    print("OK")
