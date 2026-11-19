# 605. Can Place Flowers (Easy)
# https://leetcode.com/problems/can-place-flowers/
# Run: python 0605_can_place_flowers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        bed = [0] + flowerbed + [0]
        for i in range(1, len(bed) - 1):
            if bed[i - 1] == bed[i] == bed[i + 1] == 0:
                bed[i] = 1
                n -= 1
        return n <= 0


if __name__ == "__main__":
    assert Solution().canPlaceFlowers([1, 0, 0, 0, 1], 1) is True
    assert Solution().canPlaceFlowers([1, 0, 0, 0, 1], 2) is False
    assert Solution().canPlaceFlowers([0], 1) is True
    assert Solution().canPlaceFlowers([0, 0, 1, 0, 0], 2) is True
    assert Solution().canPlaceFlowers([1, 0, 1, 0, 1], 0) is True
    print("OK")
