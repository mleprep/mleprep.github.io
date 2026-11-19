# 1472. Design Browser History (Medium)
# https://leetcode.com/problems/design-browser-history/
# Run: python 1472_design_browser_history.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class BrowserHistory:
    def __init__(self, homepage: str):
        self.pages, self.cur, self.last = [homepage], 0, 0

    def visit(self, url: str) -> None:
        self.cur += 1
        if self.cur == len(self.pages):
            self.pages.append(url)
        else:
            self.pages[self.cur] = url
        self.last = self.cur

    def back(self, steps: int) -> str:
        self.cur = max(0, self.cur - steps)
        return self.pages[self.cur]

    def forward(self, steps: int) -> str:
        self.cur = min(self.last, self.cur + steps)
        return self.pages[self.cur]


if __name__ == "__main__":
    b = BrowserHistory("leetcode.com")
    for u in ("google.com", "facebook.com", "youtube.com"):
        b.visit(u)
    assert b.back(1) == "facebook.com" and b.back(1) == "google.com" and b.forward(1) == "facebook.com"
    b.visit("linkedin.com")
    assert b.forward(2) == "linkedin.com" and b.back(2) == "google.com" and b.back(7) == "leetcode.com"
    print("OK")
