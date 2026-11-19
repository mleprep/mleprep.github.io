# 997. Find the Town Judge (Easy)
# https://leetcode.com/problems/find-the-town-judge/
# Run: python 0997_find_the_town_judge.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        score = [0] * (n + 1)
        for a, b in trust:
            score[a] -= 1
            score[b] += 1
        return next((p for p in range(1, n + 1) if score[p] == n - 1), -1)


if __name__ == "__main__":
    assert Solution().findJudge(2, [[1, 2]]) == 2
    assert Solution().findJudge(3, [[1, 3], [2, 3]]) == 3
    assert Solution().findJudge(3, [[1, 3], [2, 3], [3, 1]]) == -1
    assert Solution().findJudge(1, []) == 1
    print("OK")
