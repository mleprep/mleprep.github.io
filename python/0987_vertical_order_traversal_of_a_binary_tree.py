# 987. Vertical Order Traversal of a Binary Tree (Hard)
# https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
# Run: python 0987_vertical_order_traversal_of_a_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        cells = []

        def dfs(node, r, c):
            if node:
                cells.append((c, r, node.val))
                dfs(node.left, r + 1, c - 1)
                dfs(node.right, r + 1, c + 1)

        dfs(root, 0, 0)
        cells.sort()
        res, last = [], None
        for c, _, v in cells:
            if c != last:
                res.append([])
                last = c
            res[-1].append(v)
        return res


if __name__ == "__main__":
    assert Solution().verticalTraversal(make_tree([3, 9, 20, None, None, 15, 7])) == [[9], [3, 15], [20], [7]]
    assert Solution().verticalTraversal(make_tree([1, 2, 3, 4, 6, 5, 7])) == [[4], [2], [1, 5, 6], [3], [7]]
    print("OK")
