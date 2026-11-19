# 144. Binary Tree Preorder Traversal (Easy)
# https://leetcode.com/problems/binary-tree-preorder-traversal/
# Run: python 0144_binary_tree_preorder_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        out, st = [], [root] if root else []
        while st:
            node = st.pop()
            out.append(node.val)
            if node.right:
                st.append(node.right)
            if node.left:
                st.append(node.left)
        return out


if __name__ == "__main__":
    assert Solution().preorderTraversal(make_tree([1, None, 2, 3])) == [1, 2, 3]
    assert Solution().preorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [1, 2, 4, 5, 6, 7, 3, 8, 9]
    assert Solution().preorderTraversal(make_tree([])) == []
    print("OK")
