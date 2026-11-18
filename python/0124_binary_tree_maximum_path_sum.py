# 124. Binary Tree Maximum Path Sum (Hard)
# https://leetcode.com/problems/binary-tree-maximum-path-sum/
# Run: python 0124_binary_tree_maximum_path_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        best = float("-inf")

        def gain(node):
            nonlocal best
            if not node:
                return 0
            l, r = max(gain(node.left), 0), max(gain(node.right), 0)
            best = max(best, node.val + l + r)
            return node.val + max(l, r)

        gain(root)
        return best


if __name__ == "__main__":
    assert Solution().maxPathSum(make_tree([1, 2, 3])) == 6
    assert Solution().maxPathSum(make_tree([-10, 9, 20, None, None, 15, 7])) == 42
    assert Solution().maxPathSum(make_tree([-3])) == -3
    print("OK")
