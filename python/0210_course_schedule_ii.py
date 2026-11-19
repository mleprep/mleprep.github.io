# 210. Course Schedule II (Medium)
# https://leetcode.com/problems/course-schedule-ii/
# Run: python 0210_course_schedule_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        g, indeg = defaultdict(list), [0] * numCourses
        for a, b in prerequisites:
            g[b].append(a)
            indeg[a] += 1
        q = deque(i for i in range(numCourses) if indeg[i] == 0)
        order = []
        while q:
            u = q.popleft()
            order.append(u)
            for v in g[u]:
                indeg[v] -= 1
                if indeg[v] == 0:
                    q.append(v)
        return order if len(order) == numCourses else []


if __name__ == "__main__":
    # any valid topological order is accepted
    def valid(n, pre, order):
        pos = {c: i for i, c in enumerate(order)}
        return len(order) == n == len(pos) and all(pos[b] < pos[a] for a, b in pre)
    assert valid(4, [[1, 0], [2, 0], [3, 1], [3, 2]], Solution().findOrder(4, [[1, 0], [2, 0], [3, 1], [3, 2]]))
    assert valid(2, [[1, 0]], Solution().findOrder(2, [[1, 0]]))
    assert Solution().findOrder(2, [[0, 1], [1, 0]]) == []
    print("OK")
