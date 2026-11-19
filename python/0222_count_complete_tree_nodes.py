# 222. Count Complete Tree Nodes (Medium)
# https://leetcode.com/problems/count-complete-tree-nodes/
# Run: python 0222_count_complete_tree_nodes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        l = r = root
        hl = hr = 0
        while l:
            l, hl = l.left, hl + 1
        while r:
            r, hr = r.right, hr + 1
        if hl == hr:
            return (1 << hl) - 1
        return 1 + self.countNodes(root.left) + self.countNodes(root.right)


if __name__ == "__main__":
    assert Solution().countNodes(make_tree([1, 2, 3, 4, 5, 6])) == 6
    assert Solution().countNodes(make_tree(list(range(1, 16)))) == 15
    assert Solution().countNodes(None) == 0
    print("OK")
