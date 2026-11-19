# 1039. Minimum Score Triangulation of Polygon (Medium)
# https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
# Run: python 1039_minimum_score_triangulation_of_polygon.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minScoreTriangulation(self, values: List[int]) -> int:
        n = len(values)
        dp = [[0] * n for _ in range(n)]
        for gap in range(2, n):
            for i in range(n - gap):
                j = i + gap
                dp[i][j] = min(dp[i][k] + values[i] * values[k] * values[j] + dp[k][j] for k in range(i + 1, j))
        return dp[0][n - 1]


if __name__ == "__main__":
    assert Solution().minScoreTriangulation([1, 2, 3]) == 6
    assert Solution().minScoreTriangulation([3, 7, 4, 5]) == 144
    assert Solution().minScoreTriangulation([1, 3, 1, 4, 1, 5]) == 13
    print("OK")
