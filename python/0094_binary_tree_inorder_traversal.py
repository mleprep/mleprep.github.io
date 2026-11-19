# 94. Binary Tree Inorder Traversal (Easy)
# https://leetcode.com/problems/binary-tree-inorder-traversal/
# Run: python 0094_binary_tree_inorder_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res, st, node = [], [], root
        while st or node:
            while node:
                st.append(node)
                node = node.left
            node = st.pop()
            res.append(node.val)
            node = node.right
        return res


if __name__ == "__main__":
    assert Solution().inorderTraversal(make_tree([1, None, 2, 3])) == [1, 3, 2]
    assert Solution().inorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [4, 2, 6, 5, 7, 1, 3, 9, 8]
    print("OK")
