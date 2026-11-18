# 20. Valid Parentheses (Easy)
# https://leetcode.com/problems/valid-parentheses/
# Run: python 0020_valid_parentheses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isValid(self, s: str) -> bool:
        pair = {")": "(", "]": "[", "}": "{"}
        st = []
        for ch in s:
            if ch in pair:
                if not st or st.pop() != pair[ch]:
                    return False
            else:
                st.append(ch)
        return not st


if __name__ == "__main__":
    assert Solution().isValid("()[]{}") is True
    assert Solution().isValid("(]") is False
    assert Solution().isValid("([)]") is False
    assert Solution().isValid("{[]}") is True
    print("OK")
