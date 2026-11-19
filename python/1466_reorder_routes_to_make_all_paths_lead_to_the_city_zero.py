# 1466. Reorder Routes to Make All Paths Lead to the City Zero (Medium)
# https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
# Run: python 1466_reorder_routes_to_make_all_paths_lead_to_the_city_zero.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minReorder(self, n: int, connections: List[List[int]]) -> int:
        g = defaultdict(list)
        for a, b in connections:
            g[a].append((b, 1))     # original direction, points away from 0 -> needs flip
            g[b].append((a, 0))
        seen, st, flips = {0}, [0], 0
        while st:
            u = st.pop()
            for v, cost in g[u]:
                if v not in seen:
                    seen.add(v)
                    flips += cost
                    st.append(v)
        return flips


if __name__ == "__main__":
    assert Solution().minReorder(6, [[0, 1], [1, 3], [2, 3], [4, 0], [4, 5]]) == 3
    assert Solution().minReorder(5, [[1, 0], [1, 2], [3, 2], [3, 4]]) == 2
    print("OK")
