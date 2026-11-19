# 27. Remove Element (Easy)
# https://leetcode.com/problems/remove-element/
# Run: python 0027_remove_element.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        k = 0
        for x in nums:
            if x != val:
                nums[k] = x
                k += 1
        return k


if __name__ == "__main__":
    a = [0, 1, 2, 2, 3, 0, 4, 2]
    k = Solution().removeElement(a, 2)
    assert k == 5 and sorted(a[:k]) == [0, 0, 1, 3, 4]
    print("OK")
