# Trees

**Two traversal families**
- **DFS** (recursion): pre/in/post-order. Ask "what does each subtree need to *return* to its parent?" (height, sum, is-valid, best path…). Post-order = combine children results.
- **BFS** (queue): level-order, shortest depth, right-side view, connect levels.

**BST**: in-order traversal is sorted; every node bounds its subtree `(lo, hi)`.

```python
def dfs(node):                      # post-order template
    if not node: return BASE
    L, R = dfs(node.left), dfs(node.right)
    update_global_answer(node, L, R)
    return combine(node, L, R)

q = deque([root])                   # BFS level-order template
while q:
    for _ in range(len(q)):
        node = q.popleft()
        ...; q.extend(c for c in (node.left, node.right) if c)
```

Local tests build trees from LeetCode's level-order format: Python `make_tree([3, 9, 20, None, None, 15, 7])`, Java `TreeNode.of(3, 9, 20, null, null, 15, 7)`.

## 226. Invert Binary Tree | Easy
related: 101, 100
tags: blind75, dfs

Swap the children of every node (any traversal order works).

```viz
     4              4
   2   7    ->    7   2
  1 3 6 9        9 6 3 1
```

**Complexity:** O(n) time, O(h) stack.

```python
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if root:
            root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root
# test
assert tree_vals(Solution().invertTree(make_tree([4, 2, 7, 1, 3, 6, 9]))) == [4, 7, 2, 9, 6, 3, 1]
```

```java
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if (root == null) return null;
        TreeNode l = invertTree(root.left);
        root.left = invertTree(root.right);
        root.right = l;
        return root;
    }
}
// test
check(new Solution().invertTree(TreeNode.of(4, 2, 7, 1, 3, 6, 9)), TreeNode.of(4, 7, 2, 9, 6, 3, 1));
```

## 104. Maximum Depth of Binary Tree | Easy
related: 110, 543, 111
tags: blind75, dfs

`depth(node) = 1 + max(depth(left), depth(right))`, `depth(None) = 0`.

**Complexity:** O(n).

```python
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))
# test
assert Solution().maxDepth(make_tree([3, 9, 20, None, None, 15, 7])) == 3
assert Solution().maxDepth(None) == 0
```

```java
class Solution {
    public int maxDepth(TreeNode root) {
        return root == null ? 0 : 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
    }
}
// test
check(new Solution().maxDepth(TreeNode.of(3, 9, 20, null, null, 15, 7)), 3);
```

## 543. Diameter of Binary Tree | Easy
related: 124, 104
tags: dfs, global-answer

The longest path *through* a node = `height(left) + height(right)`. Return height upward, update a global best on the way.

**Complexity:** O(n).

```python
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        best = 0

        def height(node):
            nonlocal best
            if not node:
                return 0
            l, r = height(node.left), height(node.right)
            best = max(best, l + r)
            return 1 + max(l, r)

        height(root)
        return best
# test
assert Solution().diameterOfBinaryTree(make_tree([1, 2, 3, 4, 5])) == 3
assert Solution().diameterOfBinaryTree(make_tree([1, 2])) == 1
```

```java
class Solution {
    private int best = 0;

    public int diameterOfBinaryTree(TreeNode root) {
        height(root);
        return best;
    }

    private int height(TreeNode n) {
        if (n == null) return 0;
        int l = height(n.left), r = height(n.right);
        best = Math.max(best, l + r);
        return 1 + Math.max(l, r);
    }
}
// test
check(new Solution().diameterOfBinaryTree(TreeNode.of(1, 2, 3, 4, 5)), 3);
```

## 110. Balanced Binary Tree | Easy
related: 104
tags: dfs

Return height, or `-1` as soon as any subtree is unbalanced (children heights differ by > 1). Single pass, O(n).

**Complexity:** O(n).

```python
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def h(node):
            if not node:
                return 0
            l, r = h(node.left), h(node.right)
            if l < 0 or r < 0 or abs(l - r) > 1:
                return -1
            return 1 + max(l, r)

        return h(root) >= 0
# test
assert Solution().isBalanced(make_tree([3, 9, 20, None, None, 15, 7])) is True
assert Solution().isBalanced(make_tree([1, 2, 2, 3, 3, None, None, 4, 4])) is False
```

