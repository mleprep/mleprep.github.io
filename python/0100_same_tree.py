# 100. Same Tree (Easy)
# https://leetcode.com/problems/same-tree/
# Run: python 0100_same_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p or not q:
            return p is q
        return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)


if __name__ == "__main__":
    assert Solution().isSameTree(make_tree([1, 2, 3]), make_tree([1, 2, 3])) is True
    assert Solution().isSameTree(make_tree([1, 2]), make_tree([1, None, 2])) is False
    print("OK")
