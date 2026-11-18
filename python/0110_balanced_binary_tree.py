# 110. Balanced Binary Tree (Easy)
# https://leetcode.com/problems/balanced-binary-tree/
# Run: python 0110_balanced_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def h(node):
            if not node:
                return 0
            l, r = h(node.left), h(node.right)
            if l < 0 or r < 0 or abs(l - r) > 1:
                return -1
            return 1 + max(l, r)

        return h(root) >= 0


if __name__ == "__main__":
    assert Solution().isBalanced(make_tree([3, 9, 20, None, None, 15, 7])) is True
    assert Solution().isBalanced(make_tree([1, 2, 2, 3, 3, None, None, 4, 4])) is False
    print("OK")