```java
class Solution {
    public boolean isBalanced(TreeNode root) { return h(root) >= 0; }

    private int h(TreeNode n) {
        if (n == null) return 0;
        int l = h(n.left), r = h(n.right);
        if (l < 0 || r < 0 || Math.abs(l - r) > 1) return -1;
        return 1 + Math.max(l, r);
    }
}
// test
check(new Solution().isBalanced(TreeNode.of(3, 9, 20, null, null, 15, 7)), true);
check(new Solution().isBalanced(TreeNode.of(1, 2, 2, 3, 3, null, null, 4, 4)), false);
```

## 100. Same Tree | Easy
related: 101, 572
tags: blind75, dfs

Both null → true; one null or values differ → false; else recurse on both sides.

**Complexity:** O(n).

```python
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p or not q:
            return p is q
        return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
# test
assert Solution().isSameTree(make_tree([1, 2, 3]), make_tree([1, 2, 3])) is True
assert Solution().isSameTree(make_tree([1, 2]), make_tree([1, None, 2])) is False
```

```java
class Solution {
    public boolean isSameTree(TreeNode p, TreeNode q) {
        if (p == null || q == null) return p == q;
        return p.val == q.val && isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
// test
check(new Solution().isSameTree(TreeNode.of(1, 2, 3), TreeNode.of(1, 2, 3)), true);
check(new Solution().isSameTree(TreeNode.of(1, 2), TreeNode.of(1, null, 2)), false);
```

## 101. Symmetric Tree | Easy
related: 100, 226
tags: dfs

A tree is symmetric if left and right subtrees are **mirrors**: `mirror(a, b) = a.val == b.val && mirror(a.left, b.right) && mirror(a.right, b.left)`.

**Complexity:** O(n).

```python
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def mirror(a, b):
            if not a or not b:
                return a is b
            return a.val == b.val and mirror(a.left, b.right) and mirror(a.right, b.left)

        return mirror(root.left, root.right)
# test
assert Solution().isSymmetric(make_tree([1, 2, 2, 3, 4, 4, 3])) is True
assert Solution().isSymmetric(make_tree([1, 2, 2, None, 3, None, 3])) is False
```

```java
class Solution {
    public boolean isSymmetric(TreeNode root) { return mirror(root.left, root.right); }

    private boolean mirror(TreeNode a, TreeNode b) {
        if (a == null || b == null) return a == b;
        return a.val == b.val && mirror(a.left, b.right) && mirror(a.right, b.left);
    }
}
// test
check(new Solution().isSymmetric(TreeNode.of(1, 2, 2, 3, 4, 4, 3)), true);
check(new Solution().isSymmetric(TreeNode.of(1, 2, 2, null, 3, null, 3)), false);
```

## 572. Subtree of Another Tree | Easy
related: 100
tags: blind75, dfs

For every node of `root`, check `isSameTree(node, subRoot)`. (O(m+n) alternative: serialize both and do substring search / hashing.)

**Complexity:** O(m·n) worst case.

```python
class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def same(a, b):
            if not a or not b:
                return a is b
            return a.val == b.val and same(a.left, b.left) and same(a.right, b.right)

        if not root:
            return False
        return same(root, subRoot) or self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)
# test
assert Solution().isSubtree(make_tree([3, 4, 5, 1, 2]), make_tree([4, 1, 2])) is True
assert Solution().isSubtree(make_tree([3, 4, 5, 1, 2, None, None, None, None, 0]), make_tree([4, 1, 2])) is False
```

```java
class Solution {
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        if (root == null) return false;
        return same(root, subRoot) || isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot);
    }

    private boolean same(TreeNode a, TreeNode b) {
        if (a == null || b == null) return a == b;
        return a.val == b.val && same(a.left, b.left) && same(a.right, b.right);
    }
}
// test
check(new Solution().isSubtree(TreeNode.of(3, 4, 5, 1, 2), TreeNode.of(4, 1, 2)), true);
```

## 235. Lowest Common Ancestor of a Binary Search Tree | Medium
related: 236
tags: blind75, bst

Use BST ordering: if both `p, q` are smaller go left, both larger go right; otherwise the current node splits them → LCA.

**Complexity:** O(h) time, O(1) space.

