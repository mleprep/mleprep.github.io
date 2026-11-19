# 1325. Delete Leaves With a Given Value (Medium)
# https://leetcode.com/problems/delete-leaves-with-a-given-value/
# Run: python 1325_delete_leaves_with_a_given_value.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        if not root:
            return None
        root.left = self.removeLeafNodes(root.left, target)
        root.right = self.removeLeafNodes(root.right, target)
        if not root.left and not root.right and root.val == target:
            return None
        return root


if __name__ == "__main__":
    assert tree_vals(Solution().removeLeafNodes(make_tree([1, 2, 3, 2, None, 2, 4]), 2)) == [1, None, 3, None, 4]
    assert tree_vals(Solution().removeLeafNodes(make_tree([1, 3, 3, 3, 2]), 3)) == [1, 3, None, None, 2]
    assert tree_vals(Solution().removeLeafNodes(make_tree([1, 2, None, 2, None, 2]), 2)) == [1]
    assert Solution().removeLeafNodes(make_tree([1, 1, 1]), 1) is None
    print("OK")
