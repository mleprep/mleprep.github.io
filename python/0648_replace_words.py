# 648. Replace Words (Medium)
# https://leetcode.com/problems/replace-words/
# Run: python 0648_replace_words.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        root = {}
        for w in dictionary:
            node = root
            for ch in w:
                node = node.setdefault(ch, {})
            node["$"] = True

        def shortest(word):
            node = root
            for i, ch in enumerate(word):
                if ch not in node:
                    return word
                node = node[ch]
                if "$" in node:
                    return word[:i + 1]
            return word

        return " ".join(shortest(w) for w in sentence.split())


if __name__ == "__main__":
    assert Solution().replaceWords(["cat", "bat", "rat"], "the cattle was rattled by the battery") == "the cat was rat by the bat"
    assert Solution().replaceWords(["a", "b", "c"], "aadsfasf absbs bbab cadsfafs") == "a a b c"
    print("OK")
