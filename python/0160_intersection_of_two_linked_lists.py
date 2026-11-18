# 160. Intersection of Two Linked Lists (Easy)
# https://leetcode.com/problems/intersection-of-two-linked-lists/
# Run: python 0160_intersection_of_two_linked_lists.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        a, b = headA, headB
        while a is not b:
            a = a.next if a else headB
            b = b.next if b else headA
        return a


if __name__ == "__main__":
    common = make_list([8, 4, 5])
    A = make_list([4, 1]); A.next.next = common
    B = make_list([5, 6, 1]); B.next.next.next = common
    assert Solution().getIntersectionNode(A, B) is common
    assert Solution().getIntersectionNode(make_list([1]), make_list([2])) is None
    print("OK")