```python
class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        while root:
            if p.val < root.val and q.val < root.val:
                root = root.left
            elif p.val > root.val and q.val > root.val:
                root = root.right
            else:
                return root
# test
t = make_tree([6, 2, 8, 0, 4, 7, 9, None, None, 3, 5])
assert Solution().lowestCommonAncestor(t, find_node(t, 2), find_node(t, 8)).val == 6
assert Solution().lowestCommonAncestor(t, find_node(t, 2), find_node(t, 4)).val == 2
```

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        while (root != null) {
            if (p.val < root.val && q.val < root.val) root = root.left;
            else if (p.val > root.val && q.val > root.val) root = root.right;
            else return root;
        }
        return null;
    }
}
// test
TreeNode t = TreeNode.of(6, 2, 8, 0, 4, 7, 9, null, null, 3, 5);
check(new Solution().lowestCommonAncestor(t, TreeNode.find(t, 2), TreeNode.find(t, 8)).val, 6);
check(new Solution().lowestCommonAncestor(t, TreeNode.find(t, 2), TreeNode.find(t, 4)).val, 2);
```

## 236. Lowest Common Ancestor of a Binary Tree | Medium
related: 235, 863
tags: dfs

DFS returns `p` or `q` if found in the subtree (else null). If both sides return non-null, the current node is the LCA; otherwise bubble up whichever side is non-null.

**Complexity:** O(n).

```python
class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        if not root or root is p or root is q:
            return root
        l = self.lowestCommonAncestor(root.left, p, q)
        r = self.lowestCommonAncestor(root.right, p, q)
        return root if l and r else l or r
# test
t = make_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
assert Solution().lowestCommonAncestor(t, find_node(t, 5), find_node(t, 1)).val == 3
assert Solution().lowestCommonAncestor(t, find_node(t, 5), find_node(t, 4)).val == 5
```

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || root == p || root == q) return root;
        TreeNode l = lowestCommonAncestor(root.left, p, q), r = lowestCommonAncestor(root.right, p, q);
        if (l != null && r != null) return root;
        return l != null ? l : r;
    }
}
// test
TreeNode t = TreeNode.of(3, 5, 1, 6, 2, 0, 8, null, null, 7, 4);
check(new Solution().lowestCommonAncestor(t, TreeNode.find(t, 5), TreeNode.find(t, 4)).val, 5);
```

## 102. Binary Tree Level Order Traversal | Medium
related: 103, 199, 637
tags: blind75, bfs

BFS; process the queue in chunks of its current size = one level.

**Complexity:** O(n).

```python
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        res, q = [], deque([root] if root else [])
        while q:
            level = []
            for _ in range(len(q)):
                node = q.popleft()
                level.append(node.val)
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
            res.append(level)
        return res
# test
assert Solution().levelOrder(make_tree([3, 9, 20, None, None, 15, 7])) == [[3], [9, 20], [15, 7]]
assert Solution().levelOrder(None) == []
```

```java
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        Deque<TreeNode> q = new ArrayDeque<>();
        if (root != null) q.add(root);
        while (!q.isEmpty()) {
            List<Integer> level = new ArrayList<>();
            for (int i = q.size(); i > 0; i--) {
                TreeNode n = q.poll();
                level.add(n.val);
                if (n.left != null) q.add(n.left);
                if (n.right != null) q.add(n.right);
            }
            res.add(level);
        }
        return res;
    }
}
// test
check(new Solution().levelOrder(TreeNode.of(3, 9, 20, null, null, 15, 7)), "[[3], [9, 20], [15, 7]]");
```

## 103. Binary Tree Zigzag Level Order Traversal | Medium
related: 102
tags: bfs

Level-order BFS; reverse every other level (or append to the front of a deque).

**Complexity:** O(n).

```python
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        res, q = [], deque([root] if root else [])
        while q:
            level = []
            for _ in range(len(q)):
                node = q.popleft()
                level.append(node.val)
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
            res.append(level[::-1] if len(res) % 2 else level)
        return res
# test
assert Solution().zigzagLevelOrder(make_tree([3, 9, 20, None, None, 15, 7])) == [[3], [20, 9], [15, 7]]
```

```java
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        Deque<TreeNode> q = new ArrayDeque<>();
        if (root != null) q.add(root);
        while (!q.isEmpty()) {
            LinkedList<Integer> level = new LinkedList<>();
            boolean rev = res.size() % 2 == 1;
            for (int i = q.size(); i > 0; i--) {
                TreeNode n = q.poll();
                if (rev) level.addFirst(n.val); else level.addLast(n.val);
                if (n.left != null) q.add(n.left);
                if (n.right != null) q.add(n.right);
            }
            res.add(level);
        }
        return res;
    }
}
// test
check(new Solution().zigzagLevelOrder(TreeNode.of(3, 9, 20, null, null, 15, 7)), "[[3], [20, 9], [15, 7]]");
```

## 199. Binary Tree Right Side View | Medium
related: 102
tags: bfs

Level-order BFS; the last node of each level is visible from the right. (DFS alternative: visit right first, record first node at each new depth.)

**Complexity:** O(n).

