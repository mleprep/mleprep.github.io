# 621. Task Scheduler (Medium)
# https://leetcode.com/problems/task-scheduler/
# Run: python 0621_task_scheduler.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        cnt = Counter(tasks).values()
        maxf = max(cnt)
        tied = sum(1 for c in cnt if c == maxf)
        return max(len(tasks), (maxf - 1) * (n + 1) + tied)


if __name__ == "__main__":
    assert Solution().leastInterval(["A", "A", "A", "B", "B", "B"], 2) == 8
    assert Solution().leastInterval(["A", "C", "A", "B", "D", "B"], 1) == 6
    assert Solution().leastInterval(["A", "A", "A", "B", "B", "B"], 3) == 10
    print("OK")
