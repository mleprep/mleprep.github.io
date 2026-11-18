# 173. Binary Search Tree Iterator (Medium)
# https://leetcode.com/problems/binary-search-tree-iterator/
# Run: python 0173_binary_search_tree_iterator.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class BSTIterator:
    def __init__(self, root: Optional[TreeNode]):
        self.st = []
        self._push_left(root)

    def _push_left(self, node):
        while node:
            self.st.append(node)
            node = node.left

    def next(self) -> int:
        node = self.st.pop()
        self._push_left(node.right)
        return node.val

    def hasNext(self) -> bool:
        return bool(self.st)


if __name__ == "__main__":
    it = BSTIterator(make_tree([7, 3, 15, None, None, 9, 20]))
    out = []
    while it.hasNext():
        out.append(it.next())
    assert out == [3, 7, 9, 15, 20]
    print("OK")
