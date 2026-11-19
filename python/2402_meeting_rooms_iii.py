# 2402. Meeting Rooms III (Hard)
# https://leetcode.com/problems/meeting-rooms-iii/
# Run: python 2402_meeting_rooms_iii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        free, busy, count = list(range(n)), [], [0] * n
        for s, e in sorted(meetings):
            while busy and busy[0][0] <= s:
                heappush(free, heappop(busy)[1])
            if free:
                room = heappop(free)
                heappush(busy, (e, room))
            else:
                end, room = heappop(busy)
                heappush(busy, (end + e - s, room))
            count[room] += 1
        return count.index(max(count))


if __name__ == "__main__":
    assert Solution().mostBooked(2, [[0, 10], [1, 5], [2, 7], [3, 4]]) == 0
    assert Solution().mostBooked(3, [[1, 20], [2, 10], [3, 5], [4, 9], [6, 8]]) == 1
    assert Solution().mostBooked(1, [[0, 1]]) == 0
    assert Solution().mostBooked(4, [[18, 19], [3, 12], [17, 19], [2, 13], [7, 10]]) == 0
    print("OK")
