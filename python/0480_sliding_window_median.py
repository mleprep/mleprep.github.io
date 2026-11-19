# 480. Sliding Window Median (Hard)
# https://leetcode.com/problems/sliding-window-median/
# Run: python 0480_sliding_window_median.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        win = sorted(nums[:k])
        res = []
        for i in range(k, len(nums) + 1):
            res.append((win[k // 2] + win[(k - 1) // 2]) / 2)
            if i == len(nums):
                break
            win.pop(bisect.bisect_left(win, nums[i - k]))
            bisect.insort(win, nums[i])
        return res


if __name__ == "__main__":
    assert Solution().medianSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3) == [1.0, -1.0, -1.0, 3.0, 5.0, 6.0]
    assert Solution().medianSlidingWindow([1, 2, 3, 4, 2, 3, 1, 4, 2], 3) == [2.0, 3.0, 3.0, 3.0, 2.0, 3.0, 2.0]
    print("OK")
