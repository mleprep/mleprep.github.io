# 535. Encode and Decode TinyURL (Medium)
# https://leetcode.com/problems/encode-and-decode-tinyurl/
# Run: python 0535_encode_and_decode_tinyurl.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Codec:
    ALPHA = string.digits + string.ascii_letters

    def __init__(self):
        self.long_to_short, self.short_to_long = {}, {}

    def _base62(self, n):
        s = ""
        while True:
            n, r = divmod(n, 62)
            s = self.ALPHA[r] + s
            if n == 0:
                return s

    def encode(self, longUrl: str) -> str:
        if longUrl not in self.long_to_short:
            code = self._base62(len(self.short_to_long) + 1)
            self.long_to_short[longUrl] = code
            self.short_to_long[code] = longUrl
        return "http://tinyurl.com/" + self.long_to_short[longUrl]

    def decode(self, shortUrl: str) -> str:
        return self.short_to_long[shortUrl.rsplit("/", 1)[1]]


if __name__ == "__main__":
    c = Codec()
    u = "https://leetcode.com/problems/design-tinyurl"
    assert c.decode(c.encode(u)) == u and c.encode(u) == c.encode(u)
    print("OK")
