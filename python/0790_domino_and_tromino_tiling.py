# 790. Domino and Tromino Tiling (Medium)
# https://leetcode.com/problems/domino-and-tromino-tiling/
# Run: python 0790_domino_and_tromino_tiling.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numTilings(self, n: int) -> int:
        MOD = 10**9 + 7
        a, b, c = 1, 1, 2                       # f(0), f(1), f(2)
        if n < 3:
            return [1, 1, 2][n]
        for _ in range(n - 2):
            a, b, c = b, c, (2 * c + a) % MOD
        return c


if __name__ == "__main__":
    assert [Solution().numTilings(n) for n in range(1, 7)] == [1, 2, 5, 11, 24, 53]
    # cross-check against the two-state (full / jagged) DP for larger n
    f, p = [1, 1], [0, 1]                           # p(n): a 2 x n board plus one protruding cell
    for n in range(2, 1001):
        f.append((f[n - 1] + f[n - 2] + 2 * p[n - 2]) % (10**9 + 7))
        p.append((p[n - 1] + f[n - 1]) % (10**9 + 7))
    assert Solution().numTilings(1000) == f[1000]
    assert Solution().numTilings(30) == f[30]
    print("OK")
