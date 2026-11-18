# 108. Convert Sorted Array to Binary Search Tree (Easy)
# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
# Run: python 0108_convert_sorted_array_to_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def build(lo, hi):
            if lo > hi:
                return None
            mid = (lo + hi) // 2
            return TreeNode(nums[mid], build(lo, mid - 1), build(mid + 1, hi))

        return build(0, len(nums) - 1)


if __name__ == "__main__":
    assert tree_vals(Solution().sortedArrayToBST([-10, -3, 0, 5, 9])) == [0, -10, 5, None, -3, None, 9]
    print("OK")
