# 316. Remove Duplicate Letters (Medium)
# https://leetcode.com/problems/remove-duplicate-letters/
# Run: python 0316_remove_duplicate_letters.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        last = {ch: i for i, ch in enumerate(s)}
        st, used = [], set()
        for i, ch in enumerate(s):
            if ch in used:
                continue
            while st and st[-1] > ch and last[st[-1]] > i:
                used.remove(st.pop())
            st.append(ch)
            used.add(ch)
        return "".join(st)


if __name__ == "__main__":
    assert Solution().removeDuplicateLetters("bcabc") == "abc"
    assert Solution().removeDuplicateLetters("cbacdcbc") == "acdb"
    print("OK")
