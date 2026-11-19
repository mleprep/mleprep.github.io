# 637. Average of Levels in Binary Tree (Easy)
# https://leetcode.com/problems/average-of-levels-in-binary-tree/
# Run: python 0637_average_of_levels_in_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        res, q = [], deque([root])
        while q:
            s, k = 0, len(q)
            for _ in range(k):
                n = q.popleft()
                s += n.val
                q.extend(c for c in (n.left, n.right) if c)
            res.append(s / k)
        return res


if __name__ == "__main__":
    assert Solution().averageOfLevels(make_tree([3, 9, 20, None, None, 15, 7])) == [3.0, 14.5, 11.0]
    print("OK")
