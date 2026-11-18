# 981. Time Based Key-Value Store (Medium)
# https://leetcode.com/problems/time-based-key-value-store/
# Run: python 0981_time_based_key_value_store.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class TimeMap:
    def __init__(self):
        self.data = defaultdict(list)  # key -> [(ts, value)]

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.data[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        arr = self.data.get(key, [])
        i = bisect.bisect_right(arr, (timestamp, chr(0x10FFFF)))  # first ts > timestamp
        return arr[i - 1][1] if i else ""


if __name__ == "__main__":
    tm = TimeMap()
    tm.set("foo", "bar", 1)
    assert tm.get("foo", 1) == "bar" and tm.get("foo", 3) == "bar"
    tm.set("foo", "bar2", 4)
    assert tm.get("foo", 4) == "bar2" and tm.get("foo", 5) == "bar2" and tm.get("foo", 0) == ""
    print("OK")
