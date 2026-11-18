# 297. Serialize and Deserialize Binary Tree (Hard)
# https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
# Run: python 0297_serialize_and_deserialize_binary_tree.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Codec:
    def serialize(self, root: Optional[TreeNode]) -> str:
        out = []

        def dfs(node):
            if not node:
                out.append("#")
                return
            out.append(str(node.val))
            dfs(node.left)
            dfs(node.right)

        dfs(root)
        return ",".join(out)

    def deserialize(self, data: str) -> Optional[TreeNode]:
        it = iter(data.split(","))

        def build():
            v = next(it)
            if v == "#":
                return None
            node = TreeNode(int(v))
            node.left = build()
            node.right = build()
            return node

        return build()


if __name__ == "__main__":
    c = Codec()
    t = make_tree([1, 2, 3, None, None, 4, 5])
    assert c.serialize(t) == "1,2,#,#,3,4,#,#,5,#,#"
    assert tree_vals(c.deserialize(c.serialize(t))) == [1, 2, 3, None, None, 4, 5]
    assert c.deserialize(c.serialize(None)) is None
    print("OK")
