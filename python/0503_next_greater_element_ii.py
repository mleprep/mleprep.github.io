# 503. Next Greater Element II (Medium)
# https://leetcode.com/problems/next-greater-element-ii/
# Run: python 0503_next_greater_element_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res, st = [-1] * n, []
        for i in range(2 * n):
            x = nums[i % n]
            while st and nums[st[-1]] < x:
                res[st.pop()] = x
            if i < n:
                st.append(i)
        return res


if __name__ == "__main__":
    assert Solution().nextGreaterElements([1, 2, 1]) == [2, -1, 2]
    assert Solution().nextGreaterElements([1, 2, 3, 4, 3]) == [2, 3, 4, -1, 4]
    print("OK")
