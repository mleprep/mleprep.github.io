# 199. Binary Tree Right Side View (Medium)
# https://leetcode.com/problems/binary-tree-right-side-view/
# Run: python 0199_binary_tree_right_side_view.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        res, q = [], deque([root] if root else [])
        while q:
            res.append(q[-1].val)
            for _ in range(len(q)):
                node = q.popleft()
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
        return res


if __name__ == "__main__":
    assert Solution().rightSideView(make_tree([1, 2, 3, None, 5, None, 4])) == [1, 3, 4]
    assert Solution().rightSideView(make_tree([1, 2, 3, 4])) == [1, 3, 4]
    print("OK")
