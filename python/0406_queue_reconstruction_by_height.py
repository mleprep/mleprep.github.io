# 406. Queue Reconstruction by Height (Medium)
# https://leetcode.com/problems/queue-reconstruction-by-height/
# Run: python 0406_queue_reconstruction_by_height.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        res = []
        for h, k in sorted(people, key=lambda p: (-p[0], p[1])):
            res.insert(k, [h, k])
        return res


if __name__ == "__main__":
    P = [[7, 0], [4, 4], [7, 1], [5, 0], [6, 1], [5, 2]]
    assert Solution().reconstructQueue(P) == [[5, 0], [7, 0], [5, 2], [6, 1], [4, 4], [7, 1]]
    print("OK")
