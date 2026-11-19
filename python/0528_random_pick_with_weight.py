# 528. Random Pick with Weight (Medium)
# https://leetcode.com/problems/random-pick-with-weight/
# Run: python 0528_random_pick_with_weight.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def __init__(self, w: List[int]):
        self.pre = list(itertools.accumulate(w))

    def pickIndex(self) -> int:
        r = random.random() * self.pre[-1]
        return bisect.bisect_right(self.pre, r)


if __name__ == "__main__":
    random.seed(1)
    s = Solution([1, 3])
    picks = [s.pickIndex() for _ in range(4000)]
    assert set(picks) == {0, 1} and 0.70 < picks.count(1) / 4000 < 0.80
    assert Solution([5]).pickIndex() == 0
    print("OK")
