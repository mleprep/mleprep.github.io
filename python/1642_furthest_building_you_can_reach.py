# 1642. Furthest Building You Can Reach (Medium)
# https://leetcode.com/problems/furthest-building-you-can-reach/
# Run: python 1642_furthest_building_you_can_reach.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        h = []
        for i in range(len(heights) - 1):
            d = heights[i + 1] - heights[i]
            if d <= 0:
                continue
            heappush(h, d)
            if len(h) > ladders:
                bricks -= heappop(h)
                if bricks < 0:
                    return i
        return len(heights) - 1


if __name__ == "__main__":
    assert Solution().furthestBuilding([4, 2, 7, 6, 9, 14, 12], 5, 1) == 4
    assert Solution().furthestBuilding([4, 12, 2, 7, 3, 18, 20, 3, 19], 10, 2) == 7
    assert Solution().furthestBuilding([14, 3, 19, 3], 17, 0) == 3
    print("OK")
