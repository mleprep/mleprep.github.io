# 134. Gas Station (Medium)
# https://leetcode.com/problems/gas-station/
# Run: python 0134_gas_station.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        if sum(gas) < sum(cost):
            return -1
        start = tank = 0
        for i, (g, c) in enumerate(zip(gas, cost)):
            tank += g - c
            if tank < 0:
                start, tank = i + 1, 0
        return start


if __name__ == "__main__":
    assert Solution().canCompleteCircuit([1, 2, 3, 4, 5], [3, 4, 5, 1, 2]) == 3
    assert Solution().canCompleteCircuit([2, 3, 4], [3, 4, 3]) == -1
    print("OK")
