# 394. Decode String (Medium)
# https://leetcode.com/problems/decode-string/
# Run: python 0394_decode_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def decodeString(self, s: str) -> str:
        st, cur, k = [], "", 0
        for ch in s:
            if ch.isdigit():
                k = k * 10 + int(ch)
            elif ch == "[":
                st.append((cur, k))
                cur, k = "", 0
            elif ch == "]":
                prev, times = st.pop()
                cur = prev + cur * times
            else:
                cur += ch
        return cur


if __name__ == "__main__":
    assert Solution().decodeString("3[a]2[bc]") == "aaabcbc"
    assert Solution().decodeString("3[a2[c]]") == "accaccacc"
    assert Solution().decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef"
    print("OK")
