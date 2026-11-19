# 643. Maximum Average Subarray I (Easy)
# https://leetcode.com/problems/maximum-average-subarray-i/
# Run: python 0643_maximum_average_subarray_i.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        s = best = sum(nums[:k])
        for i in range(k, len(nums)):
            s += nums[i] - nums[i - k]
            best = max(best, s)
        return best / k


if __name__ == "__main__":
    assert Solution().findMaxAverage([1, 12, -5, -6, 50, 3], 4) == 12.75
    print("OK")
