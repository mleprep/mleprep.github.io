# 2807. Insert Greatest Common Divisors in Linked List (Medium)
# https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
# Run: python 2807_insert_greatest_common_divisors_in_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def insertGreatestCommonDivisors(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        while cur and cur.next:
            cur.next = ListNode(math.gcd(cur.val, cur.next.val), cur.next)
            cur = cur.next.next
        return head


if __name__ == "__main__":
    assert list_vals(Solution().insertGreatestCommonDivisors(make_list([18, 6, 10, 3]))) == [18, 6, 6, 2, 10, 1, 3]
    assert list_vals(Solution().insertGreatestCommonDivisors(make_list([7]))) == [7]
    assert list_vals(Solution().insertGreatestCommonDivisors(make_list([4, 4]))) == [4, 4, 4]
    print("OK")
