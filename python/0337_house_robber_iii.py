# 337. House Robber III (Medium)
# https://leetcode.com/problems/house-robber-iii/
# Run: python 0337_house_robber_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rob(self, root: Optional[TreeNode]) -> int:
        def dfs(node):
            if not node:
                return 0, 0
            lr, ls = dfs(node.left)
            rr, rs = dfs(node.right)
            return node.val + ls + rs, max(lr, ls) + max(rr, rs)

        return max(dfs(root))


if __name__ == "__main__":
    assert Solution().rob(make_tree([3, 2, 3, None, 3, None, 1])) == 7
    assert Solution().rob(make_tree([3, 4, 5, 1, 3, None, 1])) == 9
    print("OK")
