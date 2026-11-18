# 230. Kth Smallest Element in a BST (Medium)
# https://leetcode.com/problems/kth-smallest-element-in-a-bst/
# Run: python 0230_kth_smallest_element_in_a_bst.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        st, node = [], root
        while True:
            while node:
                st.append(node)
                node = node.left
            node = st.pop()
            k -= 1
            if k == 0:
                return node.val
            node = node.right


if __name__ == "__main__":
    assert Solution().kthSmallest(make_tree([3, 1, 4, None, 2]), 1) == 1
    assert Solution().kthSmallest(make_tree([5, 3, 6, 2, 4, None, None, 1]), 3) == 3
    print("OK")
