# 904. Fruit Into Baskets (Medium)
# https://leetcode.com/problems/fruit-into-baskets/
# Run: python 0904_fruit_into_baskets.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        cnt = defaultdict(int)
        l = best = 0
        for r, f in enumerate(fruits):
            cnt[f] += 1
            while len(cnt) > 2:
                cnt[fruits[l]] -= 1
                if cnt[fruits[l]] == 0:
                    del cnt[fruits[l]]
                l += 1
            best = max(best, r - l + 1)
        return best


if __name__ == "__main__":
    assert Solution().totalFruit([1, 2, 1]) == 3
    assert Solution().totalFruit([0, 1, 2, 2]) == 3
    assert Solution().totalFruit([1, 2, 3, 2, 2]) == 4
    print("OK")
