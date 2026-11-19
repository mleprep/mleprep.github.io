# 1268. Search Suggestions System (Medium)
# https://leetcode.com/problems/search-suggestions-system/
# Run: python 1268_search_suggestions_system.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def suggestedProducts(self, products: List[str], searchWord: str) -> List[List[str]]:
        products.sort()
        res, prefix = [], ""
        for ch in searchWord:
            prefix += ch
            i = bisect.bisect_left(products, prefix)
            res.append([p for p in products[i:i + 3] if p.startswith(prefix)])
        return res


if __name__ == "__main__":
    P = ["mobile", "mouse", "moneypot", "monitor", "mousepad"]
    assert Solution().suggestedProducts(P, "mouse") == [
        ["mobile", "moneypot", "monitor"], ["mobile", "moneypot", "monitor"],
        ["mouse", "mousepad"], ["mouse", "mousepad"], ["mouse", "mousepad"]]
    assert Solution().suggestedProducts(["havana"], "tatiana") == [[], [], [], [], [], [], []]
    assert Solution().suggestedProducts(["bags", "baggage", "banner", "box", "cloths"], "bags") == [["baggage", "bags", "banner"], ["baggage", "bags", "banner"], ["baggage", "bags"], ["bags"]]
    print("OK")
