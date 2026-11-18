# 98. Validate Binary Search Tree (Medium)
# https://leetcode.com/problems/validate-binary-search-tree/
# Run: python 0098_validate_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def ok(node, lo, hi):
            if not node:
                return True
            if not lo < node.val < hi:
                return False
            return ok(node.left, lo, node.val) and ok(node.right, node.val, hi)

        return ok(root, float("-inf"), float("inf"))


if __name__ == "__main__":
    assert Solution().isValidBST(make_tree([2, 1, 3])) is True
    assert Solution().isValidBST(make_tree([5, 1, 6, None, None, 3, 7])) is False
    print("OK")
