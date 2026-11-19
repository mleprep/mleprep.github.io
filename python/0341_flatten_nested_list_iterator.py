# 341. Flatten Nested List Iterator (Medium)
# https://leetcode.com/problems/flatten-nested-list-iterator/
# Run: python 0341_flatten_nested_list_iterator.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class NestedInteger:  # LeetCode's interface (minimal stand-in)
    def __init__(self, v): self.v = v
    def isInteger(self): return isinstance(self.v, int)
    def getInteger(self): return self.v
    def getList(self): return self.v

class NestedIterator:
    def __init__(self, nestedList: List["NestedInteger"]):
        self.st = nestedList[::-1]

    def next(self) -> int:
        return self.st.pop().getInteger()      # hasNext() guarantees an integer on top

    def hasNext(self) -> bool:
        while self.st and not self.st[-1].isInteger():
            self.st.extend(reversed(self.st.pop().getList()))
        return bool(self.st)


if __name__ == "__main__":
    def wrap(x): return NestedInteger(x if isinstance(x, int) else [wrap(y) for y in x])
    it, out = NestedIterator([wrap(x) for x in [[1, 1], 2, [1, 1]]]), []
    while it.hasNext():
        out.append(it.next())
    assert out == [1, 1, 2, 1, 1]
    it, out = NestedIterator([wrap(x) for x in [1, [4, [6]], [[]]]]), []
    while it.hasNext():
        out.append(it.next())
    assert out == [1, 4, 6]
    print("OK")
