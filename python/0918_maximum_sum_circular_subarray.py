# 918. Maximum Sum Circular Subarray (Medium)
# https://leetcode.com/problems/maximum-sum-circular-subarray/
# Run: python 0918_maximum_sum_circular_subarray.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        cur_max = cur_min = 0
        best_max, best_min = nums[0], nums[0]
        for x in nums:
            cur_max = max(x, cur_max + x)
            best_max = max(best_max, cur_max)
            cur_min = min(x, cur_min + x)
            best_min = min(best_min, cur_min)
        if best_max < 0:
            return best_max
        return max(best_max, sum(nums) - best_min)


if __name__ == "__main__":
    assert Solution().maxSubarraySumCircular([1, -2, 3, -2]) == 3
    assert Solution().maxSubarraySumCircular([5, -3, 5]) == 10
    assert Solution().maxSubarraySumCircular([-3, -2, -3]) == -2
    print("OK")
