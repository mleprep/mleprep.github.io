# 93. Restore IP Addresses (Medium)
# https://leetcode.com/problems/restore-ip-addresses/
# Run: python 0093_restore_ip_addresses.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []

        def bt(i, parts):
            if len(parts) == 4:
                if i == len(s):
                    res.append(".".join(parts))
                return
            rem = 4 - len(parts)
            if not rem <= len(s) - i <= 3 * rem:
                return
            for L in range(1, 4):
                p = s[i:i + L]
                if len(p) < L or (p[0] == "0" and L > 1) or int(p) > 255:
                    break
                bt(i + L, parts + [p])

        bt(0, [])
        return res


if __name__ == "__main__":
    assert sorted(Solution().restoreIpAddresses("25525511135")) == sorted(["255.255.11.135", "255.255.111.35"])
    assert sorted(Solution().restoreIpAddresses("0000")) == sorted(["0.0.0.0"])
    assert sorted(Solution().restoreIpAddresses("101023")) == sorted(["1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"])
    print("OK")
