# 451. Sort Characters By Frequency (Medium)
# https://leetcode.com/problems/sort-characters-by-frequency/
# Run: python 0451_sort_characters_by_frequency.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def frequencySort(self, s: str) -> str:
        return "".join(ch * c for ch, c in Counter(s).most_common())


if __name__ == "__main__":
    assert Solution().frequencySort("tree") in ("eert", "eetr")
    assert Solution().frequencySort("Aabb") in ("bbAa", "bbaA")
    print("OK")
