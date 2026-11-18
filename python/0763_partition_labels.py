# 763. Partition Labels (Medium)
# https://leetcode.com/problems/partition-labels/
# Run: python 0763_partition_labels.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        last = {ch: i for i, ch in enumerate(s)}
        res, start, end = [], 0, 0
        for i, ch in enumerate(s):
            end = max(end, last[ch])
            if i == end:
                res.append(end - start + 1)
                start = i + 1
        return res


if __name__ == "__main__":
    assert Solution().partitionLabels("ababcbacadefegdehijhklij") == [9, 7, 8]
    assert Solution().partitionLabels("eccbbbbdec") == [10]
    print("OK")
