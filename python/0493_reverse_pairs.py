# 493. Reverse Pairs (Hard)
# https://leetcode.com/problems/reverse-pairs/
# Run: python 0493_reverse_pairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reversePairs(self, nums: List[int]) -> int:
        def sort(a):
            if len(a) <= 1:
                return a, 0
            mid = len(a) // 2
            left, cl = sort(a[:mid])
            right, cr = sort(a[mid:])
            count, j = cl + cr, 0
            for x in left:
                while j < len(right) and x > 2 * right[j]:
                    j += 1
                count += j
            merged, i, k = [], 0, 0
            while i < len(left) and k < len(right):
                if left[i] <= right[k]:
                    merged.append(left[i]); i += 1
                else:
                    merged.append(right[k]); k += 1
            return merged + left[i:] + right[k:], count

        return sort(nums)[1]


if __name__ == "__main__":
    assert Solution().reversePairs([1, 3, 2, 3, 1]) == 2
    assert Solution().reversePairs([2, 4, 3, 5, 1]) == 3
    assert Solution().reversePairs([2147483647, 2147483647, 2147483647]) == 0
    print("OK")
