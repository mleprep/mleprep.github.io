# 632. Smallest Range Covering Elements from K Lists (Hard)
# https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/
# Run: python 0632_smallest_range_covering_elements_from_k_lists.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        heap = [(row[0], i, 0) for i, row in enumerate(nums)]
        heapify(heap)
        hi = max(row[0] for row in nums)
        best = [heap[0][0], hi]
        while True:
            lo, i, j = heappop(heap)
            if hi - lo < best[1] - best[0]:
                best = [lo, hi]
            if j + 1 == len(nums[i]):
                return best
            nxt = nums[i][j + 1]
            hi = max(hi, nxt)
            heappush(heap, (nxt, i, j + 1))


if __name__ == "__main__":
    assert Solution().smallestRange([[4, 10, 15, 24, 26], [0, 9, 12, 20], [5, 18, 22, 30]]) == [20, 24]
    assert Solution().smallestRange([[1, 2, 3], [1, 2, 3], [1, 2, 3]]) == [1, 1]
    assert Solution().smallestRange([[1], [2], [3]]) == [1, 3]
    assert Solution().smallestRange([[10, 10], [11, 11]]) == [10, 11]
    print("OK")
