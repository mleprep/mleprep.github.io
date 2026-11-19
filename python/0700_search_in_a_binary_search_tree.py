# 700. Search in a Binary Search Tree (Easy)
# https://leetcode.com/problems/search-in-a-binary-search-tree/
# Run: python 0700_search_in_a_binary_search_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def searchBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        while root and root.val != val:
            root = root.left if val < root.val else root.right
        return root


if __name__ == "__main__":
    assert tree_vals(Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 2)) == [2, 1, 3]
    assert Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 5) is None
    assert tree_vals(Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 7)) == [7]
    print("OK")
