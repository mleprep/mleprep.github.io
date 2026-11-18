# 26. Remove Duplicates from Sorted Array (Easy)
# https://leetcode.com/problems/remove-duplicates-from-sorted-array/
# Run: python 0026_remove_duplicates_from_sorted_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k == 0 or x != nums[k - 1]:
                nums[k] = x
                k += 1
        return k


if __name__ == "__main__":
    a = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
    k = Solution().removeDuplicates(a)
    assert k == 5 and a[:k] == [0, 1, 2, 3, 4]
    print("OK")