```python
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        res, q = [], deque([root] if root else [])
        while q:
            res.append(q[-1].val)
            for _ in range(len(q)):
                node = q.popleft()
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
        return res
# test
assert Solution().rightSideView(make_tree([1, 2, 3, None, 5, None, 4])) == [1, 3, 4]
assert Solution().rightSideView(make_tree([1, 2, 3, 4])) == [1, 3, 4]
```

```java
class Solution {
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        dfs(root, 0, res);
        return res;
    }

    private void dfs(TreeNode n, int depth, List<Integer> res) {
        if (n == null) return;
        if (depth == res.size()) res.add(n.val);   // first node seen at this depth (right-first)
        dfs(n.right, depth + 1, res);
        dfs(n.left, depth + 1, res);
    }
}
// test
check(new Solution().rightSideView(TreeNode.of(1, 2, 3, null, 5, null, 4)), "[1, 3, 4]");
check(new Solution().rightSideView(TreeNode.of(1, 2, 3, 4)), "[1, 3, 4]");
```

## 1448. Count Good Nodes in Binary Tree | Medium
related: 98
tags: dfs

Pass down the max value on the path from the root; a node is good if `node.val ≥ pathMax`.

**Complexity:** O(n).

```python
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, mx):
            if not node:
                return 0
            good = node.val >= mx
            mx = max(mx, node.val)
            return good + dfs(node.left, mx) + dfs(node.right, mx)

        return dfs(root, root.val)
# test
assert Solution().goodNodes(make_tree([3, 1, 4, 3, None, 1, 5])) == 4
assert Solution().goodNodes(make_tree([3, 3, None, 4, 2])) == 3
```

```java
class Solution {
    public int goodNodes(TreeNode root) { return dfs(root, Integer.MIN_VALUE); }

    private int dfs(TreeNode n, int mx) {
        if (n == null) return 0;
        int good = n.val >= mx ? 1 : 0;
        mx = Math.max(mx, n.val);
        return good + dfs(n.left, mx) + dfs(n.right, mx);
    }
}
// test
check(new Solution().goodNodes(TreeNode.of(3, 1, 4, 3, null, 1, 5)), 4);
```

## 98. Validate Binary Search Tree | Medium
related: 230, 1448
tags: blind75, bst, bounds

Each node must lie strictly within `(lo, hi)` inherited from ancestors — checking only direct children is the classic bug.

```viz
      5            (−∞, ∞)
    1   6          1:(−∞,5)  6:(5,∞)
       3  7        3:(5,6)  ✗  3 < 5
```

**Complexity:** O(n).

```python
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def ok(node, lo, hi):
            if not node:
                return True
            if not lo < node.val < hi:
                return False
            return ok(node.left, lo, node.val) and ok(node.right, node.val, hi)

        return ok(root, float("-inf"), float("inf"))
# test
assert Solution().isValidBST(make_tree([2, 1, 3])) is True
assert Solution().isValidBST(make_tree([5, 1, 6, None, None, 3, 7])) is False
```

```java
class Solution {
    public boolean isValidBST(TreeNode root) { return ok(root, Long.MIN_VALUE, Long.MAX_VALUE); }

    private boolean ok(TreeNode n, long lo, long hi) {
        if (n == null) return true;
        if (n.val <= lo || n.val >= hi) return false;
        return ok(n.left, lo, n.val) && ok(n.right, n.val, hi);
    }
}
// test
check(new Solution().isValidBST(TreeNode.of(2, 1, 3)), true);
check(new Solution().isValidBST(TreeNode.of(5, 1, 6, null, null, 3, 7)), false);
```

## 230. Kth Smallest Element in a BST | Medium
related: 98, 173
tags: blind75, bst, inorder

In-order traversal of a BST is sorted; stop at the k-th visited node (iterative stack makes early exit easy).

**Complexity:** O(h + k).

```python
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        st, node = [], root
        while True:
            while node:
                st.append(node)
                node = node.left
            node = st.pop()
            k -= 1
            if k == 0:
                return node.val
            node = node.right
# test
assert Solution().kthSmallest(make_tree([3, 1, 4, None, 2]), 1) == 1
assert Solution().kthSmallest(make_tree([5, 3, 6, 2, 4, None, None, 1]), 3) == 3
```

```java
class Solution {
    public int kthSmallest(TreeNode root, int k) {
        Deque<TreeNode> st = new ArrayDeque<>();
        TreeNode n = root;
        while (true) {
            while (n != null) { st.push(n); n = n.left; }
            n = st.pop();
            if (--k == 0) return n.val;
            n = n.right;
        }
    }
}
// test
check(new Solution().kthSmallest(TreeNode.of(5, 3, 6, 2, 4, null, null, 1), 3), 3);
```

