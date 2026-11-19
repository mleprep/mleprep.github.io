# 433. Minimum Genetic Mutation (Medium)
# https://leetcode.com/problems/minimum-genetic-mutation/
# Run: python 0433_minimum_genetic_mutation.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minMutation(self, startGene: str, endGene: str, bank: List[str]) -> int:
        valid = set(bank)
        q, seen, steps = deque([startGene]), {startGene}, 0
        while q:
            for _ in range(len(q)):
                g = q.popleft()
                if g == endGene:
                    return steps
                for i in range(len(g)):
                    for ch in "ACGT":
                        nxt = g[:i] + ch + g[i + 1:]
                        if nxt in valid and nxt not in seen:
                            seen.add(nxt)
                            q.append(nxt)
            steps += 1
        return -1


if __name__ == "__main__":
    assert Solution().minMutation("AACCGGTT", "AACCGGTA", ["AACCGGTA"]) == 1
    assert Solution().minMutation("AACCGGTT", "AAACGGTA", ["AACCGGTA", "AACCGCTA", "AAACGGTA"]) == 2
    assert Solution().minMutation("AAAAACCC", "AACCCCCC", ["AAAACCCC", "AAACCCCC", "AACCCCCC"]) == 3
    print("OK")
