# 698. Partition to K Equal Sum Subsets (Medium)
# https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
# Run: python 0698_partition_to_k_equal_sum_subsets.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        total = sum(nums)
        if total % k:
            return False
        target = total // k
        nums.sort(reverse=True)
        if nums[0] > target:
            return False
        buckets = [0] * k

        def bt(i):
            if i == len(nums):
                return True
            for b in range(k):
                if buckets[b] + nums[i] <= target:
                    buckets[b] += nums[i]
                    if bt(i + 1):
                        return True
                    buckets[b] -= nums[i]
                if buckets[b] == 0:
                    break                     # symmetric empty buckets
            return False

        return bt(0)


if __name__ == "__main__":
    assert Solution().canPartitionKSubsets([4, 3, 2, 3, 5, 2, 1], 4) is True
    assert Solution().canPartitionKSubsets([1, 2, 3, 4], 3) is False
    print("OK")
