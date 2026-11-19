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


## 94. Binary Tree Inorder Traversal | Easy
related: 230, 173, 144
tags: inorder, stack

Iterative in-order: push the left spine, pop a node, visit it, then continue with its right subtree. (Morris traversal does it in O(1) space by threading.)

**Complexity:** O(n) time, O(h) space.

```python
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res, st, node = [], [], root
        while st or node:
            while node:
                st.append(node)
                node = node.left
            node = st.pop()
            res.append(node.val)
            node = node.right
        return res
# test
assert Solution().inorderTraversal(make_tree([1, None, 2, 3])) == [1, 3, 2]
assert Solution().inorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [4, 2, 6, 5, 7, 1, 3, 9, 8]
```

```java
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        Deque<TreeNode> st = new ArrayDeque<>();
        TreeNode n = root;
        while (n != null || !st.isEmpty()) {
            while (n != null) { st.push(n); n = n.left; }
            n = st.pop();
            res.add(n.val);
            n = n.right;
        }
        return res;
    }
}
// test
check(new Solution().inorderTraversal(TreeNode.of(1, null, 2, 3)), "[1, 3, 2]");
```

## 111. Minimum Depth of Binary Tree | Easy
related: 104
tags: bfs

BFS level by level; the first **leaf** reached gives the minimum depth (DFS must be careful: a node with one child isn't a leaf).

**Complexity:** O(n).

```python
class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        q, depth = deque([root]), 1
        while q:
            for _ in range(len(q)):
                n = q.popleft()
                if not n.left and not n.right:
                    return depth
                q.extend(c for c in (n.left, n.right) if c)
            depth += 1
# test
assert Solution().minDepth(make_tree([3, 9, 20, None, None, 15, 7])) == 2
assert Solution().minDepth(make_tree([2, None, 3, None, 4, None, 5, None, 6])) == 5
```

```java
class Solution {
    public int minDepth(TreeNode root) {
        if (root == null) return 0;
        if (root.left == null) return 1 + minDepth(root.right);
        if (root.right == null) return 1 + minDepth(root.left);
        return 1 + Math.min(minDepth(root.left), minDepth(root.right));
    }
}
// test
check(new Solution().minDepth(TreeNode.of(2, null, 3, null, 4, null, 5, null, 6)), 5);
```

## 106. Construct Binary Tree from Inorder and Postorder Traversal | Medium
related: 105
tags: divide-and-conquer

Mirror of 105: the **last** postorder element is the root. Consume postorder from the end and build the **right** subtree first.

**Complexity:** O(n).

```python
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        pos = {v: i for i, v in enumerate(inorder)}

        def build(lo, hi):
            if lo >= hi:
                return None
            root = TreeNode(postorder.pop())
            m = pos[root.val]
            root.right = build(m + 1, hi)
            root.left = build(lo, m)
            return root

        return build(0, len(inorder))
# test
assert tree_vals(Solution().buildTree([9, 3, 15, 20, 7], [9, 15, 7, 20, 3])) == [3, 9, 20, None, None, 15, 7]
```

```java
class Solution {
    private int post;
    private final Map<Integer, Integer> pos = new HashMap<>();

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        for (int i = 0; i < inorder.length; i++) pos.put(inorder[i], i);
        post = postorder.length - 1;
        return build(postorder, 0, inorder.length);
    }

    private TreeNode build(int[] postorder, int lo, int hi) {
        if (lo >= hi) return null;
        TreeNode root = new TreeNode(postorder[post--]);
        int m = pos.get(root.val);
        root.right = build(postorder, m + 1, hi);
        root.left = build(postorder, lo, m);
        return root;
    }
}
// test
check(new Solution().buildTree(new int[]{9, 3, 15, 20, 7}, new int[]{9, 15, 7, 20, 3}), TreeNode.of(3, 9, 20, null, null, 15, 7));
```

## 109. Convert Sorted List to Binary Search Tree | Medium
related: 108
tags: inorder-simulation

Count the nodes, then build the tree **in in-order**: recursively build the left half, consume the current list node as root, build the right half. The list pointer advances exactly in sorted order → O(n).

**Complexity:** O(n) time, O(log n) stack.

```python
class Solution:
    def sortedListToBST(self, head: Optional[ListNode]) -> Optional[TreeNode]:
        n, p = 0, head
        while p:
            n, p = n + 1, p.next
        cur = head

        def build(lo, hi):
            nonlocal cur
            if lo > hi:
                return None
            mid = (lo + hi) // 2
            left = build(lo, mid - 1)
            root = TreeNode(cur.val, left)
            cur = cur.next
            root.right = build(mid + 1, hi)
            return root

        return build(0, n - 1)
# test
assert tree_vals(Solution().sortedListToBST(make_list([-10, -3, 0, 5, 9]))) == [0, -10, 5, None, -3, None, 9]
```

```java
class Solution {
    private ListNode cur;

    public TreeNode sortedListToBST(ListNode head) {
        int n = 0;
        for (ListNode p = head; p != null; p = p.next) n++;
        cur = head;
        return build(0, n - 1);
    }

    private TreeNode build(int lo, int hi) {
        if (lo > hi) return null;
        int mid = (lo + hi) / 2;
        TreeNode left = build(lo, mid - 1);
        TreeNode root = new TreeNode(cur.val);
        cur = cur.next;
        root.left = left;
        root.right = build(mid + 1, hi);
        return root;
    }
}
// test
check(new Solution().sortedListToBST(ListNode.of(-10, -3, 0, 5, 9)), TreeNode.of(0, -10, 5, null, -3, null, 9));
```

## 117. Populating Next Right Pointers in Each Node II | Medium
related: 116, 102
tags: bfs, level-linking

Arbitrary tree. Walk the current level through its `next` pointers while building the next level's chain with a **dummy head + tail** pointer. O(1) extra space.

**Complexity:** O(n) time, O(1) space.

```python
class Node:  # provided by LeetCode
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val, self.left, self.right, self.next = val, left, right, next

class Solution:
    def connect(self, root: "Node") -> "Node":
        level = root
        while level:
            dummy = tail = Node(0)
            node = level
            while node:
                for child in (node.left, node.right):
                    if child:
                        tail.next = tail = child
                node = node.next
            level = dummy.next
        return root
# test
n = {i: Node(i) for i in range(1, 8)}
n[1].left, n[1].right = n[2], n[3]
n[2].left, n[2].right, n[3].right = n[4], n[5], n[7]
Solution().connect(n[1])
assert n[2].next is n[3] and n[4].next is n[5] and n[5].next is n[7] and n[7].next is None
```

```java
class Node { // provided by LeetCode
    int val; Node left, right, next;
    Node(int val) { this.val = val; }
}

class Solution {
    public Node connect(Node root) {
        for (Node level = root; level != null; ) {
            Node dummy = new Node(0), tail = dummy;
            for (Node n = level; n != null; n = n.next) {
                if (n.left != null) tail = tail.next = n.left;
                if (n.right != null) tail = tail.next = n.right;
            }
            level = dummy.next;
        }
        return root;
    }
}
// test
Node[] n = new Node[8];
for (int i = 1; i < 8; i++) n[i] = new Node(i);
n[1].left = n[2]; n[1].right = n[3]; n[2].left = n[4]; n[2].right = n[5]; n[3].right = n[7];
new Solution().connect(n[1]);
check(n[4].next == n[5] && n[5].next == n[7] && n[7].next == null && n[2].next == n[3], true);
```

## 129. Sum Root to Leaf Numbers | Medium
related: 112, 257
tags: dfs

Carry the number built so far: `cur = cur·10 + val`; at a leaf, add it to the total.

**Complexity:** O(n).

```python
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        def dfs(node, cur):
            if not node:
                return 0
            cur = cur * 10 + node.val
            if not node.left and not node.right:
                return cur
            return dfs(node.left, cur) + dfs(node.right, cur)

        return dfs(root, 0)
# test
assert Solution().sumNumbers(make_tree([1, 2, 3])) == 25
assert Solution().sumNumbers(make_tree([4, 9, 0, 5, 1])) == 1026
```

```java
class Solution {
    public int sumNumbers(TreeNode root) { return dfs(root, 0); }

    private int dfs(TreeNode n, int cur) {
        if (n == null) return 0;
        cur = cur * 10 + n.val;
        if (n.left == null && n.right == null) return cur;
        return dfs(n.left, cur) + dfs(n.right, cur);
    }
}
// test
check(new Solution().sumNumbers(TreeNode.of(4, 9, 0, 5, 1)), 1026);
```

## 257. Binary Tree Paths | Easy
related: 113, 129
tags: dfs, backtracking

DFS with the path so far; at a leaf, join it with `->`.

**Complexity:** O(n · h).

```python
class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        res = []

        def dfs(node, path):
            if not node:
                return
            path = path + [str(node.val)]
            if not node.left and not node.right:
                res.append("->".join(path))
            dfs(node.left, path)
            dfs(node.right, path)

        dfs(root, [])
        return res
# test
assert Solution().binaryTreePaths(make_tree([1, 2, 3, None, 5])) == ["1->2->5", "1->3"]
```

```java
class Solution {
    public List<String> binaryTreePaths(TreeNode root) {
        List<String> res = new ArrayList<>();
        dfs(root, "", res);
        return res;
    }

    private void dfs(TreeNode n, String path, List<String> res) {
        if (n == null) return;
        path = path.isEmpty() ? String.valueOf(n.val) : path + "->" + n.val;
        if (n.left == null && n.right == null) res.add(path);
        dfs(n.left, path, res);
        dfs(n.right, path, res);
    }
}
// test
check(new Solution().binaryTreePaths(TreeNode.of(1, 2, 3, null, 5)), "[1->2->5, 1->3]");
```

## 222. Count Complete Tree Nodes | Easy
related: 958
tags: binary-search, complete-tree

Compare the leftmost and rightmost depths. Equal ⇒ the subtree is perfect: `2^h − 1` nodes. Otherwise recurse into both children — only one side is non-perfect at each level.

**Complexity:** O(log² n).

```python
class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        l = r = root
        hl = hr = 0
        while l:
            l, hl = l.left, hl + 1
        while r:
            r, hr = r.right, hr + 1
        if hl == hr:
            return (1 << hl) - 1
        return 1 + self.countNodes(root.left) + self.countNodes(root.right)
# test
assert Solution().countNodes(make_tree([1, 2, 3, 4, 5, 6])) == 6
assert Solution().countNodes(make_tree(list(range(1, 16)))) == 15
assert Solution().countNodes(None) == 0
```

```java
class Solution {
    public int countNodes(TreeNode root) {
        if (root == null) return 0;
        int hl = 0, hr = 0;
        for (TreeNode l = root; l != null; l = l.left) hl++;
        for (TreeNode r = root; r != null; r = r.right) hr++;
        if (hl == hr) return (1 << hl) - 1;
        return 1 + countNodes(root.left) + countNodes(root.right);
    }
}
// test
check(new Solution().countNodes(TreeNode.of(1, 2, 3, 4, 5, 6)), 6);
```

## 337. House Robber III | Medium
related: 198, 213
tags: tree-dp

Each node returns a pair `(rob this node, skip this node)`. Rob → children must be skipped; skip → each child takes its best.

**Complexity:** O(n).

```python
class Solution:
    def rob(self, root: Optional[TreeNode]) -> int:
        def dfs(node):
            if not node:
                return 0, 0
            lr, ls = dfs(node.left)
            rr, rs = dfs(node.right)
            return node.val + ls + rs, max(lr, ls) + max(rr, rs)

        return max(dfs(root))
# test
assert Solution().rob(make_tree([3, 2, 3, None, 3, None, 1])) == 7
assert Solution().rob(make_tree([3, 4, 5, 1, 3, None, 1])) == 9
```

```java
class Solution {
    public int rob(TreeNode root) {
        int[] r = dfs(root);
        return Math.max(r[0], r[1]);
    }

    private int[] dfs(TreeNode n) {   // {rob, skip}
        if (n == null) return new int[2];
        int[] l = dfs(n.left), r = dfs(n.right);
        return new int[]{n.val + l[1] + r[1], Math.max(l[0], l[1]) + Math.max(r[0], r[1])};
    }
}
// test
check(new Solution().rob(TreeNode.of(3, 4, 5, 1, 3, null, 1)), 9);
```

## 450. Delete Node in a BST | Medium
related: 701, 98
tags: bst

Find the node. 0 or 1 child → replace it with that child. 2 children → copy the **in-order successor** (min of the right subtree) into it, then delete the successor from the right subtree.

**Complexity:** O(h).

```python
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if not root:
            return None
        if key < root.val:
            root.left = self.deleteNode(root.left, key)
        elif key > root.val:
            root.right = self.deleteNode(root.right, key)
        else:
            if not root.left or not root.right:
                return root.left or root.right
            succ = root.right
            while succ.left:
                succ = succ.left
            root.val = succ.val
            root.right = self.deleteNode(root.right, succ.val)
        return root
# test
assert tree_vals(Solution().deleteNode(make_tree([5, 3, 6, 2, 4, None, 7]), 3)) == [5, 4, 6, 2, None, None, 7]
assert tree_vals(Solution().deleteNode(make_tree([5, 3, 6, 2, 4, None, 7]), 0)) == [5, 3, 6, 2, 4, None, 7]
```

```java
class Solution {
    public TreeNode deleteNode(TreeNode root, int key) {
        if (root == null) return null;
        if (key < root.val) root.left = deleteNode(root.left, key);
        else if (key > root.val) root.right = deleteNode(root.right, key);
        else {
            if (root.left == null) return root.right;
            if (root.right == null) return root.left;
            TreeNode succ = root.right;
            while (succ.left != null) succ = succ.left;
            root.val = succ.val;
            root.right = deleteNode(root.right, succ.val);
        }
        return root;
    }
}
// test
check(new Solution().deleteNode(TreeNode.of(5, 3, 6, 2, 4, null, 7), 3), TreeNode.of(5, 4, 6, 2, null, null, 7));
```

## 530. Minimum Absolute Difference in BST | Easy
related: 98, 230
tags: inorder

In-order visits values in sorted order, so the minimum difference is between consecutive visited values.

**Complexity:** O(n).

```python
class Solution:
    def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
        best, prev = float("inf"), None

        def inorder(node):
            nonlocal best, prev
            if not node:
                return
            inorder(node.left)
            if prev is not None:
                best = min(best, node.val - prev)
            prev = node.val
            inorder(node.right)

        inorder(root)
        return best
# test
assert Solution().getMinimumDifference(make_tree([4, 2, 6, 1, 3])) == 1
assert Solution().getMinimumDifference(make_tree([1, 0, 48, None, None, 12, 49])) == 1
```

```java
class Solution {
    private int best = Integer.MAX_VALUE;
    private Integer prev = null;

    public int getMinimumDifference(TreeNode root) {
        inorder(root);
        return best;
    }

    private void inorder(TreeNode n) {
        if (n == null) return;
        inorder(n.left);
        if (prev != null) best = Math.min(best, n.val - prev);
        prev = n.val;
        inorder(n.right);
    }
}
// test
check(new Solution().getMinimumDifference(TreeNode.of(4, 2, 6, 1, 3)), 1);
```

## 637. Average of Levels in Binary Tree | Easy
related: 102
tags: bfs

Level-order BFS; average each level's values.

**Complexity:** O(n).

```python
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        res, q = [], deque([root])
        while q:
            s, k = 0, len(q)
            for _ in range(k):
                n = q.popleft()
                s += n.val
                q.extend(c for c in (n.left, n.right) if c)
            res.append(s / k)
        return res
# test
assert Solution().averageOfLevels(make_tree([3, 9, 20, None, None, 15, 7])) == [3.0, 14.5, 11.0]
```

```java
class Solution {
    public List<Double> averageOfLevels(TreeNode root) {
        List<Double> res = new ArrayList<>();
        Deque<TreeNode> q = new ArrayDeque<>(List.of(root));
        while (!q.isEmpty()) {
            int k = q.size();
            long s = 0;
            for (int i = 0; i < k; i++) {
                TreeNode n = q.poll();
                s += n.val;
                if (n.left != null) q.add(n.left);
                if (n.right != null) q.add(n.right);
            }
            res.add((double) s / k);
        }
        return res;
    }
}
// test
check(new Solution().averageOfLevels(TreeNode.of(3, 9, 20, null, null, 15, 7)), "[3.0, 14.5, 11.0]");
```

## 662. Maximum Width of Binary Tree | Medium
related: 102, 958
tags: bfs, heap-indexing

Give nodes heap indices (`left = 2i`, `right = 2i+1`). Width of a level = last index − first index + 1. Re-base indices per level to avoid huge numbers.

**Complexity:** O(n).

```python
class Solution:
    def widthOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        best, q = 0, [(root, 0)]
        while q:
            base = q[0][1]
            best = max(best, q[-1][1] - base + 1)
            nxt = []
            for node, i in q:
                i -= base
                if node.left:
                    nxt.append((node.left, 2 * i))
                if node.right:
                    nxt.append((node.right, 2 * i + 1))
            q = nxt
        return best
# test
assert Solution().widthOfBinaryTree(make_tree([1, 3, 2, 5, 3, None, 9])) == 4
assert Solution().widthOfBinaryTree(make_tree([1, 3, 2, 5, None, None, 9, 6, None, 7])) == 7
```

```java
class Solution {
    public int widthOfBinaryTree(TreeNode root) {
        int best = 0;
        Deque<Object[]> q = new ArrayDeque<>();
        q.add(new Object[]{root, 0L});
        while (!q.isEmpty()) {
            long base = (long) q.peekFirst()[1], last = (long) q.peekLast()[1];
            best = (int) Math.max(best, last - base + 1);
            for (int k = q.size(); k > 0; k--) {
                Object[] e = q.poll();
                TreeNode n = (TreeNode) e[0];
                long i = (long) e[1] - base;
                if (n.left != null) q.add(new Object[]{n.left, 2 * i});
                if (n.right != null) q.add(new Object[]{n.right, 2 * i + 1});
            }
        }
        return best;
    }
}
// test
check(new Solution().widthOfBinaryTree(TreeNode.of(1, 3, 2, 5, null, null, 9, 6, null, 7)), 7);
```

## 687. Longest Univalue Path | Medium
related: 543, 124
tags: dfs, global-answer

Like the diameter: each node returns the longest downward arm with its own value. Arms from children count only if the child has the same value; update the global with `left + right`.

**Complexity:** O(n).

```python
class Solution:
    def longestUnivaluePath(self, root: Optional[TreeNode]) -> int:
        best = 0

        def arm(node):
            nonlocal best
            if not node:
                return 0
            l, r = arm(node.left), arm(node.right)
            l = l + 1 if node.left and node.left.val == node.val else 0
            r = r + 1 if node.right and node.right.val == node.val else 0
            best = max(best, l + r)
            return max(l, r)

        arm(root)
        return best
# test
assert Solution().longestUnivaluePath(make_tree([5, 4, 5, 1, 1, None, 5])) == 2
assert Solution().longestUnivaluePath(make_tree([1, 4, 5, 4, 4, None, 5])) == 2
```

```java
class Solution {
    private int best = 0;

    public int longestUnivaluePath(TreeNode root) {
        arm(root);
        return best;
    }

    private int arm(TreeNode n) {
        if (n == null) return 0;
        int l = arm(n.left), r = arm(n.right);
        l = (n.left != null && n.left.val == n.val) ? l + 1 : 0;
        r = (n.right != null && n.right.val == n.val) ? r + 1 : 0;
        best = Math.max(best, l + r);
        return Math.max(l, r);
    }
}
// test
check(new Solution().longestUnivaluePath(TreeNode.of(5, 4, 5, 1, 1, null, 5)), 2);
```

## 958. Check Completeness of a Binary Tree | Medium
related: 222, 662
tags: bfs

BFS including null children. In a complete tree, once you see a null, every later node in BFS order must also be null.

**Complexity:** O(n).

```python
class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        q, seen_null = deque([root]), False
        while q:
            node = q.popleft()
            if not node:
                seen_null = True
                continue
            if seen_null:
                return False
            q.append(node.left)
            q.append(node.right)
        return True
# test
assert Solution().isCompleteTree(make_tree([1, 2, 3, 4, 5, 6])) is True
assert Solution().isCompleteTree(make_tree([1, 2, 3, 4, 5, None, 7])) is False
```

```java
class Solution {
    public boolean isCompleteTree(TreeNode root) {
        Deque<TreeNode> q = new LinkedList<>();   // LinkedList allows null elements
        q.add(root);
        boolean seenNull = false;
        while (!q.isEmpty()) {
            TreeNode n = q.poll();
            if (n == null) { seenNull = true; continue; }
            if (seenNull) return false;
            q.add(n.left);
            q.add(n.right);
        }
        return true;
    }
}
// test
check(new Solution().isCompleteTree(TreeNode.of(1, 2, 3, 4, 5, 6)), true);
check(new Solution().isCompleteTree(TreeNode.of(1, 2, 3, 4, 5, null, 7)), false);
```

## 987. Vertical Order Traversal of a Binary Tree | Hard
related: 102, 199
tags: dfs, sort

Record `(col, row, val)` for every node (left child col−1, right col+1), sort, then group by column. Sorting by `(row, val)` within a column handles same-position ties.

**Complexity:** O(n log n).

```python
class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        cells = []

        def dfs(node, r, c):
            if node:
                cells.append((c, r, node.val))
                dfs(node.left, r + 1, c - 1)
                dfs(node.right, r + 1, c + 1)

        dfs(root, 0, 0)
        cells.sort()
        res, last = [], None
        for c, _, v in cells:
            if c != last:
                res.append([])
                last = c
            res[-1].append(v)
        return res
# test
assert Solution().verticalTraversal(make_tree([3, 9, 20, None, None, 15, 7])) == [[9], [3, 15], [20], [7]]
assert Solution().verticalTraversal(make_tree([1, 2, 3, 4, 6, 5, 7])) == [[4], [2], [1, 5, 6], [3], [7]]
```

```java
class Solution {
    private final List<int[]> cells = new ArrayList<>();   // {col, row, val}

    public List<List<Integer>> verticalTraversal(TreeNode root) {
        dfs(root, 0, 0);
        cells.sort((a, b) -> a[0] != b[0] ? a[0] - b[0] : a[1] != b[1] ? a[1] - b[1] : a[2] - b[2]);
        List<List<Integer>> res = new ArrayList<>();
        Integer last = null;
        for (int[] c : cells) {
            if (last == null || c[0] != last) { res.add(new ArrayList<>()); last = c[0]; }
            res.get(res.size() - 1).add(c[2]);
        }
        return res;
    }

    private void dfs(TreeNode n, int r, int c) {
        if (n == null) return;
        cells.add(new int[]{c, r, n.val});
        dfs(n.left, r + 1, c - 1);
        dfs(n.right, r + 1, c + 1);
    }
}
// test
check(new Solution().verticalTraversal(TreeNode.of(1, 2, 3, 4, 6, 5, 7)), "[[4], [2], [1, 5, 6], [3], [7]]");
```
