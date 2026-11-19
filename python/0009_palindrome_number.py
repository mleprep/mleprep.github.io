# 9. Palindrome Number (Easy)
# https://leetcode.com/problems/palindrome-number/
# Run: python 0009_palindrome_number.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0 or (x % 10 == 0 and x != 0):
            return False
        rev = 0
        while x > rev:
            x, d = divmod(x, 10)
            rev = rev * 10 + d
        return x == rev or x == rev // 10


if __name__ == "__main__":
    assert Solution().isPalindrome(121) is True
    assert Solution().isPalindrome(-121) is False
    assert Solution().isPalindrome(10) is False
    assert Solution().isPalindrome(1221) is True
    print("OK")
