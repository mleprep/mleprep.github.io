# 974. Subarray Sums Divisible by K (Medium)
# https://leetcode.com/problems/subarray-sums-divisible-by-k/
# Run: python 0974_subarray_sums_divisible_by_k.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        cnt = [0] * k
        cnt[0] = 1
        pre = res = 0
        for x in nums:
            pre = (pre + x) % k           # Python % is already non-negative
            res += cnt[pre]
            cnt[pre] += 1
        return res


if __name__ == "__main__":
    assert Solution().subarraysDivByK([4, 5, 0, -2, -3, 1], 5) == 7
    assert Solution().subarraysDivByK([5], 9) == 0
    print("OK")
