# 933. Number of Recent Calls (Easy)
# https://leetcode.com/problems/number-of-recent-calls/
# Run: python 0933_number_of_recent_calls.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class RecentCounter:
    def __init__(self):
        self.q = deque()

    def ping(self, t: int) -> int:
        self.q.append(t)
        while self.q[0] < t - 3000:
            self.q.popleft()
        return len(self.q)


if __name__ == "__main__":
    rc = RecentCounter()
    assert [rc.ping(t) for t in (1, 100, 3001, 3002)] == [1, 2, 3, 3]
    rc2 = RecentCounter()
    assert [rc2.ping(t) for t in (1, 3001, 3002, 6003)] == [1, 2, 2, 1]
    print("OK")
