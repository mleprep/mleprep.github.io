# 560. Subarray Sum Equals K (Medium)
# https://leetcode.com/problems/subarray-sum-equals-k/
# Run: python 0560_subarray_sum_equals_k.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count, pre = 0, 0
        seen = defaultdict(int, {0: 1})
        for x in nums:
            pre += x
            count += seen[pre - k]
            seen[pre] += 1
        return count


if __name__ == "__main__":
    assert Solution().subarraySum([1, 1, 1], 2) == 2
    assert Solution().subarraySum([1, 2, 3], 3) == 2
    assert Solution().subarraySum([1, -1, 0], 0) == 3
    print("OK")
