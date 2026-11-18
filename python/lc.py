"""LeetCode's ListNode / TreeNode plus tiny helpers so every solution file runs locally."""
from collections import deque


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __repr__(self):
        return f"ListNode({list_vals(self)})"


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self):
        return f"TreeNode({tree_vals(self)})"


def make_list(vals):
    """[1, 2, 3] -> 1 -> 2 -> 3"""
    dummy = tail = ListNode()
    for v in vals:
        tail.next = ListNode(v)
        tail = tail.next
    return dummy.next


def list_vals(head, limit=10_000):
    out = []
    while head and len(out) < limit:
        out.append(head.val)
        head = head.next
    return out


def make_tree(vals):
    """LeetCode level-order list, e.g. [3, 9, 20, None, None, 15, 7]."""
    if not vals or vals[0] is None:
        return None
    root = TreeNode(vals[0])
    q, i = deque([root]), 1
    while q and i < len(vals):
        node = q.popleft()
        if i < len(vals) and vals[i] is not None:
            node.left = TreeNode(vals[i])
            q.append(node.left)
        i += 1
        if i < len(vals) and vals[i] is not None:
            node.right = TreeNode(vals[i])
            q.append(node.right)
        i += 1
    return root


def tree_vals(root):
    """Inverse of make_tree (level order, trailing Nones trimmed)."""
    out, q = [], deque([root])
    while q:
        n = q.popleft()
        if n is None:
            out.append(None)
            continue
        out.append(n.val)
        q.append(n.left)
        q.append(n.right)
    while out and out[-1] is None:
        out.pop()
    return out


def find_node(root, val):
    if root is None or root.val == val:
        return root
    return find_node(root.left, val) or find_node(root.right, val)
