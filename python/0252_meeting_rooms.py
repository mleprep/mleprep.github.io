# 252. Meeting Rooms (Easy)
# https://leetcode.com/problems/meeting-rooms/
# Run: python 0252_meeting_rooms.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def canAttendMeetings(self, intervals: List[List[int]]) -> bool:
        intervals.sort()
        return all(intervals[i][0] >= intervals[i - 1][1] for i in range(1, len(intervals)))


if __name__ == "__main__":
    assert Solution().canAttendMeetings([[0, 30], [5, 10], [15, 20]]) is False
    assert Solution().canAttendMeetings([[7, 10], [2, 4]]) is True
    print("OK")
