# 83. Remove Duplicates from Sorted List (Easy)
# https://leetcode.com/problems/remove-duplicates-from-sorted-list/
# Run: python 0083_remove_duplicates_from_sorted_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        while cur and cur.next:
            if cur.next.val == cur.val:
                cur.next = cur.next.next
            else:
                cur = cur.next
        return head


if __name__ == "__main__":
    assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 2, 3, 3]))) == [1, 2, 3]
    assert list_vals(Solution().deleteDuplicates(make_list([]))) == []
    assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 1]))) == [1]
    print("OK")
