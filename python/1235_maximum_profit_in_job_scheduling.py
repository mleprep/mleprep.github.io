# 1235. Maximum Profit in Job Scheduling (Hard)
# https://leetcode.com/problems/maximum-profit-in-job-scheduling/
# Run: python 1235_maximum_profit_in_job_scheduling.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
        jobs = sorted(zip(endTime, startTime, profit))
        ends = [e for e, _, _ in jobs]
        dp = [0] * (len(jobs) + 1)
        for i, (e, s, p) in enumerate(jobs, 1):
            j = bisect.bisect_right(ends, s, 0, i - 1)     # jobs finishing by time s
            dp[i] = max(dp[i - 1], dp[j] + p)
        return dp[-1]


if __name__ == "__main__":
    assert Solution().jobScheduling([1, 2, 3, 3], [3, 4, 5, 6], [50, 10, 40, 70]) == 120
    assert Solution().jobScheduling([1, 2, 3, 4, 6], [3, 5, 10, 6, 9], [20, 20, 100, 70, 60]) == 150
    assert Solution().jobScheduling([1, 1, 1], [2, 3, 4], [5, 6, 4]) == 6
    print("OK")
