# 437. Path Sum III (Medium)
# https://leetcode.com/problems/path-sum-iii/
# Run: python 0437_path_sum_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        seen = defaultdict(int, {0: 1})

        def dfs(node, cur):
            if not node:
                return 0
            cur += node.val
            cnt = seen[cur - targetSum]
            seen[cur] += 1
            cnt += dfs(node.left, cur) + dfs(node.right, cur)
            seen[cur] -= 1                     # backtrack
            return cnt

        return dfs(root, 0)


if __name__ == "__main__":
    assert Solution().pathSum(make_tree([10, 5, -3, 3, 2, None, 11, 3, -2, None, 1]), 8) == 3
    assert Solution().pathSum(make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1]), 22) == 3
    print("OK")
