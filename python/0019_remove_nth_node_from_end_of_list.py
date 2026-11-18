# 19. Remove Nth Node From End of List (Medium)
# https://leetcode.com/problems/remove-nth-node-from-end-of-list/
# Run: python 0019_remove_nth_node_from_end_of_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        slow = fast = dummy
        for _ in range(n):
            fast = fast.next
        while fast.next:
            slow, fast = slow.next, fast.next
        slow.next = slow.next.next
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().removeNthFromEnd(make_list([1, 2, 3, 4, 5]), 2)) == [1, 2, 3, 5]
    assert list_vals(Solution().removeNthFromEnd(make_list([1]), 1)) == []
    print("OK")
