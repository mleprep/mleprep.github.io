# 15. 3Sum (Medium)
# https://leetcode.com/problems/3sum/
# Run: python 0015_3sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        for i in range(len(nums) - 2):
            if nums[i] > 0:
                break
            if i and nums[i] == nums[i - 1]:
                continue
            l, r = i + 1, len(nums) - 1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if s < 0:
                    l += 1
                elif s > 0:
                    r -= 1
                else:
                    res.append([nums[i], nums[l], nums[r]])
                    l += 1
                    while l < r and nums[l] == nums[l - 1]:
                        l += 1
                    r -= 1
        return res


if __name__ == "__main__":
    assert sorted(map(sorted, Solution().threeSum([-1, 0, 1, 2, -1, -4]))) == sorted(map(sorted, [[-1, -1, 2], [-1, 0, 1]]))
    assert sorted(map(sorted, Solution().threeSum([0, 0, 0]))) == sorted(map(sorted, [[0, 0, 0]]))
    print("OK")
