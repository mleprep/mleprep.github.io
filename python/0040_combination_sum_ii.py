# 40. Combination Sum II (Medium)
# https://leetcode.com/problems/combination-sum-ii/
# Run: python 0040_combination_sum_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res, path = [], []

        def bt(start, rest):
            if rest == 0:
                res.append(path[:])
                return
            for i in range(start, len(candidates)):
                if i > start and candidates[i] == candidates[i - 1]:
                    continue
                if candidates[i] > rest:
                    break
                path.append(candidates[i])
                bt(i + 1, rest - candidates[i])
                path.pop()

        bt(0, target)
        return res


if __name__ == "__main__":
    assert sorted(map(sorted, Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))) == sorted(map(sorted, [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]))
    assert sorted(map(sorted, Solution().combinationSum2([2, 5, 2, 1, 2], 5))) == [[1, 2, 2], [5]]
    assert Solution().combinationSum2([1], 2) == []
    print("OK")
