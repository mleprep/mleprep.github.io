# 1431. Kids With the Greatest Number of Candies (Easy)
# https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
# Run: python 1431_kids_with_the_greatest_number_of_candies.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        best = max(candies)
        return [c + extraCandies >= best for c in candies]


if __name__ == "__main__":
    assert Solution().kidsWithCandies([2, 3, 5, 1, 3], 3) == [True, True, True, False, True]
    assert Solution().kidsWithCandies([4, 2, 1, 1, 2], 1) == [True, False, False, False, False]
    assert Solution().kidsWithCandies([12, 1, 12], 10) == [True, False, True]
    print("OK")
