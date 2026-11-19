# 450. Delete Node in a BST (Medium)
# https://leetcode.com/problems/delete-node-in-a-bst/
# Run: python 0450_delete_node_in_a_bst.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if not root:
            return None
        if key < root.val:
            root.left = self.deleteNode(root.left, key)
        elif key > root.val:
            root.right = self.deleteNode(root.right, key)
        else:
            if not root.left or not root.right:
                return root.left or root.right
            succ = root.right
            while succ.left:
                succ = succ.left
            root.val = succ.val
            root.right = self.deleteNode(root.right, succ.val)
        return root


if __name__ == "__main__":
    assert tree_vals(Solution().deleteNode(make_tree([5, 3, 6, 2, 4, None, 7]), 3)) == [5, 4, 6, 2, None, None, 7]
    assert tree_vals(Solution().deleteNode(make_tree([5, 3, 6, 2, 4, None, 7]), 0)) == [5, 3, 6, 2, 4, None, 7]
    print("OK")
