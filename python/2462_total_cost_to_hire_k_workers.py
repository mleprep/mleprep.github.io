# 2462. Total Cost to Hire K Workers (Medium)
# https://leetcode.com/problems/total-cost-to-hire-k-workers/
# Run: python 2462_total_cost_to_hire_k_workers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def totalCost(self, costs: List[int], k: int, candidates: int) -> int:
        i, j = 0, len(costs) - 1
        front, back, total = [], [], 0
        for _ in range(k):
            while len(front) < candidates and i <= j:
                heappush(front, costs[i])
                i += 1
            while len(back) < candidates and i <= j:
                heappush(back, costs[j])
                j -= 1
            if not back or (front and front[0] <= back[0]):
                total += heappop(front)
            else:
                total += heappop(back)
        return total


if __name__ == "__main__":
    assert Solution().totalCost([17, 12, 10, 2, 7, 2, 11, 20, 8], 3, 4) == 11
    assert Solution().totalCost([1, 2, 4, 1], 3, 3) == 4
    assert Solution().totalCost([5], 1, 1) == 5
    assert Solution().totalCost([2, 2, 2, 2], 2, 1) == 4
    print("OK")
