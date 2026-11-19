# 735. Asteroid Collision (Medium)
# https://leetcode.com/problems/asteroid-collision/
# Run: python 0735_asteroid_collision.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        st = []
        for x in asteroids:
            alive = True
            while alive and x < 0 and st and st[-1] > 0:
                if st[-1] < -x:
                    st.pop()
                    continue
                if st[-1] == -x:
                    st.pop()
                alive = False
            if alive:
                st.append(x)
        return st


if __name__ == "__main__":
    assert Solution().asteroidCollision([5, 10, -5]) == [5, 10]
    assert Solution().asteroidCollision([8, -8]) == []
    assert Solution().asteroidCollision([10, 2, -5]) == [10]
    assert Solution().asteroidCollision([-2, -1, 1, 2]) == [-2, -1, 1, 2]
    print("OK")
