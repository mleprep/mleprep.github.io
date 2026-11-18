# 622. Design Circular Queue (Medium)
# https://leetcode.com/problems/design-circular-queue/
# Run: python 0622_design_circular_queue.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class MyCircularQueue:
    def __init__(self, k: int):
        self.a, self.head, self.n = [0] * k, 0, 0

    def enQueue(self, value: int) -> bool:
        if self.isFull():
            return False
        self.a[(self.head + self.n) % len(self.a)] = value
        self.n += 1
        return True

    def deQueue(self) -> bool:
        if self.isEmpty():
            return False
        self.head = (self.head + 1) % len(self.a)
        self.n -= 1
        return True

    def Front(self) -> int:
        return -1 if self.isEmpty() else self.a[self.head]

    def Rear(self) -> int:
        return -1 if self.isEmpty() else self.a[(self.head + self.n - 1) % len(self.a)]

    def isEmpty(self) -> bool:
        return self.n == 0

    def isFull(self) -> bool:
        return self.n == len(self.a)


if __name__ == "__main__":
    cq = MyCircularQueue(3)
    assert cq.enQueue(1) and cq.enQueue(2) and cq.enQueue(3) and not cq.enQueue(4)
    assert cq.Rear() == 3 and cq.isFull() and cq.deQueue() and cq.enQueue(4) and cq.Rear() == 4
    print("OK")
