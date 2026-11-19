# 106. Construct Binary Tree from Inorder and Postorder Traversal (Medium)
# https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
# Run: python 0106_construct_binary_tree_from_inorder_and_postorder_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        pos = {v: i for i, v in enumerate(inorder)}

        def build(lo, hi):
            if lo >= hi:
                return None
            root = TreeNode(postorder.pop())
            m = pos[root.val]
            root.right = build(m + 1, hi)
            root.left = build(lo, m)
            return root

        return build(0, len(inorder))


if __name__ == "__main__":
    assert tree_vals(Solution().buildTree([9, 3, 15, 20, 7], [9, 15, 7, 20, 3])) == [3, 9, 20, None, None, 15, 7]
    print("OK")
