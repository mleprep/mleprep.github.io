# 82. Remove Duplicates from Sorted List II (Medium)
# https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
# Run: python 0082_remove_duplicates_from_sorted_list_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = prev = ListNode(0, head)
        while prev.next:
            cur = prev.next
            if cur.next and cur.next.val == cur.val:
                while cur.next and cur.next.val == cur.val:
                    cur = cur.next
                prev.next = cur.next          # drop the whole run
            else:
                prev = cur
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().deleteDuplicates(make_list([1, 2, 3, 3, 4, 4, 5]))) == [1, 2, 5]
    assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 1, 2, 3]))) == [2, 3]
    print("OK")
