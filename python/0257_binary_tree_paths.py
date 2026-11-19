# 257. Binary Tree Paths (Easy)
# https://leetcode.com/problems/binary-tree-paths/
# Run: python 0257_binary_tree_paths.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        res = []

        def dfs(node, path):
            if not node:
                return
            path = path + [str(node.val)]
            if not node.left and not node.right:
                res.append("->".join(path))
            dfs(node.left, path)
            dfs(node.right, path)

        dfs(root, [])
        return res


if __name__ == "__main__":
    assert Solution().binaryTreePaths(make_tree([1, 2, 3, None, 5])) == ["1->2->5", "1->3"]
    print("OK")
