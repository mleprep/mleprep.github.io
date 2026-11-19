# 179. Largest Number (Medium)
# https://leetcode.com/problems/largest-number/
# Run: python 0179_largest_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        s = sorted(map(str, nums), key=functools.cmp_to_key(lambda a, b: -1 if a + b > b + a else 1))
        res = "".join(s)
        return "0" if res[0] == "0" else res


if __name__ == "__main__":
    assert Solution().largestNumber([10, 2]) == "210"
    assert Solution().largestNumber([3, 30, 34, 5, 9]) == "9534330"
    assert Solution().largestNumber([0, 0]) == "0"
    print("OK")
