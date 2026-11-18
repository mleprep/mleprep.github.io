# 863. All Nodes Distance K in Binary Tree (Medium)
# https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
# Run: python 0863_all_nodes_distance_k_in_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        parent = {root: None}
        st = [root]
        while st:
            node = st.pop()
            for c in (node.left, node.right):
                if c:
                    parent[c] = node
                    st.append(c)
        q, seen = deque([target]), {target}
        for _ in range(k):
            for _ in range(len(q)):
                node = q.popleft()
                for nb in (node.left, node.right, parent[node]):
                    if nb and nb not in seen:
                        seen.add(nb)
                        q.append(nb)
        return [n.val for n in q]


if __name__ == "__main__":
    t = make_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
    assert sorted(Solution().distanceK(t, find_node(t, 5), 2)) == [1, 4, 7]
    print("OK")
