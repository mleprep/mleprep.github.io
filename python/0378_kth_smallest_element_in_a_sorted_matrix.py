# 378. Kth Smallest Element in a Sorted Matrix (Medium)
# https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
# Run: python 0378_kth_smallest_element_in_a_sorted_matrix.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n = len(matrix)
        h = [(matrix[r][0], r, 0) for r in range(min(n, k))]
        heapify(h)
        for _ in range(k - 1):
            _, r, c = heappop(h)
            if c + 1 < n:
                heappush(h, (matrix[r][c + 1], r, c + 1))
        return h[0][0]


if __name__ == "__main__":
    assert Solution().kthSmallest([[1, 5, 9], [10, 11, 13], [12, 13, 15]], 8) == 13
    assert Solution().kthSmallest([[-5]], 1) == -5
    print("OK")
