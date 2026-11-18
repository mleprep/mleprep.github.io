# 101. Symmetric Tree (Easy)
# https://leetcode.com/problems/symmetric-tree/
# Run: python 0101_symmetric_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def mirror(a, b):
            if not a or not b:
                return a is b
            return a.val == b.val and mirror(a.left, b.right) and mirror(a.right, b.left)

        return mirror(root.left, root.right)


if __name__ == "__main__":
    assert Solution().isSymmetric(make_tree([1, 2, 2, 3, 4, 4, 3])) is True
    assert Solution().isSymmetric(make_tree([1, 2, 2, None, 3, None, 3])) is False
    print("OK")
