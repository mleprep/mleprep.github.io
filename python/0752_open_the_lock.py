# 752. Open the Lock (Medium)
# https://leetcode.com/problems/open-the-lock/
# Run: python 0752_open_the_lock.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        dead = set(deadends)
        if "0000" in dead:
            return -1
        q, seen, steps = deque(["0000"]), {"0000"}, 0
        while q:
            for _ in range(len(q)):
                cur = q.popleft()
                if cur == target:
                    return steps
                for i in range(4):
                    d = int(cur[i])
                    for nd in ((d + 1) % 10, (d - 1) % 10):
                        nxt = cur[:i] + str(nd) + cur[i + 1:]
                        if nxt not in seen and nxt not in dead:
                            seen.add(nxt)
                            q.append(nxt)
            steps += 1
        return -1


if __name__ == "__main__":
    assert Solution().openLock(["0201", "0101", "0102", "1212", "2002"], "0202") == 6
    assert Solution().openLock(["8888"], "0009") == 1
    assert Solution().openLock(["8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"], "8888") == -1
    print("OK")