## 105. Construct Binary Tree from Preorder and Inorder Traversal | Medium
related: 106, 297
tags: blind75, divide-and-conquer

`preorder[0]` is the root; its position `m` in inorder splits left/right subtrees (left size = `m − inLo`). A value→index map makes each split O(1).

```viz
pre [3 | 9 | 20 15 7]     in [9 | 3 | 15 20 7]
root 3, left = {9}, right = {20,15,7}  -> recurse
```

**Complexity:** O(n).

```python
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        pos = {v: i for i, v in enumerate(inorder)}
        it = iter(preorder)

        def build(lo, hi):              # inorder range [lo, hi)
            if lo >= hi:
                return None
            root = TreeNode(next(it))
            m = pos[root.val]
            root.left = build(lo, m)    # consume preorder left subtree first
            root.right = build(m + 1, hi)
            return root

        return build(0, len(inorder))
# test
assert tree_vals(Solution().buildTree([3, 9, 20, 15, 7], [9, 3, 15, 20, 7])) == [3, 9, 20, None, None, 15, 7]
```

```java
class Solution {
    private int pre = 0;
    private final Map<Integer, Integer> pos = new HashMap<>();

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        for (int i = 0; i < inorder.length; i++) pos.put(inorder[i], i);
        return build(preorder, 0, inorder.length);
    }

    private TreeNode build(int[] preorder, int lo, int hi) {
        if (lo >= hi) return null;
        TreeNode root = new TreeNode(preorder[pre++]);
        int m = pos.get(root.val);
        root.left = build(preorder, lo, m);
        root.right = build(preorder, m + 1, hi);
        return root;
    }
}
// test
check(new Solution().buildTree(new int[]{3, 9, 20, 15, 7}, new int[]{9, 3, 15, 20, 7}), TreeNode.of(3, 9, 20, null, null, 15, 7));
```

## 124. Binary Tree Maximum Path Sum | Hard
related: 543, 687
tags: blind75, dfs, global-answer

Each node returns the best **downward** path gain (`max(0, left, right) + val`, negatives dropped). The best path *through* the node uses both sides: `val + max(0,L) + max(0,R)` → update global answer.

```viz
     -10
     /  \
    9    20        through 20: 15 + 20 + 7 = 42  (answer)
        /  \       returns up: 20 + 15 = 35
       15   7
```

**Complexity:** O(n).

```python
class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        best = float("-inf")

        def gain(node):
            nonlocal best
            if not node:
                return 0
            l, r = max(gain(node.left), 0), max(gain(node.right), 0)
            best = max(best, node.val + l + r)
            return node.val + max(l, r)

        gain(root)
        return best
# test
assert Solution().maxPathSum(make_tree([1, 2, 3])) == 6
assert Solution().maxPathSum(make_tree([-10, 9, 20, None, None, 15, 7])) == 42
assert Solution().maxPathSum(make_tree([-3])) == -3
```

```java
class Solution {
    private int best = Integer.MIN_VALUE;

    public int maxPathSum(TreeNode root) {
        gain(root);
        return best;
    }

    private int gain(TreeNode n) {
        if (n == null) return 0;
        int l = Math.max(gain(n.left), 0), r = Math.max(gain(n.right), 0);
        best = Math.max(best, n.val + l + r);
        return n.val + Math.max(l, r);
    }
}
// test
check(new Solution().maxPathSum(TreeNode.of(-10, 9, 20, null, null, 15, 7)), 42);
check(new Solution().maxPathSum(TreeNode.of(-3)), -3);
```

## 297. Serialize and Deserialize Binary Tree | Hard
related: 105, 449
tags: blind75, design, preorder

Preorder with explicit null markers (`#`) uniquely identifies a tree. Deserialize by consuming tokens in the same preorder.

```viz
    1
   / \        "1,2,#,#,3,4,#,#,5,#,#"
  2   3
     / \
    4   5
```

**Complexity:** O(n) both ways.

```python
class Codec:
    def serialize(self, root: Optional[TreeNode]) -> str:
        out = []

        def dfs(node):
            if not node:
                out.append("#")
                return
            out.append(str(node.val))
            dfs(node.left)
            dfs(node.right)

        dfs(root)
        return ",".join(out)

    def deserialize(self, data: str) -> Optional[TreeNode]:
        it = iter(data.split(","))

        def build():
            v = next(it)
            if v == "#":
                return None
            node = TreeNode(int(v))
            node.left = build()
            node.right = build()
            return node

        return build()
# test
c = Codec()
t = make_tree([1, 2, 3, None, None, 4, 5])
assert c.serialize(t) == "1,2,#,#,3,4,#,#,5,#,#"
assert tree_vals(c.deserialize(c.serialize(t))) == [1, 2, 3, None, None, 4, 5]
assert c.deserialize(c.serialize(None)) is None
```

