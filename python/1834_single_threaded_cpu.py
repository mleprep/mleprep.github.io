# 1834. Single-Threaded CPU (Medium)
# https://leetcode.com/problems/single-threaded-cpu/
# Run: python 1834_single_threaded_cpu.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        order = sorted(range(len(tasks)), key=lambda i: tasks[i][0])
        h, res, time, i = [], [], 0, 0
        while len(res) < len(tasks):
            if not h and time < tasks[order[i]][0]:
                time = tasks[order[i]][0]
            while i < len(order) and tasks[order[i]][0] <= time:
                j = order[i]
                heappush(h, (tasks[j][1], j))
                i += 1
            p, j = heappop(h)
            time += p
            res.append(j)
        return res


if __name__ == "__main__":
    assert Solution().getOrder([[1, 2], [2, 4], [3, 2], [4, 1]]) == [0, 2, 3, 1]
    assert Solution().getOrder([[7, 10], [7, 12], [7, 5], [7, 4], [7, 2]]) == [4, 3, 2, 0, 1]
    print("OK")
