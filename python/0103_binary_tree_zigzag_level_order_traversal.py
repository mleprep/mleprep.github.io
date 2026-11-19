# 103. Binary Tree Zigzag Level Order Traversal (Medium)
# https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
# Run: python 0103_binary_tree_zigzag_level_order_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        res, q = [], deque([root] if root else [])
        while q:
            level = []
            for _ in range(len(q)):
                node = q.popleft()
                level.append(node.val)
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
            res.append(level[::-1] if len(res) % 2 else level)
        return res


if __name__ == "__main__":
    assert Solution().zigzagLevelOrder(make_tree([3, 9, 20, None, None, 15, 7])) == [[3], [20, 9], [15, 7]]
    assert Solution().zigzagLevelOrder(make_tree([1])) == [[1]]
    assert Solution().zigzagLevelOrder(make_tree([])) == []
    assert Solution().zigzagLevelOrder(make_tree([1, 2, 3, 4, None, None, 5])) == [[1], [3, 2], [4, 5]]
    print("OK")
