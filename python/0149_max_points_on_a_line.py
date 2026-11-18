# 149. Max Points on a Line (Hard)
# https://leetcode.com/problems/max-points-on-a-line/
# Run: python 0149_max_points_on_a_line.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        best = 1
        for i, (x1, y1) in enumerate(points):
            slopes = defaultdict(int)
            for x2, y2 in points[i + 1:]:
                dx, dy = x2 - x1, y2 - y1
                g = math.gcd(dx, dy)
                dx, dy = dx // g, dy // g
                if dx < 0 or (dx == 0 and dy < 0):
                    dx, dy = -dx, -dy
                slopes[(dx, dy)] += 1
                best = max(best, slopes[(dx, dy)] + 1)
        return best


if __name__ == "__main__":
    assert Solution().maxPoints([[1, 1], [2, 2], [3, 3]]) == 3
    assert Solution().maxPoints([[1, 1], [3, 2], [5, 3], [4, 1], [2, 3], [1, 4]]) == 4
    print("OK")
