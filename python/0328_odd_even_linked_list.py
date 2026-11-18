# 328. Odd Even Linked List (Medium)
# https://leetcode.com/problems/odd-even-linked-list/
# Run: python 0328_odd_even_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head
        odd, even = head, head.next
        even_head = even
        while even and even.next:
            odd.next = even.next
            odd = odd.next
            even.next = odd.next
            even = even.next
        odd.next = even_head
        return head


if __name__ == "__main__":
    assert list_vals(Solution().oddEvenList(make_list([1, 2, 3, 4, 5]))) == [1, 3, 5, 2, 4]
    assert list_vals(Solution().oddEvenList(make_list([2, 1, 3, 5, 6, 4, 7]))) == [2, 3, 6, 7, 1, 5, 4]
    print("OK")
