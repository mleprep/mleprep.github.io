# 1146. Snapshot Array (Medium)
# https://leetcode.com/problems/snapshot-array/
# Run: python 1146_snapshot_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class SnapshotArray:
    def __init__(self, length: int):
        self.hist = [[(-1, 0)] for _ in range(length)]
        self.snap_id = 0

    def set(self, index: int, val: int) -> None:
        h = self.hist[index]
        if h[-1][0] == self.snap_id:
            h[-1] = (self.snap_id, val)
        else:
            h.append((self.snap_id, val))

    def snap(self) -> int:
        self.snap_id += 1
        return self.snap_id - 1

    def get(self, index: int, snap_id: int) -> int:
        h = self.hist[index]
        i = bisect.bisect_right(h, (snap_id, float("inf"))) - 1
        return h[i][1]


if __name__ == "__main__":
    sa = SnapshotArray(3)
    sa.set(0, 5)
    assert sa.snap() == 0
    sa.set(0, 6)
    assert sa.get(0, 0) == 5 and sa.get(0, 1) == 6 and sa.get(1, 0) == 0
    print("OK")
