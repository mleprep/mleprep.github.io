# 841. Keys and Rooms (Medium)
# https://leetcode.com/problems/keys-and-rooms/
# Run: python 0841_keys_and_rooms.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        seen, st = {0}, [0]
        while st:
            for k in rooms[st.pop()]:
                if k not in seen:
                    seen.add(k)
                    st.append(k)
        return len(seen) == len(rooms)


if __name__ == "__main__":
    assert Solution().canVisitAllRooms([[1], [2], [3], []]) is True
    assert Solution().canVisitAllRooms([[1, 3], [3, 0, 1], [2], [0]]) is False
    print("OK")
