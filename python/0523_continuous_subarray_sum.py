# 523. Continuous Subarray Sum (Medium)
# https://leetcode.com/problems/continuous-subarray-sum/
# Run: python 0523_continuous_subarray_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        first = {0: -1}
        pre = 0
        for i, x in enumerate(nums):
            pre = (pre + x) % k
            if pre in first:
                if i - first[pre] >= 2:
                    return True
            else:
                first[pre] = i
        return False


if __name__ == "__main__":
    assert Solution().checkSubarraySum([23, 2, 4, 6, 7], 6) is True
    assert Solution().checkSubarraySum([23, 2, 6, 4, 7], 13) is False
    assert Solution().checkSubarraySum([5, 0, 0, 0], 3) is True
    print("OK")
