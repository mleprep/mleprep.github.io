# 155. Min Stack (Medium)
# https://leetcode.com/problems/min-stack/
# Run: python 0155_min_stack.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MinStack:
    def __init__(self):
        self.st = []  # (val, current min)

    def push(self, val: int) -> None:
        self.st.append((val, min(val, self.st[-1][1]) if self.st else val))

    def pop(self) -> None:
        self.st.pop()

    def top(self) -> int:
        return self.st[-1][0]

    def getMin(self) -> int:
        return self.st[-1][1]


if __name__ == "__main__":
    m = MinStack()
    m.push(-2); m.push(0); m.push(-3)
    assert m.getMin() == -3
    m.pop()
    assert m.top() == 0 and m.getMin() == -2
    print("OK")
