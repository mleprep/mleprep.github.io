# 49. Group Anagrams (Medium)
# https://leetcode.com/problems/group-anagrams/
# Run: python 0049_group_anagrams.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)
        for w in strs:
            key = [0] * 26
            for ch in w:
                key[ord(ch) - 97] += 1
            groups[tuple(key)].append(w)
        return list(groups.values())


if __name__ == "__main__":
    res = Solution().groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    assert sorted(sorted(g) for g in res) == [["ate", "eat", "tea"], ["bat"], ["nat", "tan"]]
    print("OK")
