# 253. Meeting Rooms II (Medium)
# https://leetcode.com/problems/meeting-rooms-ii/
# Run: python 0253_meeting_rooms_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        starts = sorted(s for s, _ in intervals)
        ends = sorted(e for _, e in intervals)
        rooms = best = j = 0
        for s in starts:
            while ends[j] <= s:
                rooms -= 1
                j += 1
            rooms += 1
            best = max(best, rooms)
        return best


if __name__ == "__main__":
    assert Solution().minMeetingRooms([[0, 30], [5, 10], [15, 20]]) == 2
    assert Solution().minMeetingRooms([[7, 10], [2, 4]]) == 1
    print("OK")
