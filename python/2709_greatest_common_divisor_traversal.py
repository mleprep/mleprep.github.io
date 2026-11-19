# 2709. Greatest Common Divisor Traversal (Hard)
# https://leetcode.com/problems/greatest-common-divisor-traversal/
# Run: python 2709_greatest_common_divisor_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canTraverseAllPairs(self, nums: List[int]) -> bool:
        n = len(nums)
        if n == 1:
            return True
        if 1 in nums:
            return False
        M = max(nums)
        spf = list(range(M + 1))                      # smallest prime factor
        for i in range(2, int(M ** 0.5) + 1):
            if spf[i] == i:
                for j in range(i * i, M + 1, i):
                    if spf[j] == j:
                        spf[j] = i
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        first = {}                                    # prime -> first index having it
        for i, x in enumerate(nums):
            while x > 1:
                p = spf[x]
                if p in first:
                    parent[find(i)] = find(first[p])
                else:
                    first[p] = i
                while x % p == 0:
                    x //= p
        return len({find(i) for i in range(n)}) == 1


if __name__ == "__main__":
    assert Solution().canTraverseAllPairs([2, 3, 6]) is True
    assert Solution().canTraverseAllPairs([3, 9, 5]) is False
    assert Solution().canTraverseAllPairs([4, 3, 12, 8]) is True
    assert Solution().canTraverseAllPairs([1]) is True
    assert Solution().canTraverseAllPairs([1, 1]) is False
    assert Solution().canTraverseAllPairs([2, 2]) is True
    assert Solution().canTraverseAllPairs([35, 77, 22, 6]) is True
    print("OK")
