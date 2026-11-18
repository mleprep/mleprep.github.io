# 79. Word Search (Medium)
# https://leetcode.com/problems/word-search/
# Run: python 0079_word_search.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        R, C = len(board), len(board[0])

        def dfs(r, c, i):
            if i == len(word):
                return True
            if not (0 <= r < R and 0 <= c < C) or board[r][c] != word[i]:
                return False
            board[r][c] = "#"
            found = dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or dfs(r, c + 1, i + 1) or dfs(r, c - 1, i + 1)
            board[r][c] = word[i]
            return found

        return any(dfs(r, c, 0) for r in range(R) for c in range(C))


if __name__ == "__main__":
    B = [list("ABCE"), list("SFCS"), list("ADEE")]
    assert Solution().exist(B, "ABCCED") is True
    assert Solution().exist(B, "SEE") is True
    assert Solution().exist(B, "ABCB") is False
    print("OK")
