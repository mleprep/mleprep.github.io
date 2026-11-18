# 112. Path Sum (Easy)
# https://leetcode.com/problems/path-sum/
# Run: python 0112_path_sum.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False
        if not root.left and not root.right:
            return root.val == targetSum
        rest = targetSum - root.val
        return self.hasPathSum(root.left, rest) or self.hasPathSum(root.right, rest)


if __name__ == "__main__":
    t = make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, None, 1])
    assert Solution().hasPathSum(t, 22) is True
    assert Solution().hasPathSum(make_tree([1, 2, 3]), 5) is False
    print("OK")
