# 992. Subarrays with K Different Integers (Hard)
# https://leetcode.com/problems/subarrays-with-k-different-integers/
# Run: python 0992_subarrays_with_k_different_integers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        def at_most(k):
            cnt = defaultdict(int)
            l = res = 0
            for r, x in enumerate(nums):
                cnt[x] += 1
                while len(cnt) > k:
                    cnt[nums[l]] -= 1
                    if cnt[nums[l]] == 0:
                        del cnt[nums[l]]
                    l += 1
                res += r - l + 1
            return res

        return at_most(k) - at_most(k - 1)


if __name__ == "__main__":
    assert Solution().subarraysWithKDistinct([1, 2, 1, 2, 3], 2) == 7
    assert Solution().subarraysWithKDistinct([1, 2, 1, 3, 4], 3) == 3
    print("OK")
