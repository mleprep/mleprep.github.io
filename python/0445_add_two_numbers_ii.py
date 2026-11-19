# 445. Add Two Numbers II (Medium)
# https://leetcode.com/problems/add-two-numbers-ii/
# Run: python 0445_add_two_numbers_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        a, b = [], []
        while l1:
            a.append(l1.val); l1 = l1.next
        while l2:
            b.append(l2.val); l2 = l2.next
        head, carry = None, 0
        while a or b or carry:
            s = carry + (a.pop() if a else 0) + (b.pop() if b else 0)
            carry, d = divmod(s, 10)
            head = ListNode(d, head)
        return head


if __name__ == "__main__":
    assert list_vals(Solution().addTwoNumbers(make_list([7, 2, 4, 3]), make_list([5, 6, 4]))) == [7, 8, 0, 7]
    assert list_vals(Solution().addTwoNumbers(make_list([5]), make_list([5]))) == [1, 0]
    print("OK")
