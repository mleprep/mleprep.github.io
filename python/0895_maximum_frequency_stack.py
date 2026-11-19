# 895. Maximum Frequency Stack (Hard)
# https://leetcode.com/problems/maximum-frequency-stack/
# Run: python 0895_maximum_frequency_stack.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class FreqStack:
    def __init__(self):
        self.freq = Counter()
        self.group = defaultdict(list)
        self.maxf = 0

    def push(self, val: int) -> None:
        f = self.freq[val] + 1
        self.freq[val] = f
        self.group[f].append(val)
        self.maxf = max(self.maxf, f)

    def pop(self) -> int:
        val = self.group[self.maxf].pop()
        self.freq[val] -= 1
        if not self.group[self.maxf]:
            self.maxf -= 1
        return val


if __name__ == "__main__":
    fs = FreqStack()
    for v in (5, 7, 5, 7, 4, 5):
        fs.push(v)
    assert [fs.pop(), fs.pop(), fs.pop(), fs.pop()] == [5, 7, 5, 4]
    fs.push(9)
    assert [fs.pop(), fs.pop(), fs.pop()] == [9, 7, 5]
    print("OK")
