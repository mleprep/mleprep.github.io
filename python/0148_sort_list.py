# 148. Sort List (Medium)
# https://leetcode.com/problems/sort-list/
# Run: python 0148_sort_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head
        slow, fast = head, head.next
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        mid, slow.next = slow.next, None
        a, b = self.sortList(head), self.sortList(mid)
        dummy = tail = ListNode()
        while a and b:
            if a.val <= b.val:
                tail.next, a = a, a.next
            else:
                tail.next, b = b, b.next
            tail = tail.next
        tail.next = a or b
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().sortList(make_list([4, 2, 1, 3]))) == [1, 2, 3, 4]
    assert list_vals(Solution().sortList(make_list([-1, 5, 3, 4, 0]))) == [-1, 0, 3, 4, 5]
    print("OK")
