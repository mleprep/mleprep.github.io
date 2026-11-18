# 32. Longest Valid Parentheses (Hard)
# https://leetcode.com/problems/longest-valid-parentheses/
# Run: python 0032_longest_valid_parentheses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestValidParentheses(self, s: str) -> int:
        st, best = [-1], 0
        for i, ch in enumerate(s):
            if ch == "(":
                st.append(i)
            else:
                st.pop()
                if not st:
                    st.append(i)
                else:
                    best = max(best, i - st[-1])
        return best


if __name__ == "__main__":
    assert Solution().longestValidParentheses("(()") == 2
    assert Solution().longestValidParentheses(")()())") == 4
    assert Solution().longestValidParentheses("") == 0
    print("OK")
