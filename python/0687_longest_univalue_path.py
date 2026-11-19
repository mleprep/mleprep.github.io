# 687. Longest Univalue Path (Medium)
# https://leetcode.com/problems/longest-univalue-path/
# Run: python 0687_longest_univalue_path.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestUnivaluePath(self, root: Optional[TreeNode]) -> int:
        best = 0

        def arm(node):
            nonlocal best
            if not node:
                return 0
            l, r = arm(node.left), arm(node.right)
            l = l + 1 if node.left and node.left.val == node.val else 0
            r = r + 1 if node.right and node.right.val == node.val else 0
            best = max(best, l + r)
            return max(l, r)

        arm(root)
        return best


if __name__ == "__main__":
    assert Solution().longestUnivaluePath(make_tree([5, 4, 5, 1, 1, None, 5])) == 2
    assert Solution().longestUnivaluePath(make_tree([1, 4, 5, 4, 4, None, 5])) == 2
    print("OK")
