# 1423. Maximum Points You Can Obtain from Cards (Medium)
# https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
# Run: python 1423_maximum_points_you_can_obtain_from_cards.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        n, w = len(cardPoints), len(cardPoints) - k
        s = best = sum(cardPoints[:w])
        for i in range(w, n):
            s += cardPoints[i] - cardPoints[i - w]
            best = min(best, s)
        return sum(cardPoints) - best


if __name__ == "__main__":
    assert Solution().maxScore([1, 2, 3, 4, 5, 6, 1], 3) == 12
    assert Solution().maxScore([9, 7, 7, 9, 7, 7, 9], 7) == 55
    print("OK")
