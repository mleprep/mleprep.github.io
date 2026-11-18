# 1448. Count Good Nodes in Binary Tree (Medium)
# https://leetcode.com/problems/count-good-nodes-in-binary-tree/
# Run: python 1448_count_good_nodes_in_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, mx):
            if not node:
                return 0
            good = node.val >= mx
            mx = max(mx, node.val)
            return good + dfs(node.left, mx) + dfs(node.right, mx)

        return dfs(root, root.val)


if __name__ == "__main__":
    assert Solution().goodNodes(make_tree([3, 1, 4, 3, None, 1, 5])) == 4
    assert Solution().goodNodes(make_tree([3, 3, None, 4, 2])) == 3
    print("OK")
