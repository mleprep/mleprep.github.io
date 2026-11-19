# 90. Subsets II (Medium)
# https://leetcode.com/problems/subsets-ii/
# Run: python 0090_subsets_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res, path = [], []

        def bt(start):
            res.append(path[:])
            for i in range(start, len(nums)):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                path.append(nums[i])
                bt(i + 1)
                path.pop()

        bt(0)
        return res


if __name__ == "__main__":
    assert sorted(map(sorted, Solution().subsetsWithDup([1, 2, 2]))) == sorted(map(sorted, [[], [1], [1, 2], [1, 2, 2], [2], [2, 2]]))
    assert sorted(map(sorted, Solution().subsetsWithDup([0]))) == [[], [0]]
    assert sorted(map(sorted, Solution().subsetsWithDup([4, 4, 1, 4]))) == [[], [1], [1, 4], [1, 4, 4], [1, 4, 4, 4], [4], [4, 4], [4, 4, 4]]
    print("OK")
