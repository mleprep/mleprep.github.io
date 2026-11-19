# 659. Split Array into Consecutive Subsequences (Medium)
# https://leetcode.com/problems/split-array-into-consecutive-subsequences/
# Run: python 0659_split_array_into_consecutive_subsequences.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        left, ends = Counter(nums), Counter()
        for x in nums:
            if not left[x]:
                continue
            left[x] -= 1
            if ends[x - 1]:
                ends[x - 1] -= 1
                ends[x] += 1
            elif left[x + 1] and left[x + 2]:
                left[x + 1] -= 1
                left[x + 2] -= 1
                ends[x + 2] += 1
            else:
                return False
        return True


if __name__ == "__main__":
    assert Solution().isPossible([1, 2, 3, 3, 4, 5]) is True
    assert Solution().isPossible([1, 2, 3, 3, 4, 4, 5, 5]) is True
    assert Solution().isPossible([1, 2, 3, 4, 4, 5]) is False
    print("OK")
