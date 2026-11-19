# 958. Check Completeness of a Binary Tree (Medium)
# https://leetcode.com/problems/check-completeness-of-a-binary-tree/
# Run: python 0958_check_completeness_of_a_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        q, seen_null = deque([root]), False
        while q:
            node = q.popleft()
            if not node:
                seen_null = True
                continue
            if seen_null:
                return False
            q.append(node.left)
            q.append(node.right)
        return True


if __name__ == "__main__":
    assert Solution().isCompleteTree(make_tree([1, 2, 3, 4, 5, 6])) is True
    assert Solution().isCompleteTree(make_tree([1, 2, 3, 4, 5, None, 7])) is False
    print("OK")
