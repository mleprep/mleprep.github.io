# 18. 4Sum (Medium)
# https://leetcode.com/problems/4sum/
# Run: python 0018_4sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()

        def ksum(start, k, target):
            res = []
            if k == 2:
                l, r = start, len(nums) - 1
                while l < r:
                    s = nums[l] + nums[r]
                    if s < target or (l > start and nums[l] == nums[l - 1]):
                        l += 1
                    elif s > target:
                        r -= 1
                    else:
                        res.append([nums[l], nums[r]])
                        l += 1; r -= 1
                return res
            for i in range(start, len(nums) - k + 1):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                for sub in ksum(i + 1, k - 1, target - nums[i]):
                    res.append([nums[i]] + sub)
            return res

        return ksum(0, 4, target)


if __name__ == "__main__":
    assert sorted(map(sorted, Solution().fourSum([1, 0, -1, 0, -2, 2], 0))) == sorted(map(sorted, [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]))
    assert sorted(map(sorted, Solution().fourSum([2, 2, 2, 2, 2], 8))) == sorted(map(sorted, [[2, 2, 2, 2]]))
    print("OK")
