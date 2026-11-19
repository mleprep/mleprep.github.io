# 105. Construct Binary Tree from Preorder and Inorder Traversal (Medium)
# https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
# Run: python 0105_construct_binary_tree_from_preorder_and_inorder_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        pos = {v: i for i, v in enumerate(inorder)}
        it = iter(preorder)

        def build(lo, hi):              # inorder range [lo, hi)
            if lo >= hi:
                return None
            root = TreeNode(next(it))
            m = pos[root.val]
            root.left = build(lo, m)    # consume preorder left subtree first
            root.right = build(m + 1, hi)
            return root

        return build(0, len(inorder))


if __name__ == "__main__":
    assert tree_vals(Solution().buildTree([3, 9, 20, 15, 7], [9, 3, 15, 20, 7])) == [3, 9, 20, None, None, 15, 7]
    assert tree_vals(Solution().buildTree([-1], [-1])) == [-1]
    assert tree_vals(Solution().buildTree([1, 2, 3], [3, 2, 1])) == [1, 2, None, 3]
    print("OK")
