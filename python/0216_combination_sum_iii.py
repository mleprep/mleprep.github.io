# 216. Combination Sum III (Medium)
# https://leetcode.com/problems/combination-sum-iii/
# Run: python 0216_combination_sum_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        res, path = [], []

        def bt(start, rest):
            if len(path) == k:
                if rest == 0:
                    res.append(path[:])
                return
            for d in range(start, 10):
                if d > rest:
                    break
                path.append(d)
                bt(d + 1, rest - d)
                path.pop()

        bt(1, n)
        return res


if __name__ == "__main__":
    assert Solution().combinationSum3(3, 7) == [[1, 2, 4]]
    assert Solution().combinationSum3(3, 9) == [[1, 2, 6], [1, 3, 5], [2, 3, 4]]
    assert Solution().combinationSum3(4, 1) == []
    print("OK")
