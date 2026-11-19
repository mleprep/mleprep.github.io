# 117. Populating Next Right Pointers in Each Node II (Medium)
# https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
# Run: python 0117_populating_next_right_pointers_in_each_node_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Node:  # provided by LeetCode
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val, self.left, self.right, self.next = val, left, right, next

class Solution:
    def connect(self, root: "Node") -> "Node":
        level = root
        while level:
            dummy = tail = Node(0)
            node = level
            while node:
                for child in (node.left, node.right):
                    if child:
                        tail.next = tail = child
                node = node.next
            level = dummy.next
        return root


if __name__ == "__main__":
    n = {i: Node(i) for i in range(1, 8)}
    n[1].left, n[1].right = n[2], n[3]
    n[2].left, n[2].right, n[3].right = n[4], n[5], n[7]
    Solution().connect(n[1])
    assert n[2].next is n[3] and n[4].next is n[5] and n[5].next is n[7] and n[7].next is None
    assert Solution().connect(None) is None
    print("OK")
