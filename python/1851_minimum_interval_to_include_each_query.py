# 1851. Minimum Interval to Include Each Query (Hard)
# https://leetcode.com/problems/minimum-interval-to-include-each-query/
# Run: python 1851_minimum_interval_to_include_each_query.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:
        intervals.sort()
        h, i, ans = [], 0, {}
        for q in sorted(set(queries)):
            while i < len(intervals) and intervals[i][0] <= q:
                l, r = intervals[i]
                heappush(h, (r - l + 1, r))
                i += 1
            while h and h[0][1] < q:
                heappop(h)
            ans[q] = h[0][0] if h else -1
        return [ans[q] for q in queries]


if __name__ == "__main__":
    assert Solution().minInterval([[1, 4], [2, 4], [3, 6], [4, 4]], [2, 3, 4, 5]) == [3, 3, 1, 4]
    assert Solution().minInterval([[2, 3], [2, 5], [1, 8], [20, 25]], [2, 19, 5, 22]) == [2, -1, 4, 6]
    print("OK")
