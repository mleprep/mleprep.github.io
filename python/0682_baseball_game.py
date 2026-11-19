# 682. Baseball Game (Easy)
# https://leetcode.com/problems/baseball-game/
# Run: python 0682_baseball_game.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def calPoints(self, operations: List[str]) -> int:
        st = []
        for op in operations:
            if op == "C":
                st.pop()
            elif op == "D":
                st.append(2 * st[-1])
            elif op == "+":
                st.append(st[-1] + st[-2])
            else:
                st.append(int(op))
        return sum(st)


if __name__ == "__main__":
    assert Solution().calPoints(["5", "2", "C", "D", "+"]) == 30
    assert Solution().calPoints(["5", "-2", "4", "C", "D", "9", "+", "+"]) == 27
    assert Solution().calPoints(["1", "C"]) == 0
    print("OK")
