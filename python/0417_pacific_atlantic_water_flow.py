# 417. Pacific Atlantic Water Flow (Medium)
# https://leetcode.com/problems/pacific-atlantic-water-flow/
# Run: python 0417_pacific_atlantic_water_flow.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        R, C = len(heights), len(heights[0])

        def reach(starts):
            seen, st = set(starts), list(starts)
            while st:
                r, c = st.pop()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if (0 <= nr < R and 0 <= nc < C and (nr, nc) not in seen
                            and heights[nr][nc] >= heights[r][c]):
                        seen.add((nr, nc))
                        st.append((nr, nc))
            return seen

        pac = reach([(0, c) for c in range(C)] + [(r, 0) for r in range(R)])
        atl = reach([(R - 1, c) for c in range(C)] + [(r, C - 1) for r in range(R)])
        return sorted([r, c] for r, c in pac & atl)


if __name__ == "__main__":
    H = [[1, 2, 2, 3, 5], [3, 2, 3, 4, 4], [2, 4, 5, 3, 1], [6, 7, 1, 4, 5], [5, 1, 1, 2, 4]]
    assert Solution().pacificAtlantic(H) == [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]
    print("OK")
