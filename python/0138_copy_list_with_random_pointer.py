# 138. Copy List with Random Pointer (Medium)
# https://leetcode.com/problems/copy-list-with-random-pointer/
# Run: python 0138_copy_list_with_random_pointer.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Node:  # provided by LeetCode
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val, self.next, self.random = int(x), next, random

class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        clone = {None: None}
        cur = head
        while cur:
            clone[cur] = Node(cur.val)
            cur = cur.next
        cur = head
        while cur:
            clone[cur].next = clone[cur.next]
            clone[cur].random = clone[cur.random]
            cur = cur.next
        return clone[head]


if __name__ == "__main__":
    spec = [[7, None], [13, 0], [11, 4], [10, 2], [1, 0]]
    nodes = [Node(v) for v, _ in spec]
    for i, (_, r) in enumerate(spec):
        nodes[i].next = nodes[i + 1] if i + 1 < len(nodes) else None
        nodes[i].random = nodes[r] if r is not None else None
    c, out = Solution().copyRandomList(nodes[0]), []
    copies = []
    while c:
        assert c not in nodes
        copies.append(c); c = c.next
    assert [[n.val, copies.index(n.random) if n.random else None] for n in copies] == spec
    print("OK")
