# 733. Flood Fill (Easy)
# https://leetcode.com/problems/flood-fill/
# Run: python 0733_flood_fill.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        old = image[sr][sc]
        if old == color:
            return image
        R, C = len(image), len(image[0])
        image[sr][sc] = color
        stack = [(sr, sc)]
        while stack:
            r, c = stack.pop()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and image[nr][nc] == old:
                    image[nr][nc] = color
                    stack.append((nr, nc))
        return image


if __name__ == "__main__":
    assert Solution().floodFill([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2) == [[2, 2, 2], [2, 2, 0], [2, 0, 1]]
    assert Solution().floodFill([[0, 0, 0], [0, 0, 0]], 0, 0, 0) == [[0, 0, 0], [0, 0, 0]]
    assert Solution().floodFill([[0, 0, 0], [0, 1, 1]], 1, 1, 1) == [[0, 0, 0], [0, 1, 1]]
    assert Solution().floodFill([[1]], 0, 0, 3) == [[3]]
    print("OK")
