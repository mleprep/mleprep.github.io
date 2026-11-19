# 285. Inorder Successor in BST (Medium)
# https://leetcode.com/problems/inorder-successor-in-bst/
# Run: python 0285_inorder_successor_in_bst.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def inorderSuccessor(self, root: TreeNode, p: TreeNode) -> Optional[TreeNode]:
        succ = None
        while root:
            if p.val < root.val:
                succ, root = root, root.left
            else:
                root = root.right
        return succ


if __name__ == "__main__":
    t = make_tree([2, 1, 3])
    assert Solution().inorderSuccessor(t, find_node(t, 1)).val == 2
    t = make_tree([5, 3, 6, 2, 4, None, None, 1])
    assert Solution().inorderSuccessor(t, find_node(t, 6)) is None
    assert Solution().inorderSuccessor(t, find_node(t, 4)).val == 5
    assert Solution().inorderSuccessor(t, find_node(t, 1)).val == 2
    assert Solution().inorderSuccessor(t, find_node(t, 3)).val == 4
    print("OK")
