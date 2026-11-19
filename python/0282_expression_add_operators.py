# 282. Expression Add Operators (Hard)
# https://leetcode.com/problems/expression-add-operators/
# Run: python 0282_expression_add_operators.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def addOperators(self, num: str, target: int) -> List[str]:
        res = []

        def bt(i, expr, value, prev):
            if i == len(num):
                if value == target:
                    res.append(expr)
                return
            for j in range(i + 1, len(num) + 1):
                s = num[i:j]
                if len(s) > 1 and s[0] == "0":
                    break
                cur = int(s)
                if i == 0:
                    bt(j, s, cur, cur)
                else:
                    bt(j, expr + "+" + s, value + cur, cur)
                    bt(j, expr + "-" + s, value - cur, -cur)
                    bt(j, expr + "*" + s, value - prev + prev * cur, prev * cur)

        bt(0, "", 0, 0)
        return res


if __name__ == "__main__":
    assert sorted(Solution().addOperators("123", 6)) == ["1*2*3", "1+2+3"]
    assert sorted(Solution().addOperators("232", 8)) == ["2*3+2", "2+3*2"]
    assert Solution().addOperators("3456237490", 9191) == []
    assert sorted(Solution().addOperators("105", 5)) == ["1*0+5", "10-5"]
    print("OK")
