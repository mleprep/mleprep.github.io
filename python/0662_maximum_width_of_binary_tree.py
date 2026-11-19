# 662. Maximum Width of Binary Tree (Medium)
# https://leetcode.com/problems/maximum-width-of-binary-tree/
# Run: python 0662_maximum_width_of_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def widthOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        best, q = 0, [(root, 0)]
        while q:
            base = q[0][1]
            best = max(best, q[-1][1] - base + 1)
            nxt = []
            for node, i in q:
                i -= base
                if node.left:
                    nxt.append((node.left, 2 * i))
                if node.right:
                    nxt.append((node.right, 2 * i + 1))
            q = nxt
        return best


if __name__ == "__main__":
    assert Solution().widthOfBinaryTree(make_tree([1, 3, 2, 5, 3, None, 9])) == 4
    assert Solution().widthOfBinaryTree(make_tree([1, 3, 2, 5, None, None, 9, 6, None, 7])) == 7
    print("OK")
