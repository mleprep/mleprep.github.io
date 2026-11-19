# 1462. Course Schedule IV (Medium)
# https://leetcode.com/problems/course-schedule-iv/
# Run: python 1462_course_schedule_iv.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def checkIfPrerequisite(self, numCourses: int, prerequisites: List[List[int]], queries: List[List[int]]) -> List[bool]:
        n = numCourses
        reach = [[False] * n for _ in range(n)]
        for a, b in prerequisites:
            reach[a][b] = True
        for k in range(n):
            for i in range(n):
                if reach[i][k]:
                    row_k, row_i = reach[k], reach[i]
                    for j in range(n):
                        if row_k[j]:
                            row_i[j] = True
        return [reach[u][v] for u, v in queries]


if __name__ == "__main__":
    assert Solution().checkIfPrerequisite(2, [[1, 0]], [[0, 1], [1, 0]]) == [False, True]
    assert Solution().checkIfPrerequisite(2, [], [[1, 0], [0, 1]]) == [False, False]
    assert Solution().checkIfPrerequisite(3, [[1, 2], [1, 0], [2, 0]], [[1, 0], [1, 2]]) == [True, True]
    assert Solution().checkIfPrerequisite(4, [[0, 1], [1, 2], [2, 3]], [[0, 3], [3, 0], [1, 3]]) == [True, False, True]
    print("OK")
