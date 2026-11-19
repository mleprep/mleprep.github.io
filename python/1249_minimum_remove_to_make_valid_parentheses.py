# 1249. Minimum Remove to Make Valid Parentheses (Medium)
# https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
# Run: python 1249_minimum_remove_to_make_valid_parentheses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        st, bad = [], set()
        for i, ch in enumerate(s):
            if ch == "(":
                st.append(i)
            elif ch == ")":
                if st:
                    st.pop()
                else:
                    bad.add(i)
        bad.update(st)
        return "".join(ch for i, ch in enumerate(s) if i not in bad)


if __name__ == "__main__":
    assert Solution().minRemoveToMakeValid("lee(t(c)o)de)") == "lee(t(c)o)de"
    assert Solution().minRemoveToMakeValid("a)b(c)d") == "ab(c)d"
    assert Solution().minRemoveToMakeValid("))((") == ""
    print("OK")
