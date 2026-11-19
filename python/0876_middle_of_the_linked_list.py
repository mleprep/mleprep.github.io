# 876. Middle of the Linked List (Easy)
# https://leetcode.com/problems/middle-of-the-linked-list/
# Run: python 0876_middle_of_the_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        return slow


if __name__ == "__main__":
    assert Solution().middleNode(make_list([1, 2, 3, 4, 5])).val == 3
    assert Solution().middleNode(make_list([1, 2, 3, 4, 5, 6])).val == 4
    print("OK")
