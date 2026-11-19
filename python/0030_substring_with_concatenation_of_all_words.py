# 30. Substring with Concatenation of All Words (Hard)
# https://leetcode.com/problems/substring-with-concatenation-of-all-words/
# Run: python 0030_substring_with_concatenation_of_all_words.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        L, k = len(words[0]), len(words)
        need = Counter(words)
        res = []
        for off in range(L):
            l, cnt, used = off, defaultdict(int), 0
            for r in range(off, len(s) - L + 1, L):
                w = s[r:r + L]
                if w not in need:
                    l, cnt, used = r + L, defaultdict(int), 0
                    continue
                cnt[w] += 1
                used += 1
                while cnt[w] > need[w]:
                    cnt[s[l:l + L]] -= 1
                    used -= 1
                    l += L
                if used == k:
                    res.append(l)
        return sorted(res)


if __name__ == "__main__":
    assert Solution().findSubstring("barfoothefoobarman", ["foo", "bar"]) == [0, 9]
    assert Solution().findSubstring("wordgoodgoodgoodbestword", ["word", "good", "best", "word"]) == []
    assert Solution().findSubstring("barfoofoobarthefoobarman", ["bar", "foo", "the"]) == [6, 9, 12]
    print("OK")
