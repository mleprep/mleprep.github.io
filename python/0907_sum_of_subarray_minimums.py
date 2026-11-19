# 907. Sum of Subarray Minimums (Medium)
# https://leetcode.com/problems/sum-of-subarray-minimums/
# Run: python 0907_sum_of_subarray_minimums.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        MOD, n = 10 ** 9 + 7, len(arr)
        left, right, st = [0] * n, [0] * n, []
        for i in range(n):
            while st and arr[st[-1]] > arr[i]:
                st.pop()
            left[i] = i - (st[-1] if st else -1)
            st.append(i)
        st = []
        for i in range(n - 1, -1, -1):
            while st and arr[st[-1]] >= arr[i]:
                st.pop()
            right[i] = (st[-1] if st else n) - i
            st.append(i)
        return sum(a * l * r for a, l, r in zip(arr, left, right)) % MOD


if __name__ == "__main__":
    assert Solution().sumSubarrayMins([3, 1, 2, 4]) == 17
    assert Solution().sumSubarrayMins([11, 81, 94, 43, 3]) == 444
    print("OK")
