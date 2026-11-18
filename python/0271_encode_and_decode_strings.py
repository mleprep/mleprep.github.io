# 271. Encode and Decode Strings (Medium)
# https://leetcode.com/problems/encode-and-decode-strings/
# Run: python 0271_encode_and_decode_strings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Codec:
    def encode(self, strs: List[str]) -> str:
        return "".join(f"{len(s)}#{s}" for s in strs)

    def decode(self, s: str) -> List[str]:
        res, i = [], 0
        while i < len(s):
            j = s.index("#", i)
            n = int(s[i:j])
            res.append(s[j + 1:j + 1 + n])
            i = j + 1 + n
        return res


if __name__ == "__main__":
    c = Codec()
    for data in (["neet", "code", "love", "you"], ["we", "say", ":", "yes", "#4#", ""], []):
        assert c.decode(c.encode(data)) == data
    print("OK")
