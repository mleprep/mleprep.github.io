# 427. Construct Quad Tree (Medium)
# https://leetcode.com/problems/construct-quad-tree/
# Run: python 0427_construct_quad_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Node:  # provided by LeetCode
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight


class Solution:
    def construct(self, grid: List[List[int]]) -> 'Node':
        def build(r, c, n):
            if n == 1:
                return Node(grid[r][c] == 1, True, None, None, None, None)
            h = n // 2
            kids = [build(r, c, h), build(r, c + h, h), build(r + h, c, h), build(r + h, c + h, h)]
            if all(k.isLeaf for k in kids) and len({k.val for k in kids}) == 1:
                return Node(kids[0].val, True, None, None, None, None)
            return Node(True, False, *kids)

        return build(0, 0, len(grid))


if __name__ == "__main__":
    def quad(node: "Node"):  # level order: "*" = internal node, "0"/"1" = leaf value
        out, q = [], deque([node])
        while q:
            n = q.popleft()
            out.append(("1" if n.val else "0") if n.isLeaf else "*")
            if not n.isLeaf:
                q.extend([n.topLeft, n.topRight, n.bottomLeft, n.bottomRight])
        return "".join(out)
    assert quad(Solution().construct([[0, 1], [1, 0]])) == "*0110"
    assert quad(Solution().construct([[1]])) == "1"
    assert quad(Solution().construct([[1, 1], [1, 1]])) == "1"
    assert quad(Solution().construct([[1, 1, 0, 0], [1, 1, 0, 0], [1, 1, 1, 1], [1, 1, 1, 1]])) == "*1011"
    assert quad(Solution().construct([[1, 1, 0, 0], [1, 0, 0, 0], [1, 1, 1, 1], [1, 1, 1, 1]])) == "**0111110"
    print("OK")
