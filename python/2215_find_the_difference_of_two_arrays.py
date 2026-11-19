# 2215. Find the Difference of Two Arrays (Easy)
# https://leetcode.com/problems/find-the-difference-of-two-arrays/
# Run: python 2215_find_the_difference_of_two_arrays.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        a, b = set(nums1), set(nums2)
        return [list(a - b), list(b - a)]


if __name__ == "__main__":
    assert [sorted(x) for x in Solution().findDifference([1, 2, 3], [2, 4, 6])] == [[1, 3], [4, 6]]
    assert [sorted(x) for x in Solution().findDifference([1, 2, 3, 3], [1, 1, 2, 2])] == [[3], []]
    assert [sorted(x) for x in Solution().findDifference([-1], [-1])] == [[], []]
    print("OK")
