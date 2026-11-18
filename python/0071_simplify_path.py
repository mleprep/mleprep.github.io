# 71. Simplify Path (Medium)
# https://leetcode.com/problems/simplify-path/
# Run: python 0071_simplify_path.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def simplifyPath(self, path: str) -> str:
        st = []
        for part in path.split("/"):
            if part == "..":
                if st:
                    st.pop()
            elif part and part != ".":
                st.append(part)
        return "/" + "/".join(st)


if __name__ == "__main__":
    assert Solution().simplifyPath("/home/") == "/home"
    assert Solution().simplifyPath("/../") == "/"
    assert Solution().simplifyPath("/home//foo/") == "/home/foo"
    assert Solution().simplifyPath("/a/./b/../../c/") == "/c"
    print("OK")
