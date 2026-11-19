# 301. Remove Invalid Parentheses (Hard)
# https://leetcode.com/problems/remove-invalid-parentheses/
# Run: python 0301_remove_invalid_parentheses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        def valid(t):
            bal = 0
            for ch in t:
                if ch == "(":
                    bal += 1
                elif ch == ")":
                    bal -= 1
                    if bal < 0:
                        return False
            return bal == 0

        level = {s}
        while True:
            good = [t for t in level if valid(t)]
            if good:
                return good
            level = {t[:i] + t[i + 1:] for t in level for i in range(len(t)) if t[i] in "()"}


if __name__ == "__main__":
    assert sorted(Solution().removeInvalidParentheses("()())()")) == ["(())()", "()()()"]
    assert sorted(Solution().removeInvalidParentheses("(a)())()")) == ["(a())()", "(a)()()"]
    assert Solution().removeInvalidParentheses(")(") == [""]
    print("OK")
