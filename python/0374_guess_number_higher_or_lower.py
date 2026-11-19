# 374. Guess Number Higher or Lower (Easy)
# https://leetcode.com/problems/guess-number-higher-or-lower/
# Run: python 0374_guess_number_higher_or_lower.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def guessNumber(self, n: int) -> int:
        lo, hi = 1, n
        while True:
            mid = (lo + hi) // 2
            g = guess(mid)          # provided by LeetCode
            if g == 0:
                return mid
            if g < 0:
                hi = mid - 1
            else:
                lo = mid + 1


if __name__ == "__main__":
    for n, PICK in ((10, 6), (1, 1), (2, 1), (2, 2), (2126753390, 1702766719)):
        def guess(num, pick=PICK):
            return 0 if num == pick else (-1 if pick < num else 1)
        assert Solution().guessNumber(n) == PICK
    print("OK")
