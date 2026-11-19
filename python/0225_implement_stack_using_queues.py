# 225. Implement Stack using Queues (Easy)
# https://leetcode.com/problems/implement-stack-using-queues/
# Run: python 0225_implement_stack_using_queues.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyStack:
    def __init__(self):
        self.q = deque()

    def push(self, x: int) -> None:
        self.q.append(x)
        for _ in range(len(self.q) - 1):
            self.q.append(self.q.popleft())

    def pop(self) -> int:
        return self.q.popleft()

    def top(self) -> int:
        return self.q[0]

    def empty(self) -> bool:
        return not self.q


if __name__ == "__main__":
    s = MyStack()
    s.push(1); s.push(2)
    assert s.top() == 2 and s.pop() == 2 and not s.empty() and s.pop() == 1 and s.empty()
    print("OK")
