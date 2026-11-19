# 145. Binary Tree Postorder Traversal (Easy)
# https://leetcode.com/problems/binary-tree-postorder-traversal/
# Run: python 0145_binary_tree_postorder_traversal.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        out, st = [], [root] if root else []
        while st:
            node = st.pop()
            out.append(node.val)
            if node.left:
                st.append(node.left)
            if node.right:
                st.append(node.right)
        return out[::-1]


if __name__ == "__main__":
    assert Solution().postorderTraversal(make_tree([1, None, 2, 3])) == [3, 2, 1]
    assert Solution().postorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [4, 6, 7, 5, 2, 9, 8, 3, 1]
    assert Solution().postorderTraversal(make_tree([])) == []
    print("OK")
