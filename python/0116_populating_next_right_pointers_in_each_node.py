# 116. Populating Next Right Pointers in Each Node (Medium)
# https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
# Run: python 0116_populating_next_right_pointers_in_each_node.py
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
    def connect(self, root: "Optional[Node]") -> "Optional[Node]":
        leftmost = root
        while leftmost and leftmost.left:
            node = leftmost
            while node:
                node.left.next = node.right
                if node.next:
                    node.right.next = node.next.left
                node = node.next
            leftmost = leftmost.left
        return root


if __name__ == "__main__":
    n = [Node(i) for i in range(8)]  # 1..7 used, heap layout
    for i in range(1, 4):
        n[i].left, n[i].right = n[2 * i], n[2 * i + 1]
    Solution().connect(n[1])
    assert n[4].next is n[5] and n[5].next is n[6] and n[6].next is n[7] and n[7].next is None and n[2].next is n[3]
    print("OK")
