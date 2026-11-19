# 2392. Build a Matrix With Conditions (Hard)
# https://leetcode.com/problems/build-a-matrix-with-conditions/
# Run: python 2392_build_a_matrix_with_conditions.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def buildMatrix(self, k: int, rowConditions: List[List[int]], colConditions: List[List[int]]) -> List[List[int]]:
        def topo(conds):
            adj, indeg = defaultdict(list), [0] * (k + 1)
            for a, b in conds:
                adj[a].append(b)
                indeg[b] += 1
            q = deque(v for v in range(1, k + 1) if indeg[v] == 0)
            order = []
            while q:
                v = q.popleft()
                order.append(v)
                for w in adj[v]:
                    indeg[w] -= 1
                    if indeg[w] == 0:
                        q.append(w)
            return {v: i for i, v in enumerate(order)} if len(order) == k else None

        row, col = topo(rowConditions), topo(colConditions)
        if row is None or col is None:
            return []
        m = [[0] * k for _ in range(k)]
        for v in range(1, k + 1):
            m[row[v]][col[v]] = v
        return m


if __name__ == "__main__":
    def valid(k, rc, cc, m):  # any matrix satisfying the conditions is accepted
        pos = {m[i][j]: (i, j) for i in range(k) for j in range(k) if m[i][j]}
        return (len(m) == k and all(len(r) == k for r in m) and sorted(pos) == list(range(1, k + 1))
                and all(pos[a][0] < pos[b][0] for a, b in rc) and all(pos[a][1] < pos[b][1] for a, b in cc))
    assert valid(3, [[1, 2], [3, 2]], [[2, 1], [3, 2]], Solution().buildMatrix(3, [[1, 2], [3, 2]], [[2, 1], [3, 2]]))
    assert Solution().buildMatrix(3, [[1, 2], [2, 3], [3, 1], [2, 3]], [[2, 1]]) == []
    assert valid(4, [[4, 3], [3, 2], [2, 1]], [[1, 2], [2, 3]], Solution().buildMatrix(4, [[4, 3], [3, 2], [2, 1]], [[1, 2], [2, 3]]))
    print("OK")
