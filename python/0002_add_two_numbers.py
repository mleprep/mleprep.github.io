# 2. Add Two Numbers (Medium)
# https://leetcode.com/problems/add-two-numbers/
# Run: python 0002_add_two_numbers.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = tail = ListNode()
        carry = 0
        while l1 or l2 or carry:
            s = carry + (l1.val if l1 else 0) + (l2.val if l2 else 0)
            carry, d = divmod(s, 10)
            tail.next = ListNode(d)
            tail = tail.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().addTwoNumbers(make_list([2, 4, 3]), make_list([5, 6, 4]))) == [7, 0, 8]
    assert list_vals(Solution().addTwoNumbers(make_list([9, 9, 9]), make_list([1]))) == [0, 0, 0, 1]
    print("OK")
