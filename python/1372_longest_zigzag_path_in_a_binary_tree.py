# 1372. Longest ZigZag Path in a Binary Tree (Medium)
# https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
# Run: python 1372_longest_zigzag_path_in_a_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        best = 0

        def dfs(node):                       # -> (zigzag length going left, going right), in edges
            nonlocal best
            if not node:
                return -1, -1
            left = dfs(node.left)[1] + 1
            right = dfs(node.right)[0] + 1
            best = max(best, left, right)
            return left, right

        dfs(root)
        return best


if __name__ == "__main__":
    assert Solution().longestZigZag(make_tree([1, None, 1, 1, 1, None, None, 1, 1, None, 1, None, None, None, 1])) == 3
    assert Solution().longestZigZag(make_tree([1, 1, 1, None, 1, None, None, 1, 1, None, 1])) == 4
    assert Solution().longestZigZag(make_tree([1])) == 0
    print("OK")
