# 1929. Concatenation of Array (Easy)
# https://leetcode.com/problems/concatenation-of-array/
# Run: python 1929_concatenation_of_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        return nums + nums


if __name__ == "__main__":
    assert Solution().getConcatenation([1, 2, 1]) == [1, 2, 1, 1, 2, 1]
    assert Solution().getConcatenation([1, 3, 2, 1]) == [1, 3, 2, 1, 1, 3, 2, 1]
    assert Solution().getConcatenation([5]) == [5, 5]
    print("OK")
