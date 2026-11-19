# 229. Majority Element II (Medium)
# https://leetcode.com/problems/majority-element-ii/
# Run: python 0229_majority_element_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        c1, c2, n1, n2 = None, None, 0, 0
        for x in nums:
            if x == c1:
                n1 += 1
            elif x == c2:
                n2 += 1
            elif n1 == 0:
                c1, n1 = x, 1
            elif n2 == 0:
                c2, n2 = x, 1
            else:
                n1, n2 = n1 - 1, n2 - 1
        return [c for c in (c1, c2) if c is not None and nums.count(c) > len(nums) // 3]


if __name__ == "__main__":
    assert Solution().majorityElement([3, 2, 3]) == [3]
    assert Solution().majorityElement([1]) == [1]
    assert sorted(Solution().majorityElement([1, 2])) == [1, 2]
    assert sorted(Solution().majorityElement([2, 2, 1, 1, 1, 2, 2])) == [1, 2]
    assert Solution().majorityElement([1, 2, 3, 4]) == []
    print("OK")
