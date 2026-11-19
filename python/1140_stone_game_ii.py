# 1140. Stone Game II (Medium)
# https://leetcode.com/problems/stone-game-ii/
# Run: python 1140_stone_game_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def stoneGameII(self, piles: List[int]) -> int:
        n = len(piles)
        suffix = list(itertools.accumulate(reversed(piles)))[::-1] + [0]

        @cache
        def best(i, m):
            if i + 2 * m >= n:
                return suffix[i]
            return max(suffix[i] - best(i + x, max(m, x)) for x in range(1, 2 * m + 1))

        return best(0, 1)


if __name__ == "__main__":
    assert Solution().stoneGameII([2, 7, 9, 4, 4]) == 10
    assert Solution().stoneGameII([1, 2, 3, 4, 5, 100]) == 104
    assert Solution().stoneGameII([5]) == 5
    print("OK")
