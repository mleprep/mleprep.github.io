# 129. Sum Root to Leaf Numbers (Medium)
# https://leetcode.com/problems/sum-root-to-leaf-numbers/
# Run: python 0129_sum_root_to_leaf_numbers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        def dfs(node, cur):
            if not node:
                return 0
            cur = cur * 10 + node.val
            if not node.left and not node.right:
                return cur
            return dfs(node.left, cur) + dfs(node.right, cur)

        return dfs(root, 0)


if __name__ == "__main__":
    assert Solution().sumNumbers(make_tree([1, 2, 3])) == 25
    assert Solution().sumNumbers(make_tree([4, 9, 0, 5, 1])) == 1026
    print("OK")
