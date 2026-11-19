# 68. Text Justification (Hard)
# https://leetcode.com/problems/text-justification/
# Run: python 0068_text_justification.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        res, line, letters = [], [], 0
        for w in words:
            if letters + len(w) + len(line) > maxWidth:
                gaps = len(line) - 1 or 1
                for i in range(maxWidth - letters):
                    line[i % gaps] += " "
                res.append("".join(line))
                line, letters = [], 0
            line.append(w)
            letters += len(w)
        res.append(" ".join(line).ljust(maxWidth))
        return res


if __name__ == "__main__":
    assert Solution().fullJustify(["This", "is", "an", "example", "of", "text", "justification."], 16) == [
        "This    is    an", "example  of text", "justification.  "]
    assert Solution().fullJustify(["What", "must", "be", "acknowledgment", "shall", "be"], 16) == [
        "What   must   be", "acknowledgment  ", "shall be        "]
    assert Solution().fullJustify(["Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
                                   "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"], 20) == [
        "Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
        "everything  else  we", "do                  "]
    assert Solution().fullJustify(["a"], 1) == ["a"]
    print("OK")
