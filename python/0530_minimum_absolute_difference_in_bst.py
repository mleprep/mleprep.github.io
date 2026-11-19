# 530. Minimum Absolute Difference in BST (Easy)
# https://leetcode.com/problems/minimum-absolute-difference-in-bst/
# Run: python 0530_minimum_absolute_difference_in_bst.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
        best, prev = float("inf"), None

        def inorder(node):
            nonlocal best, prev
            if not node:
                return
            inorder(node.left)
            if prev is not None:
                best = min(best, node.val - prev)
            prev = node.val
            inorder(node.right)

        inorder(root)
        return best


if __name__ == "__main__":
    assert Solution().getMinimumDifference(make_tree([4, 2, 6, 1, 3])) == 1
    assert Solution().getMinimumDifference(make_tree([1, 0, 48, None, None, 12, 49])) == 1
    print("OK")
