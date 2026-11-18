# 846. Hand of Straights (Medium)
# https://leetcode.com/problems/hand-of-straights/
# Run: python 0846_hand_of_straights.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isNStraightHand(self, hand: List[int], groupSize: int) -> bool:
        if len(hand) % groupSize:
            return False
        cnt = Counter(hand)
        for x in sorted(cnt):
            c = cnt[x]
            if c == 0:
                continue
            for y in range(x, x + groupSize):
                if cnt[y] < c:
                    return False
                cnt[y] -= c
        return True


if __name__ == "__main__":
    assert Solution().isNStraightHand([1, 2, 3, 6, 2, 3, 4, 7, 8], 3) is True
    assert Solution().isNStraightHand([1, 2, 3, 4, 5], 4) is False
    print("OK")
