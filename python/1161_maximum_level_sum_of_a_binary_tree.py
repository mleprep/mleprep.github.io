# 1161. Maximum Level Sum of a Binary Tree (Medium)
# https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
# Run: python 1161_maximum_level_sum_of_a_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        best, best_level, level, q = float("-inf"), 0, 0, deque([root])
        while q:
            level += 1
            total = 0
            for _ in range(len(q)):
                n = q.popleft()
                total += n.val
                q.extend(c for c in (n.left, n.right) if c)
            if total > best:
                best, best_level = total, level
        return best_level


if __name__ == "__main__":
    assert Solution().maxLevelSum(make_tree([1, 7, 0, 7, -8, None, None])) == 2
    assert Solution().maxLevelSum(make_tree([989, None, 10250, 98693, -89388, None, None, None, -32127])) == 2
    assert Solution().maxLevelSum(make_tree([-1, -2, -3])) == 1
    print("OK")
