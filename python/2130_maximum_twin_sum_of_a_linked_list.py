# 2130. Maximum Twin Sum of a Linked List (Medium)
# https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
# Run: python 2130_maximum_twin_sum_of_a_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        prev = None                                  # reverse the second half
        while slow:
            slow.next, prev, slow = prev, slow, slow.next
        best = 0
        while prev:
            best = max(best, head.val + prev.val)
            head, prev = head.next, prev.next
        return best


if __name__ == "__main__":
    assert Solution().pairSum(make_list([5, 4, 2, 1])) == 6
    assert Solution().pairSum(make_list([4, 2, 2, 3])) == 7
    assert Solution().pairSum(make_list([1, 100000])) == 100001
    print("OK")
