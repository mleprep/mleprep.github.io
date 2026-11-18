# 232. Implement Queue using Stacks (Easy)
# https://leetcode.com/problems/implement-queue-using-stacks/
# Run: python 0232_implement_queue_using_stacks.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyQueue:
    def __init__(self):
        self.inp, self.out = [], []

    def push(self, x: int) -> None:
        self.inp.append(x)

    def _move(self):
        if not self.out:
            while self.inp:
                self.out.append(self.inp.pop())

    def pop(self) -> int:
        self._move()
        return self.out.pop()

    def peek(self) -> int:
        self._move()
        return self.out[-1]

    def empty(self) -> bool:
        return not self.inp and not self.out


if __name__ == "__main__":
    q = MyQueue()
    q.push(1); q.push(2)
    assert q.peek() == 1 and q.pop() == 1 and not q.empty()
    print("OK")
