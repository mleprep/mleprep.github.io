# 1838. Frequency of the Most Frequent Element (Medium)
# https://leetcode.com/problems/frequency-of-the-most-frequent-element/
# Run: python 1838_frequency_of_the_most_frequent_element.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        nums.sort()
        l = s = best = 0
        for r, x in enumerate(nums):
            s += x
            while x * (r - l + 1) - s > k:
                s -= nums[l]
                l += 1
            best = max(best, r - l + 1)
        return best


if __name__ == "__main__":
    assert Solution().maxFrequency([1, 2, 4], 5) == 3
    assert Solution().maxFrequency([1, 4, 8, 13], 5) == 2
    print("OK")
