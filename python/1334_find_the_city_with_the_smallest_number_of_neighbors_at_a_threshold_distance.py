# 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance (Medium)
# https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
# Run: python 1334_find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findTheCity(self, n: int, edges: List[List[int]], distanceThreshold: int) -> int:
        INF = float("inf")
        d = [[0 if i == j else INF for j in range(n)] for i in range(n)]
        for a, b, w in edges:
            d[a][b] = d[b][a] = w
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    if d[i][k] + d[k][j] < d[i][j]:
                        d[i][j] = d[i][k] + d[k][j]
        best, city = INF, -1
        for i in range(n):
            cnt = sum(d[i][j] <= distanceThreshold for j in range(n)) - 1
            if cnt <= best:
                best, city = cnt, i
        return city


if __name__ == "__main__":
    assert Solution().findTheCity(4, [[0, 1, 3], [1, 2, 1], [1, 3, 4], [2, 3, 1]], 4) == 3
    assert Solution().findTheCity(5, [[0, 1, 2], [0, 4, 8], [1, 2, 3], [1, 4, 2], [2, 3, 1], [3, 4, 1]], 2) == 0
    print("OK")
