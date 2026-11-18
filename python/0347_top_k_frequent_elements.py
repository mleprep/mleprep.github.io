# 347. Top K Frequent Elements (Medium)
# https://leetcode.com/problems/top-k-frequent-elements/
# Run: python 0347_top_k_frequent_elements.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        cnt = Counter(nums)
        bucket = [[] for _ in range(len(nums) + 1)]
        for x, f in cnt.items():
            bucket[f].append(x)
        res = []
        for f in range(len(bucket) - 1, 0, -1):
            for x in bucket[f]:
                res.append(x)
                if len(res) == k:
                    return res
        return res


if __name__ == "__main__":
    assert sorted(Solution().topKFrequent([1, 1, 1, 2, 2, 3], 2)) == [1, 2]
    assert Solution().topKFrequent([1], 1) == [1]
    print("OK")
