# 1207. Unique Number of Occurrences (Easy)
# https://leetcode.com/problems/unique-number-of-occurrences/
# Run: python 1207_unique_number_of_occurrences.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        counts = list(Counter(arr).values())
        return len(counts) == len(set(counts))


if __name__ == "__main__":
    assert Solution().uniqueOccurrences([1, 2, 2, 1, 1, 3]) is True
    assert Solution().uniqueOccurrences([1, 2]) is False
    assert Solution().uniqueOccurrences([-3, 0, 1, -3, 1, 1, 1, -3, 10, 0]) is True
    print("OK")
