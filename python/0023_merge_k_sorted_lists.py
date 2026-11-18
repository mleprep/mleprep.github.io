# 23. Merge k Sorted Lists (Hard)
# https://leetcode.com/problems/merge-k-sorted-lists/
# Run: python 0023_merge_k_sorted_lists.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = [(node.val, i, node) for i, node in enumerate(lists) if node]
        heapify(heap)
        dummy = tail = ListNode()
        while heap:
            _, i, node = heappop(heap)      # i breaks ties (nodes aren't comparable)
            tail.next = tail = node
            if node.next:
                heappush(heap, (node.next.val, i, node.next))
        return dummy.next


if __name__ == "__main__":
    ls = [make_list([1, 4, 5]), make_list([1, 3, 4]), make_list([2, 6])]
    assert list_vals(Solution().mergeKLists(ls)) == [1, 1, 2, 3, 4, 4, 5, 6]
    assert Solution().mergeKLists([]) is None
    print("OK")
