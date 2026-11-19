# 912. Sort an Array (Medium)
# https://leetcode.com/problems/sort-an-array/
# Run: python 0912_sort_an_array.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        if len(nums) <= 1:
            return nums
        mid = len(nums) // 2
        left, right = self.sortArray(nums[:mid]), self.sortArray(nums[mid:])
        out, i, j = [], 0, 0
        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                out.append(left[i])
                i += 1
            else:
                out.append(right[j])
                j += 1
        return out + left[i:] + right[j:]


if __name__ == "__main__":
    assert Solution().sortArray([5, 2, 3, 1]) == [1, 2, 3, 5]
    assert Solution().sortArray([5, 1, 1, 2, 0, 0]) == [0, 0, 1, 1, 2, 5]
    assert Solution().sortArray([1]) == [1]
    rnd = random.Random(7)
    data = [rnd.randint(-50, 50) for _ in range(300)]
    assert Solution().sortArray(data[:]) == sorted(data)
    print("OK")
