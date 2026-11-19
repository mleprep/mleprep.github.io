# 1047. Remove All Adjacent Duplicates In String (Easy)
# https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
# Run: python 1047_remove_all_adjacent_duplicates_in_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeDuplicates(self, s: str) -> str:
        st = []
        for ch in s:
            if st and st[-1] == ch:
                st.pop()
            else:
                st.append(ch)
        return "".join(st)


if __name__ == "__main__":
    assert Solution().removeDuplicates("abbaca") == "ca"
    assert Solution().removeDuplicates("azxxzy") == "ay"
    print("OK")
