# 2095. Delete the Middle Node of a Linked List (Medium)
# https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
# Run: python 2095_delete_the_middle_node_of_a_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head.next:
            return None
        slow, fast = head, head.next.next
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        slow.next = slow.next.next
        return head


if __name__ == "__main__":
    assert list_vals(Solution().deleteMiddle(make_list([1, 3, 4, 7, 1, 2, 6]))) == [1, 3, 4, 1, 2, 6]
    assert list_vals(Solution().deleteMiddle(make_list([1, 2, 3, 4]))) == [1, 2, 4]
    assert list_vals(Solution().deleteMiddle(make_list([2, 1]))) == [2]
    assert list_vals(Solution().deleteMiddle(make_list([1]))) == []
    print("OK")
