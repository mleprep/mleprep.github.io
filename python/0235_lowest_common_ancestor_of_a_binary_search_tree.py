# 235. Lowest Common Ancestor of a Binary Search Tree (Medium)
# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
# Run: python 0235_lowest_common_ancestor_of_a_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        while root:
            if p.val < root.val and q.val < root.val:
                root = root.left
            elif p.val > root.val and q.val > root.val:
                root = root.right
            else:
                return root


if __name__ == "__main__":
    t = make_tree([6, 2, 8, 0, 4, 7, 9, None, None, 3, 5])
    assert Solution().lowestCommonAncestor(t, find_node(t, 2), find_node(t, 8)).val == 6
    assert Solution().lowestCommonAncestor(t, find_node(t, 2), find_node(t, 4)).val == 2
    print("OK")
