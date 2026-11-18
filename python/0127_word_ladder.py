# 127. Word Ladder (Hard)
# https://leetcode.com/problems/word-ladder/
# Run: python 0127_word_ladder.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        words = set(wordList)
        if endWord not in words:
            return 0
        q, steps = deque([beginWord]), 1
        words.discard(beginWord)
        while q:
            for _ in range(len(q)):
                w = q.popleft()
                if w == endWord:
                    return steps
                for i in range(len(w)):
                    for ch in string.ascii_lowercase:
                        nxt = w[:i] + ch + w[i + 1:]
                        if nxt in words:
                            words.remove(nxt)
                            q.append(nxt)
            steps += 1
        return 0


if __name__ == "__main__":
    assert Solution().ladderLength("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"]) == 5
    assert Solution().ladderLength("hit", "cog", ["hot", "dot", "dog", "lot", "log"]) == 0
    print("OK")
