# 701. Insert into a Binary Search Tree (Medium)
# https://leetcode.com/problems/insert-into-a-binary-search-tree/
# Run: python 0701_insert_into_a_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        node = TreeNode(val)
        if not root:
            return node
        cur = root
        while True:
            if val < cur.val:
                if not cur.left:
                    cur.left = node
                    return root
                cur = cur.left
            else:
                if not cur.right:
                    cur.right = node
                    return root
                cur = cur.right


if __name__ == "__main__":
    def inorder(n):
        return inorder(n.left) + [n.val] + inorder(n.right) if n else []
    for vals, v in (([4, 2, 7, 1, 3], 5), ([40, 20, 60, 10, 30, 50, 70], 25), ([], 5), ([4, 2, 7, 1, 3, None, None, None, None, None, None], 0)):
        got = inorder(Solution().insertIntoBST(make_tree(vals), v))
        assert got == sorted([x for x in vals if x is not None] + [v])     # still a BST holding every value
    print("OK")
