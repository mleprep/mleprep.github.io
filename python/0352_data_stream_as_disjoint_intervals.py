# 352. Data Stream as Disjoint Intervals (Hard)
# https://leetcode.com/problems/data-stream-as-disjoint-intervals/
# Run: python 0352_data_stream_as_disjoint_intervals.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class SummaryRanges:
    def __init__(self):
        self.iv = []                                   # sorted, disjoint [start, end]

    def addNum(self, value: int) -> None:
        iv = self.iv
        i = bisect.bisect_left(iv, [value + 1])        # first interval starting > value
        if i > 0 and iv[i - 1][1] >= value:            # already covered
            return
        left = i > 0 and iv[i - 1][1] == value - 1
        right = i < len(iv) and iv[i][0] == value + 1
        if left and right:
            iv[i - 1][1] = iv[i][1]
            iv.pop(i)
        elif left:
            iv[i - 1][1] = value
        elif right:
            iv[i][0] = value
        else:
            iv.insert(i, [value, value])

    def getIntervals(self) -> List[List[int]]:
        return [x[:] for x in self.iv]


if __name__ == "__main__":
    sr = SummaryRanges()
    out = []
    for v in (1, 3, 7, 2, 6):
        sr.addNum(v)
        out.append(sr.getIntervals())
    assert out == [[[1, 1]], [[1, 1], [3, 3]], [[1, 1], [3, 3], [7, 7]], [[1, 3], [7, 7]], [[1, 3], [6, 7]]]
    print("OK")
