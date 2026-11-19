# 109. Convert Sorted List to Binary Search Tree (Medium)
# https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
# Run: python 0109_convert_sorted_list_to_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortedListToBST(self, head: Optional[ListNode]) -> Optional[TreeNode]:
        n, p = 0, head
        while p:
            n, p = n + 1, p.next
        cur = head

        def build(lo, hi):
            nonlocal cur
            if lo > hi:
                return None
            mid = (lo + hi) // 2
            left = build(lo, mid - 1)
            root = TreeNode(cur.val, left)
            cur = cur.next
            root.right = build(mid + 1, hi)
            return root

        return build(0, n - 1)


if __name__ == "__main__":
    # any height-balanced BST with the same in-order sequence is accepted
    def balanced_bst(root, vals):
        def walk(n):  # (inorder values, height) or None if unbalanced
            if not n:
                return [], 0
            l, r = walk(n.left), walk(n.right)
            if l is None or r is None or abs(l[1] - r[1]) > 1:
                return None
            return l[0] + [n.val] + r[0], 1 + max(l[1], r[1])
        w = walk(root)
        return w is not None and w[0] == vals
    for vals in ([-10, -3, 0, 5, 9], [1, 3], [0], [1, 2, 3, 4, 5, 6, 7, 8]):
        assert balanced_bst(Solution().sortedListToBST(make_list(vals)), vals)
    assert Solution().sortedListToBST(None) is None
    print("OK")
