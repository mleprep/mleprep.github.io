# 456. 132 Pattern (Medium)
# https://leetcode.com/problems/132-pattern/
# Run: python 0456_132_pattern.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def find132pattern(self, nums: List[int]) -> bool:
        st, two = [], float("-inf")
        for x in reversed(nums):
            if x < two:
                return True
            while st and st[-1] < x:
                two = st.pop()          # x acts as "3", popped value as "2"
            st.append(x)
        return False


if __name__ == "__main__":
    assert Solution().find132pattern([1, 2, 3, 4]) is False
    assert Solution().find132pattern([3, 1, 4, 2]) is True
    assert Solution().find132pattern([-1, 3, 2, 0]) is True
    print("OK")
