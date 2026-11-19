# 111. Minimum Depth of Binary Tree (Easy)
# https://leetcode.com/problems/minimum-depth-of-binary-tree/
# Run: python 0111_minimum_depth_of_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        q, depth = deque([root]), 1
        while q:
            for _ in range(len(q)):
                n = q.popleft()
                if not n.left and not n.right:
                    return depth
                q.extend(c for c in (n.left, n.right) if c)
            depth += 1


if __name__ == "__main__":
    assert Solution().minDepth(make_tree([3, 9, 20, None, None, 15, 7])) == 2
    assert Solution().minDepth(make_tree([2, None, 3, None, 4, None, 5, None, 6])) == 5
    print("OK")
