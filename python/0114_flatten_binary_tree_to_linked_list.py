# 114. Flatten Binary Tree to Linked List (Medium)
# https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
# Run: python 0114_flatten_binary_tree_to_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        cur = root
        while cur:
            if cur.left:
                p = cur.left
                while p.right:
                    p = p.right
                p.right = cur.right
                cur.right, cur.left = cur.left, None
            cur = cur.right


if __name__ == "__main__":
    t = make_tree([1, 2, 5, 3, 4, None, 6]); Solution().flatten(t)
    assert tree_vals(t) == [1, None, 2, None, 3, None, 4, None, 5, None, 6]
    print("OK")
