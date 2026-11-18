# 84. Largest Rectangle in Histogram (Hard)
# https://leetcode.com/problems/largest-rectangle-in-histogram/
# Run: python 0084_largest_rectangle_in_histogram.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        st, best = [], 0              # increasing stack of indices
        for i, h in enumerate(heights + [0]):
            while st and heights[st[-1]] >= h:
                top = heights[st.pop()]
                left = st[-1] if st else -1
                best = max(best, top * (i - left - 1))
            st.append(i)
        return best


if __name__ == "__main__":
    assert Solution().largestRectangleArea([2, 1, 5, 6, 2, 3]) == 10
    assert Solution().largestRectangleArea([2, 4]) == 4
    print("OK")
