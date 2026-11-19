# 872. Leaf-Similar Trees (Easy)
# https://leetcode.com/problems/leaf-similar-trees/
# Run: python 0872_leaf_similar_trees.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        def leaves(node):
            if node:
                if not node.left and not node.right:
                    yield node.val
                yield from leaves(node.left)
                yield from leaves(node.right)

        return list(leaves(root1)) == list(leaves(root2))


if __name__ == "__main__":
    a = make_tree([3, 5, 1, 6, 2, 9, 8, None, None, 7, 4])
    b = make_tree([3, 5, 1, 6, 7, 4, 2, None, None, None, None, None, None, 9, 8])
    assert Solution().leafSimilar(a, b) is True
    assert Solution().leafSimilar(make_tree([1, 2, 3]), make_tree([1, 3, 2])) is False
    assert Solution().leafSimilar(make_tree([1]), make_tree([2, 1])) is True
    print("OK")
