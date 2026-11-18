# 572. Subtree of Another Tree (Easy)
# https://leetcode.com/problems/subtree-of-another-tree/
# Run: python 0572_subtree_of_another_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def same(a, b):
            if not a or not b:
                return a is b
            return a.val == b.val and same(a.left, b.left) and same(a.right, b.right)

        if not root:
            return False
        return same(root, subRoot) or self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)


if __name__ == "__main__":
    assert Solution().isSubtree(make_tree([3, 4, 5, 1, 2]), make_tree([4, 1, 2])) is True
    assert Solution().isSubtree(make_tree([3, 4, 5, 1, 2, None, None, None, None, 0]), make_tree([4, 1, 2])) is False
    print("OK")
