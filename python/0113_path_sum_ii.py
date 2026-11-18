# 113. Path Sum II (Medium)
# https://leetcode.com/problems/path-sum-ii/
# Run: python 0113_path_sum_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        res, path = [], []

        def dfs(node, rest):
            if not node:
                return
            path.append(node.val)
            rest -= node.val
            if not node.left and not node.right and rest == 0:
                res.append(path[:])
            dfs(node.left, rest)
            dfs(node.right, rest)
            path.pop()

        dfs(root, targetSum)
        return res


if __name__ == "__main__":
    t = make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1])
    assert Solution().pathSum(t, 22) == [[5, 4, 11, 2], [5, 8, 4, 5]]
    print("OK")