```java
class Codec {
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        ser(root, sb);
        return sb.substring(0, sb.length() - 1);
    }

    private void ser(TreeNode n, StringBuilder sb) {
        if (n == null) { sb.append("#,"); return; }
        sb.append(n.val).append(',');
        ser(n.left, sb);
        ser(n.right, sb);
    }

    public TreeNode deserialize(String data) {
        return des(new ArrayDeque<>(Arrays.asList(data.split(","))));
    }

    private TreeNode des(Deque<String> q) {
        String v = q.poll();
        if (v.equals("#")) return null;
        TreeNode n = new TreeNode(Integer.parseInt(v));
        n.left = des(q);
        n.right = des(q);
        return n;
    }
}
// test
Codec c = new Codec();
TreeNode t = TreeNode.of(1, 2, 3, null, null, 4, 5);
check(c.serialize(t), "1,2,#,#,3,4,#,#,5,#,#");
check(c.deserialize(c.serialize(t)), t);
```

## 112. Path Sum | Easy
related: 113, 437
tags: dfs

Subtract node values on the way down; at a **leaf**, check the remainder equals the leaf value.

**Complexity:** O(n).

```python
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False
        if not root.left and not root.right:
            return root.val == targetSum
        rest = targetSum - root.val
        return self.hasPathSum(root.left, rest) or self.hasPathSum(root.right, rest)
# test
t = make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, None, 1])
assert Solution().hasPathSum(t, 22) is True
assert Solution().hasPathSum(make_tree([1, 2, 3]), 5) is False
```

```java
class Solution {
    public boolean hasPathSum(TreeNode root, int targetSum) {
        if (root == null) return false;
        if (root.left == null && root.right == null) return root.val == targetSum;
        int rest = targetSum - root.val;
        return hasPathSum(root.left, rest) || hasPathSum(root.right, rest);
    }
}
// test
check(new Solution().hasPathSum(TreeNode.of(5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1), 22), true);
```

## 113. Path Sum II | Medium
related: 112, 437, 257
tags: dfs, backtracking

DFS with a path list; at a leaf with the right remainder, copy the path. Pop on the way back (backtracking).

**Complexity:** O(n²) worst case (copying paths).

```python
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        res, path = [], []

        def dfs(node, rest):
            if not node:
                return
            path.append(node.val)
            rest -= node.val
            if not node.left and not node.right and rest == 0:
                res.append(path[:])
            dfs(node.left, rest)
            dfs(node.right, rest)
            path.pop()

        dfs(root, targetSum)
        return res
# test
t = make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1])
assert Solution().pathSum(t, 22) == [[5, 4, 11, 2], [5, 8, 4, 5]]
```

```java
class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(root, targetSum, new ArrayList<>(), res);
        return res;
    }

    private void dfs(TreeNode n, int rest, List<Integer> path, List<List<Integer>> res) {
        if (n == null) return;
        path.add(n.val);
        rest -= n.val;
        if (n.left == null && n.right == null && rest == 0) res.add(new ArrayList<>(path));
        dfs(n.left, rest, path, res);
        dfs(n.right, rest, path, res);
        path.remove(path.size() - 1);
    }
}
// test
check(new Solution().pathSum(TreeNode.of(5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1), 22), "[[5, 4, 11, 2], [5, 8, 4, 5]]");
```

## 437. Path Sum III | Medium
related: 560, 112, 113
tags: prefix-sum, dfs

Paths go downward but can start anywhere → **prefix sums on the root-to-node path** (like 560). Count earlier prefixes equal to `cur − target`; undo the count when backtracking.

**Complexity:** O(n).

```python
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        seen = defaultdict(int, {0: 1})

        def dfs(node, cur):
            if not node:
                return 0
            cur += node.val
            cnt = seen[cur - targetSum]
            seen[cur] += 1
            cnt += dfs(node.left, cur) + dfs(node.right, cur)
            seen[cur] -= 1                     # backtrack
            return cnt

        return dfs(root, 0)
# test
assert Solution().pathSum(make_tree([10, 5, -3, 3, 2, None, 11, 3, -2, None, 1]), 8) == 3
assert Solution().pathSum(make_tree([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1]), 22) == 3
```

