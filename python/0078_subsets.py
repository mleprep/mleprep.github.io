# 78. Subsets (Medium)
# https://leetcode.com/problems/subsets/
# Run: python 0078_subsets.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res, path = [], []

        def bt(start):
            res.append(path[:])
            for i in range(start, len(nums)):
                path.append(nums[i])
                bt(i + 1)
                path.pop()

        bt(0)
        return res


if __name__ == "__main__":
    assert sorted(map(sorted, Solution().subsets([1, 2, 3]))) == sorted(map(sorted, [[], [1], [1, 2], [1, 2, 3], [1, 3], [2], [2, 3], [3]]))
    assert sorted(map(sorted, Solution().subsets([0]))) == [[], [0]]
    print("OK")
