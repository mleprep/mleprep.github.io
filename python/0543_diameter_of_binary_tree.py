# 543. Diameter of Binary Tree (Easy)
# https://leetcode.com/problems/diameter-of-binary-tree/
# Run: python 0543_diameter_of_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        best = 0

        def height(node):
            nonlocal best
            if not node:
                return 0
            l, r = height(node.left), height(node.right)
            best = max(best, l + r)
            return 1 + max(l, r)

        height(root)
        return best


if __name__ == "__main__":
    assert Solution().diameterOfBinaryTree(make_tree([1, 2, 3, 4, 5])) == 3
    assert Solution().diameterOfBinaryTree(make_tree([1, 2])) == 1
    print("OK")