```java
class Solution {
    private final Map<Long, Integer> seen = new HashMap<>();

    public int pathSum(TreeNode root, int targetSum) {
        seen.put(0L, 1);
        return dfs(root, 0L, targetSum);
    }

    private int dfs(TreeNode n, long cur, int target) {
        if (n == null) return 0;
        cur += n.val;
        int cnt = seen.getOrDefault(cur - target, 0);
        seen.merge(cur, 1, Integer::sum);
        cnt += dfs(n.left, cur, target) + dfs(n.right, cur, target);
        seen.merge(cur, -1, Integer::sum);
        return cnt;
    }
}
// test
check(new Solution().pathSum(TreeNode.of(10, 5, -3, 3, 2, null, 11, 3, -2, null, 1), 8), 3);
```

## 114. Flatten Binary Tree to Linked List | Medium
related: 430
tags: morris, in-place

For each node with a left child: find the rightmost node of the left subtree, hook the right subtree there, move left to right. Walk right. O(1) space (Morris-style).

```viz
    1            1
   / \            \
  2   5    ->      2
 / \   \            \
3   4   6            3 - 4 - 5 - 6
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        cur = root
        while cur:
            if cur.left:
                p = cur.left
                while p.right:
                    p = p.right
                p.right = cur.right
                cur.right, cur.left = cur.left, None
            cur = cur.right
# test
t = make_tree([1, 2, 5, 3, 4, None, 6]); Solution().flatten(t)
assert tree_vals(t) == [1, None, 2, None, 3, None, 4, None, 5, None, 6]
```

```java
class Solution {
    public void flatten(TreeNode root) {
        for (TreeNode cur = root; cur != null; cur = cur.right) {
            if (cur.left != null) {
                TreeNode p = cur.left;
                while (p.right != null) p = p.right;
                p.right = cur.right;
                cur.right = cur.left;
                cur.left = null;
            }
        }
    }
}
// test
TreeNode t = TreeNode.of(1, 2, 5, 3, 4, null, 6);
new Solution().flatten(t);
check(t, TreeNode.of(1, null, 2, null, 3, null, 4, null, 5, null, 6));
```

## 116. Populating Next Right Pointers in Each Node | Medium
related: 117, 102
tags: bfs, perfect-tree

Perfect tree → use the already-built `next` pointers of level *i* to wire level *i+1*: `left.next = right`, `right.next = node.next.left`. O(1) extra space.

**Complexity:** O(n) time, O(1) space.

```python
class Node:  # provided by LeetCode
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val, self.left, self.right, self.next = val, left, right, next

class Solution:
    def connect(self, root: "Optional[Node]") -> "Optional[Node]":
        leftmost = root
        while leftmost and leftmost.left:
            node = leftmost
            while node:
                node.left.next = node.right
                if node.next:
                    node.right.next = node.next.left
                node = node.next
            leftmost = leftmost.left
        return root
# test
n = [Node(i) for i in range(8)]  # 1..7 used, heap layout
for i in range(1, 4):
    n[i].left, n[i].right = n[2 * i], n[2 * i + 1]
Solution().connect(n[1])
assert n[4].next is n[5] and n[5].next is n[6] and n[6].next is n[7] and n[7].next is None and n[2].next is n[3]
```

```java
class Node { // provided by LeetCode
    int val; Node left, right, next;
    Node(int val) { this.val = val; }
}

class Solution {
    public Node connect(Node root) {
        for (Node leftmost = root; leftmost != null && leftmost.left != null; leftmost = leftmost.left)
            for (Node n = leftmost; n != null; n = n.next) {
                n.left.next = n.right;
                if (n.next != null) n.right.next = n.next.left;
            }
        return root;
    }
}
// test
Node[] n = new Node[8];
for (int i = 1; i < 8; i++) n[i] = new Node(i);
for (int i = 1; i < 4; i++) { n[i].left = n[2 * i]; n[i].right = n[2 * i + 1]; }
new Solution().connect(n[1]);
check(n[4].next == n[5] && n[5].next == n[6] && n[6].next == n[7] && n[7].next == null, true);
```

## 108. Convert Sorted Array to Binary Search Tree | Easy
related: 109, 98
tags: divide-and-conquer

Middle element as root keeps the tree height-balanced; recurse on both halves.

**Complexity:** O(n).

