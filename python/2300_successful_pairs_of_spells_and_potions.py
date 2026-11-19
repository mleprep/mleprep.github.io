# 2300. Successful Pairs of Spells and Potions (Medium)
# https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
# Run: python 2300_successful_pairs_of_spells_and_potions.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions = sorted(potions)
        m = len(potions)
        return [m - bisect.bisect_left(potions, (success + s - 1) // s) for s in spells]


if __name__ == "__main__":
    assert Solution().successfulPairs([5, 1, 3], [1, 2, 3, 4, 5], 7) == [4, 0, 3]
    assert Solution().successfulPairs([3, 1, 2], [8, 5, 8], 16) == [2, 0, 2]
    assert Solution().successfulPairs([100000], [100000], 10000000000) == [1]
    print("OK")
