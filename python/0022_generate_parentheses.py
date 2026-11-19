# 22. Generate Parentheses (Medium)
# https://leetcode.com/problems/generate-parentheses/
# Run: python 0022_generate_parentheses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def bt(cur, op, cl):
            if len(cur) == 2 * n:
                res.append(cur)
                return
            if op < n:
                bt(cur + "(", op + 1, cl)
            if cl < op:
                bt(cur + ")", op, cl + 1)

        bt("", 0, 0)
        return res


if __name__ == "__main__":
    assert sorted(Solution().generateParenthesis(3)) == sorted(["((()))", "(()())", "(())()", "()(())", "()()()"])
    assert Solution().generateParenthesis(1) == ["()"]
    assert sorted(Solution().generateParenthesis(4)) == ['(((())))', '((()()))', '((())())', '((()))()', '(()(()))', '(()()())', '(()())()', '(())(())', '(())()()', '()((()))', '()(()())', '()(())()', '()()(())', '()()()()']
    print("OK")