```python
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def build(lo, hi):
            if lo > hi:
                return None
            mid = (lo + hi) // 2
            return TreeNode(nums[mid], build(lo, mid - 1), build(mid + 1, hi))

        return build(0, len(nums) - 1)
# test
assert tree_vals(Solution().sortedArrayToBST([-10, -3, 0, 5, 9])) == [0, -10, 5, None, -3, None, 9]
```

```java
class Solution {
    public TreeNode sortedArrayToBST(int[] nums) { return build(nums, 0, nums.length - 1); }

    private TreeNode build(int[] a, int lo, int hi) {
        if (lo > hi) return null;
        int mid = (lo + hi) / 2;
        return new TreeNode(a[mid], build(a, lo, mid - 1), build(a, mid + 1, hi));
    }
}
// test
check(new Solution().sortedArrayToBST(new int[]{-10, -3, 0, 5, 9}), TreeNode.of(0, -10, 5, null, -3, null, 9));
```

## 173. Binary Search Tree Iterator | Medium
related: 230, 94
tags: design, inorder, stack

Controlled in-order traversal: stack holds the path of left children. `next()` pops, then pushes the left spine of its right child. Amortized O(1), O(h) memory.

**Complexity:** O(1) amortized per call, O(h) space.

```python
class BSTIterator:
    def __init__(self, root: Optional[TreeNode]):
        self.st = []
        self._push_left(root)

    def _push_left(self, node):
        while node:
            self.st.append(node)
            node = node.left

    def next(self) -> int:
        node = self.st.pop()
        self._push_left(node.right)
        return node.val

    def hasNext(self) -> bool:
        return bool(self.st)
# test
it = BSTIterator(make_tree([7, 3, 15, None, None, 9, 20]))
out = []
while it.hasNext():
    out.append(it.next())
assert out == [3, 7, 9, 15, 20]
```

```java
class BSTIterator {
    private final Deque<TreeNode> st = new ArrayDeque<>();

    public BSTIterator(TreeNode root) { pushLeft(root); }

    private void pushLeft(TreeNode n) { for (; n != null; n = n.left) st.push(n); }

    public int next() {
        TreeNode n = st.pop();
        pushLeft(n.right);
        return n.val;
    }

    public boolean hasNext() { return !st.isEmpty(); }
}
// test
BSTIterator it = new BSTIterator(TreeNode.of(7, 3, 15, null, null, 9, 20));
List<Integer> out = new ArrayList<>();
while (it.hasNext()) out.add(it.next());
check(out, "[3, 7, 9, 15, 20]");
```

## 863. All Nodes Distance K in Binary Tree | Medium
related: 236, 102
tags: bfs, graph

Turn the tree into an undirected graph by recording each node's **parent**, then BFS from `target` for exactly `k` steps.

**Complexity:** O(n).

```python
class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        parent = {root: None}
        st = [root]
        while st:
            node = st.pop()
            for c in (node.left, node.right):
                if c:
                    parent[c] = node
                    st.append(c)
        q, seen = deque([target]), {target}
        for _ in range(k):
            for _ in range(len(q)):
                node = q.popleft()
                for nb in (node.left, node.right, parent[node]):
                    if nb and nb not in seen:
                        seen.add(nb)
                        q.append(nb)
        return [n.val for n in q]
# test
t = make_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
assert sorted(Solution().distanceK(t, find_node(t, 5), 2)) == [1, 4, 7]
```

```java
class Solution {
    public List<Integer> distanceK(TreeNode root, TreeNode target, int k) {
        Map<TreeNode, TreeNode> parent = new HashMap<>();
        Deque<TreeNode> st = new ArrayDeque<>(List.of(root));
        while (!st.isEmpty()) {
            TreeNode n = st.pop();
            for (TreeNode c : new TreeNode[]{n.left, n.right})
                if (c != null) { parent.put(c, n); st.push(c); }
        }
        Deque<TreeNode> q = new ArrayDeque<>(List.of(target));
        Set<TreeNode> seen = new HashSet<>(List.of(target));
        for (int d = 0; d < k; d++)
            for (int i = q.size(); i > 0; i--) {
                TreeNode n = q.poll();
                for (TreeNode nb : new TreeNode[]{n.left, n.right, parent.get(n)})
                    if (nb != null && seen.add(nb)) q.add(nb);
            }
        List<Integer> res = new ArrayList<>();
        for (TreeNode n : q) res.add(n.val);
        return res;
    }
}
// test
TreeNode t = TreeNode.of(3, 5, 1, 6, 2, 0, 8, null, null, 7, 4);
List<Integer> r = new Solution().distanceK(t, TreeNode.find(t, 5), 2);
Collections.sort(r);
check(r, "[1, 4, 7]");
```
