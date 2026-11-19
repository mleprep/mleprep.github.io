# 21. Merge Two Sorted Lists (Easy)
# https://leetcode.com/problems/merge-two-sorted-lists/
# Run: python 0021_merge_two_sorted_lists.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = tail = ListNode()
        while list1 and list2:
            if list1.val <= list2.val:
                tail.next, list1 = list1, list1.next
            else:
                tail.next, list2 = list2, list2.next
            tail = tail.next
        tail.next = list1 or list2
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().mergeTwoLists(make_list([1, 2, 4]), make_list([1, 3, 4]))) == [1, 1, 2, 3, 4, 4]
    assert list_vals(Solution().mergeTwoLists(make_list([]), make_list([]))) == []
    assert list_vals(Solution().mergeTwoLists(make_list([]), make_list([0]))) == [0]
    assert list_vals(Solution().mergeTwoLists(make_list([5]), make_list([1, 2, 4]))) == [1, 2, 4, 5]
    print("OK")
