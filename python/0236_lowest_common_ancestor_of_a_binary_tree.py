# 236. Lowest Common Ancestor of a Binary Tree (Medium)
# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
# Run: python 0236_lowest_common_ancestor_of_a_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        if not root or root is p or root is q:
            return root
        l = self.lowestCommonAncestor(root.left, p, q)
        r = self.lowestCommonAncestor(root.right, p, q)
        return root if l and r else l or r


if __name__ == "__main__":
    t = make_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
    assert Solution().lowestCommonAncestor(t, find_node(t, 5), find_node(t, 1)).val == 3
    assert Solution().lowestCommonAncestor(t, find_node(t, 5), find_node(t, 4)).val == 5
    print("OK")
