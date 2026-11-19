# 881. Boats to Save People (Medium)
# https://leetcode.com/problems/boats-to-save-people/
# Run: python 0881_boats_to_save_people.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        l, r, boats = 0, len(people) - 1, 0
        while l <= r:
            if people[l] + people[r] <= limit:
                l += 1
            r -= 1
            boats += 1
        return boats


if __name__ == "__main__":
    assert Solution().numRescueBoats([3, 2, 2, 1], 3) == 3
    assert Solution().numRescueBoats([3, 5, 3, 4], 5) == 4
    print("OK")
