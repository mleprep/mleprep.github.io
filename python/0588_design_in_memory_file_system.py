# 588. Design In-Memory File System (Hard)
# https://leetcode.com/problems/design-in-memory-file-system/
# Run: python 0588_design_in_memory_file_system.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class FileSystem:
    def __init__(self):
        self.root = {"kids": {}, "file": None}       # file: None for a directory, else its content

    def _walk(self, path):
        node = self.root
        for part in filter(None, path.split("/")):
            node = node["kids"].setdefault(part, {"kids": {}, "file": None})
        return node

    def ls(self, path: str) -> List[str]:
        node = self._walk(path)
        return [path.rsplit("/", 1)[1]] if node["file"] is not None else sorted(node["kids"])

    def mkdir(self, path: str) -> None:
        self._walk(path)

    def addContentToFile(self, filePath: str, content: str) -> None:
        node = self._walk(filePath)
        node["file"] = (node["file"] or "") + content

    def readContentFromFile(self, filePath: str) -> str:
        return self._walk(filePath)["file"]


if __name__ == "__main__":
    fs = FileSystem()
    assert fs.ls("/") == []
    fs.mkdir("/a/b/c")
    fs.addContentToFile("/a/b/c/d", "hello")
    assert fs.ls("/") == ["a"]
    assert fs.readContentFromFile("/a/b/c/d") == "hello"
    fs.addContentToFile("/a/b/c/d", " world")
    assert fs.readContentFromFile("/a/b/c/d") == "hello world"
    assert fs.ls("/a/b/c/d") == ["d"]
    fs.mkdir("/a/b/e")
    fs.addContentToFile("/a/b/aa", "x")
    assert fs.ls("/a/b") == ["aa", "c", "e"]
    print("OK")
