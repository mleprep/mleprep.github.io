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

```problem
Mirror a binary tree, swapping the left and right child of every node, and return its root. (Trees are written in level order, with `null` for a missing child.)

**Example:** `[4,2,7,1,3,6,9]` → `[4,7,2,9,6,3,1]`
```

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
assert tree_vals(Solution().invertTree(make_tree([]))) == []
assert tree_vals(Solution().invertTree(make_tree([2, 1, 3]))) == [2, 3, 1]
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

```go
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
// test
assertEq(invertTree(makeTree(4, 2, 7, 1, 3, 6, 9)), []any{4, 7, 2, 9, 6, 3, 1})
assertEq(invertTree(makeTree()), []any{})
assertEq(invertTree(makeTree(2, 1, 3)), []any{2, 3, 1})
```

## 104. Maximum Depth of Binary Tree | Easy
related: 110, 543, 111
tags: blind75, dfs

```problem
Return the maximum depth of a binary tree: the number of nodes on the longest path from the root down to a leaf.

**Example:** `[3,9,20,null,null,15,7]` → `3`
```

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

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
// test
assertEq(maxDepth(makeTree(3, 9, 20, nil, nil, 15, 7)), 3)
assertEq(maxDepth(nil), 0)
```

## 543. Diameter of Binary Tree | Easy
related: 124, 104
tags: dfs, global-answer

```problem
Return the diameter of a binary tree: the number of **edges** on the longest path between any two nodes. The path doesn't have to pass through the root.

**Example:** `[1,2,3,4,5]` → `3` (path 4 → 2 → 1 → 3)
```

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

```go
func diameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := height(node.Left), height(node.Right)
		best = max(best, l+r) // longest path through node
		return 1 + max(l, r)
	}
	height(root)
	return best
}
// test
assertEq(diameterOfBinaryTree(makeTree(1, 2, 3, 4, 5)), 3)
assertEq(diameterOfBinaryTree(makeTree(1, 2)), 1)
```

## 110. Balanced Binary Tree | Easy
related: 104
tags: dfs

```problem
Decide whether a binary tree is height-balanced: at every node, the heights of the left and right subtrees differ by at most 1.

**Example:** `[3,9,20,null,null,15,7]` → `true`; `[1,2,2,3,3,null,null,4,4]` → `false`
```

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

```go
func isBalanced(root *TreeNode) bool {
	var h func(*TreeNode) int // height, or -1 if unbalanced
	h = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := h(node.Left), h(node.Right)
		if l < 0 || r < 0 || l-r > 1 || r-l > 1 {
			return -1
		}
		return 1 + max(l, r)
	}
	return h(root) >= 0
}
// test
assertEq(isBalanced(makeTree(3, 9, 20, nil, nil, 15, 7)), true)
assertEq(isBalanced(makeTree(1, 2, 2, 3, 3, nil, nil, 4, 4)), false)
```

## 100. Same Tree | Easy
related: 101, 572
tags: blind75, dfs

```problem
Decide whether two binary trees `p` and `q` are identical: the same shape with the same value in every corresponding node.

**Example:** `[1,2,3]` and `[1,2,3]` → `true`; `[1,2]` and `[1,null,2]` → `false`
```

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

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
// test
assertEq(isSameTree(makeTree(1, 2, 3), makeTree(1, 2, 3)), true)
assertEq(isSameTree(makeTree(1, 2), makeTree(1, nil, 2)), false)
```

## 101. Symmetric Tree | Easy
related: 100, 226
tags: dfs

```problem
Decide whether a binary tree is a mirror image of itself around its centre.

**Example:** `[1,2,2,3,4,4,3]` → `true`; `[1,2,2,null,3,null,3]` → `false`

**Constraints:** follow-up: both recursively and iteratively.
```

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

```go
func isSymmetric(root *TreeNode) bool {
	var mirror func(a, b *TreeNode) bool
	mirror = func(a, b *TreeNode) bool {
		if a == nil || b == nil {
			return a == b
		}
		return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
	}
	return mirror(root.Left, root.Right)
}
// test
assertEq(isSymmetric(makeTree(1, 2, 2, 3, 4, 4, 3)), true)
assertEq(isSymmetric(makeTree(1, 2, 2, nil, 3, nil, 3)), false)
```

## 572. Subtree of Another Tree | Easy
related: 100
tags: blind75, dfs

```problem
Decide whether `subRoot` appears in `root` as a subtree: some node of `root` together with **all** of its descendants is identical to `subRoot`.

**Example:** `root = [3,4,5,1,2], subRoot = [4,1,2]` → `true`; with an extra `0` under the `2` in `root` → `false`
```

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

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return same(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func same(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && same(a.Left, b.Left) && same(a.Right, b.Right)
}
// test
assertEq(isSubtree(makeTree(3, 4, 5, 1, 2), makeTree(4, 1, 2)), true)
assertEq(isSubtree(makeTree(3, 4, 5, 1, 2, nil, nil, nil, nil, 0), makeTree(4, 1, 2)), false)
```

## 235. Lowest Common Ancestor of a Binary Search Tree | Medium
related: 236
tags: blind75, bst

```problem
Given a binary search tree and two of its nodes `p` and `q`, return their lowest common ancestor: the deepest node that has both as descendants (a node counts as its own descendant).

**Example:** `[6,2,8,0,4,7,9,null,null,3,5]`, `p = 2, q = 8` → `6`; `p = 2, q = 4` → `2`
```

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

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		switch {
		case p.Val < root.Val && q.Val < root.Val:
			root = root.Left
		case p.Val > root.Val && q.Val > root.Val:
			root = root.Right
		default: // p and q split here (or one of them is root)
			return root
		}
	}
	return nil
}
// test
t := makeTree(6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5)
assertEq(lowestCommonAncestor(t, findNode(t, 2), findNode(t, 8)).Val, 6)
assertEq(lowestCommonAncestor(t, findNode(t, 2), findNode(t, 4)).Val, 2)
```

## 236. Lowest Common Ancestor of a Binary Tree | Medium
related: 235, 863
tags: dfs

```problem
Given a binary tree (not a BST) and two of its nodes `p` and `q`, return their lowest common ancestor: the deepest node that has both as descendants (a node counts as its own descendant).

**Example:** `[3,5,1,6,2,0,8,null,null,7,4]`, `p = 5, q = 1` → `3`; `p = 5, q = 4` → `5`

**Constraints:** up to 10⁵ nodes; values are unique.
```

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

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil { // one on each side
		return root
	}
	if l != nil {
		return l
	}
	return r
}
// test
t := makeTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
assertEq(lowestCommonAncestor(t, findNode(t, 5), findNode(t, 1)).Val, 3)
assertEq(lowestCommonAncestor(t, findNode(t, 5), findNode(t, 4)).Val, 5)
```

## 102. Binary Tree Level Order Traversal | Medium
related: 103, 199, 637
tags: blind75, bfs

```problem
Return the values of a binary tree level by level, each level from left to right.

**Example:** `[3,9,20,null,null,15,7]` → `[[3],[9,20],[15,7]]`
```

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

```go
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for n := len(q); n > 0; n-- {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
// test
assertEq(levelOrder(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{3}, {9, 20}, {15, 7}})
assertEq(levelOrder(nil), [][]int{})
```

## 103. Binary Tree Zigzag Level Order Traversal | Medium
related: 102
tags: bfs

```problem
Return the values of a binary tree level by level, alternating direction: the first level left to right, the next right to left, and so on.

**Example:** `[3,9,20,null,null,15,7]` → `[[3],[20,9],[15,7]]`
```

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
assert Solution().zigzagLevelOrder(make_tree([1])) == [[1]]
assert Solution().zigzagLevelOrder(make_tree([])) == []
assert Solution().zigzagLevelOrder(make_tree([1, 2, 3, 4, None, None, 5])) == [[1], [3, 2], [4, 5]]
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

```go
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for n := len(q); n > 0; n-- {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(res)%2 == 1 { // odd levels read right to left
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		res = append(res, level)
	}
	return res
}
// test
assertEq(zigzagLevelOrder(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{3}, {20, 9}, {15, 7}})
assertEq(zigzagLevelOrder(makeTree(1)), [][]int{{1}})
assertEq(zigzagLevelOrder(makeTree()), [][]int{})
assertEq(zigzagLevelOrder(makeTree(1, 2, 3, 4, nil, nil, 5)), [][]int{{1}, {3, 2}, {4, 5}})
```

## 199. Binary Tree Right Side View | Medium
related: 102
tags: bfs

```problem
Imagine standing to the right of a binary tree. Return the values you can see, from top to bottom: the rightmost node of each level.

**Example:** `[1,2,3,null,5,null,4]` → `[1,3,4]`; `[1,2,3,4]` → `[1,3,4]`
```

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

```go
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		res = append(res, q[len(q)-1].Val) // last node of each level
		for n := len(q); n > 0; n-- {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}
// test
assertEq(rightSideView(makeTree(1, 2, 3, nil, 5, nil, 4)), []int{1, 3, 4})
assertEq(rightSideView(makeTree(1, 2, 3, 4)), []int{1, 3, 4})
```

## 1448. Count Good Nodes in Binary Tree | Medium
related: 98
tags: dfs

```problem
A node is "good" if no node on the path from the root to it has a greater value. Count the good nodes of a binary tree.

**Example:** `[3,1,4,3,null,1,5]` → `4`

**Constraints:** up to 10⁵ nodes.
```

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

```go
func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, mx int) int
	dfs = func(node *TreeNode, mx int) int {
		if node == nil {
			return 0
		}
		good := 0
		if node.Val >= mx { // nothing bigger on the path from the root
			good = 1
		}
		mx = max(mx, node.Val)
		return good + dfs(node.Left, mx) + dfs(node.Right, mx)
	}
	return dfs(root, root.Val)
}
// test
assertEq(goodNodes(makeTree(3, 1, 4, 3, nil, 1, 5)), 4)
assertEq(goodNodes(makeTree(3, 3, nil, 4, 2)), 3)
```

## 98. Validate Binary Search Tree | Medium
related: 230, 1448
tags: blind75, bst, boundaries

```problem
Decide whether a binary tree is a valid binary search tree: for every node, all values in its left subtree are strictly smaller and all values in its right subtree strictly larger.

**Example:** `[2,1,3]` → `true`; `[5,1,4,null,null,3,6]` → `false` (3 sits in 5's right subtree)

**Constraints:** values span the full 32-bit range.
```

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

```go
func isValidBST(root *TreeNode) bool {
	var ok func(node *TreeNode, lo, hi int) bool // every value must lie in (lo, hi)
	ok = func(node *TreeNode, lo, hi int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lo || node.Val >= hi {
			return false
		}
		return ok(node.Left, lo, node.Val) && ok(node.Right, node.Val, hi)
	}
	return ok(root, math.MinInt, math.MaxInt)
}
// test
assertEq(isValidBST(makeTree(2, 1, 3)), true)
assertEq(isValidBST(makeTree(5, 1, 6, nil, nil, 3, 7)), false)
```

## 230. Kth Smallest Element in a BST | Medium
related: 98, 173
tags: blind75, bst, inorder

```problem
Return the k-th smallest value (1-indexed) in a binary search tree.

**Example:** `[3,1,4,null,2], k = 1` → `1`; `[5,3,6,2,4,null,null,1], k = 3` → `3`

**Constraints:** follow-up: the BST is modified often and k-th smallest is queried often.
```

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

```go
func kthSmallest(root *TreeNode, k int) int {
	st, node := []*TreeNode{}, root // iterative in-order traversal
	for {
		for node != nil {
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]
		if k--; k == 0 {
			return node.Val
		}
		node = node.Right
	}
}
// test
assertEq(kthSmallest(makeTree(3, 1, 4, nil, 2), 1), 1)
assertEq(kthSmallest(makeTree(5, 3, 6, 2, 4, nil, nil, 1), 3), 3)
```

## 105. Construct Binary Tree from Preorder and Inorder Traversal | Medium
related: 106, 297
tags: blind75, divide-and-conquer

```problem
Rebuild a binary tree from its `preorder` and `inorder` traversals (all values are unique) and return the root.

**Example:** `preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]` → `[3,9,20,null,null,15,7]`

**Constraints:** up to 3000 nodes.
```

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
assert tree_vals(Solution().buildTree([-1], [-1])) == [-1]
assert tree_vals(Solution().buildTree([1, 2, 3], [3, 2, 1])) == [1, 2, None, 3]
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

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := map[int]int{}
	for i, v := range inorder {
		pos[v] = i
	}
	next := 0 // next preorder value = next subtree root
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode { // inorder range [lo, hi)
		if lo >= hi {
			return nil
		}
		root := &TreeNode{Val: preorder[next]}
		next++
		m := pos[root.Val]
		root.Left = build(lo, m) // the left subtree comes first in preorder
		root.Right = build(m+1, hi)
		return root
	}
	return build(0, len(inorder))
}
// test
assertEq(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}), []any{3, 9, 20, nil, nil, 15, 7})
assertEq(buildTree([]int{-1}, []int{-1}), []any{-1})
assertEq(buildTree([]int{1, 2, 3}, []int{3, 2, 1}), []any{1, 2, nil, 3})
```

## 124. Binary Tree Maximum Path Sum | Hard
related: 543, 687
tags: blind75, dfs, global-answer

```problem
A path is any sequence of nodes connected by parent-child edges, using each node at most once; it doesn't need to pass through the root and may bend at one node. Return the largest sum of node values over all non-empty paths. Values can be negative.

**Example:** `[-10,9,20,null,null,15,7]` → `42` (15 → 20 → 7)
```

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

```go
func maxPathSum(root *TreeNode) int {
	best := math.MinInt
	var gain func(*TreeNode) int // best downward path starting at node
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := max(gain(node.Left), 0), max(gain(node.Right), 0)
		best = max(best, node.Val+l+r)
		return node.Val + max(l, r)
	}
	gain(root)
	return best
}
// test
assertEq(maxPathSum(makeTree(1, 2, 3)), 6)
assertEq(maxPathSum(makeTree(-10, 9, 20, nil, nil, 15, 7)), 42)
assertEq(maxPathSum(makeTree(-3)), -3)
```

## 297. Serialize and Deserialize Binary Tree | Hard
related: 105, 449
tags: blind75, design, preorder

```problem
Design a `Codec` whose `serialize(root)` turns a binary tree into a string and `deserialize(data)` rebuilds exactly the same tree from that string. Any format works as long as the round trip is exact.

**Example:** `deserialize(serialize([1,2,3,null,null,4,5]))` → `[1,2,3,null,null,4,5]`

**Constraints:** up to 10⁴ nodes; values in [−1000, 1000].
```

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

```go
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// serialize writes the preorder with "#" for missing children.
func (this *Codec) serialize(root *TreeNode) string {
	out := []string{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			out = append(out, "#")
			return
		}
		out = append(out, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(out, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals, i := strings.Split(data, ","), 0
	var build func() *TreeNode
	build = func() *TreeNode {
		v := vals[i]
		i++
		if v == "#" {
			return nil
		}
		n, _ := strconv.Atoi(v)
		node := &TreeNode{Val: n}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}
// test
c := Constructor()
t := makeTree(1, 2, 3, nil, nil, 4, 5)
assertEq(c.serialize(t), "1,2,#,#,3,4,#,#,5,#,#")
assertEq(c.deserialize(c.serialize(t)), []any{1, 2, 3, nil, nil, 4, 5})
assertEq(c.serialize(c.deserialize(c.serialize(nil))), "#")
```

## 112. Path Sum | Easy
related: 113, 437
tags: dfs

```problem
Decide whether the binary tree has a root-to-leaf path whose node values add up to `targetSum`.

**Example:** `[5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22` → `true` (5 → 4 → 11 → 2)
```

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

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	rest := targetSum - root.Val
	return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}
// test
assertEq(hasPathSum(makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1), 22), true)
assertEq(hasPathSum(makeTree(1, 2, 3), 5), false)
```

## 113. Path Sum II | Medium
related: 112, 437, 257
tags: dfs, backtracking

```problem
Return every root-to-leaf path (as a list of node values) whose values add up to `targetSum`.

**Example:** `[5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22` → `[[5,4,11,2],[5,8,4,5]]`
```

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
assert sorted(Solution().pathSum(t, 22)) == sorted([[5, 4, 11, 2], [5, 8, 4, 5]])
assert Solution().pathSum(make_tree([1, 2, 3]), 5) == []
assert Solution().pathSum(make_tree([1, 2]), 0) == []
assert Solution().pathSum(make_tree([-2, None, -3]), -5) == [[-2, -3]]
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

```go
func pathSum(root *TreeNode, targetSum int) [][]int {
	res, path := [][]int{}, []int{}
	var dfs func(node *TreeNode, rest int)
	dfs = func(node *TreeNode, rest int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		rest -= node.Val
		if node.Left == nil && node.Right == nil && rest == 0 {
			res = append(res, append([]int{}, path...)) // copy: path keeps changing
		}
		dfs(node.Left, rest)
		dfs(node.Right, rest)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return res
}
// test
t := makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1)
assertUnordered(pathSum(t, 22), [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}})
assertEq(pathSum(makeTree(1, 2, 3), 5), [][]int{})
assertEq(pathSum(makeTree(1, 2), 0), [][]int{})
assertEq(pathSum(makeTree(-2, nil, -3), -5), [][]int{{-2, -3}})
```

## 437. Path Sum III | Medium
related: 560, 112, 113
tags: prefix-sum, dfs

```problem
Count the paths whose node values add up to `targetSum`. A path may start and end at any nodes but must go downwards (parent to child only).

**Example:** `[10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8` → `3`

**Constraints:** up to 1000 nodes; sums can exceed 32 bits.
```

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

```go
func pathSum(root *TreeNode, targetSum int) int {
	seen := map[int]int{0: 1} // prefix sums on the current root-to-node path
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur += node.Val
		cnt := seen[cur-targetSum]
		seen[cur]++
		cnt += dfs(node.Left, cur) + dfs(node.Right, cur)
		seen[cur]-- // backtrack
		return cnt
	}
	return dfs(root, 0)
}
// test
assertEq(pathSum(makeTree(10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1), 8), 3)
assertEq(pathSum(makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1), 22), 3)
```

## 114. Flatten Binary Tree to Linked List | Medium
related: 430
tags: morris, in-place

```problem
Flatten a binary tree in place into a "linked list" that follows preorder: each node's `right` points to the next node in preorder and every `left` is `null`.

**Example:** `[1,2,5,3,4,null,6]` → `1 → 2 → 3 → 4 → 5 → 6` along `right` pointers

**Constraints:** follow-up: O(1) extra space.
```

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
t = make_tree([0]); Solution().flatten(t)
assert tree_vals(t) == [0]
assert Solution().flatten(None) is None
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

```go
func flatten(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left != nil {
			p := cur.Left // rightmost node of the left subtree
			for p.Right != nil {
				p = p.Right
			}
			p.Right = cur.Right
			cur.Right, cur.Left = cur.Left, nil
		}
	}
}
// test
t := makeTree(1, 2, 5, 3, 4, nil, 6)
flatten(t)
assertEq(t, []any{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6})
t2 := makeTree(0)
flatten(t2)
assertEq(t2, []any{0})
flatten(nil)
```

## 116. Populating Next Right Pointers in Each Node | Medium
related: 117, 102
tags: bfs, perfect-tree

```problem
In a **perfect** binary tree (every level full), set each node's `next` pointer to the node immediately to its right on the same level, or `null` for the rightmost node.

**Example:** `[1,2,3,4,5,6,7]` → levels linked as `1→null`, `2→3→null`, `4→5→6→7→null`

**Constraints:** follow-up: O(1) extra space (recursion stack doesn't count).
```

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
assert Solution().connect(None) is None
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

```go
type Node struct { // provided by LeetCode
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	for leftmost := root; leftmost != nil && leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next { // walk this level with the links built so far
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
// test
n := make([]*Node, 8) // 1..7 used, heap layout
for i := range n {
	n[i] = &Node{Val: i}
}
for i := 1; i < 4; i++ {
	n[i].Left, n[i].Right = n[2*i], n[2*i+1]
}
connect(n[1])
assertEq([]bool{n[4].Next == n[5], n[5].Next == n[6], n[6].Next == n[7], n[7].Next == nil, n[2].Next == n[3]}, []bool{true, true, true, true, true})
assertEq(connect(nil) == nil, true)
```

## 108. Convert Sorted Array to Binary Search Tree | Easy
related: 109, 98
tags: divide-and-conquer

```problem
Turn a sorted array into a height-balanced binary search tree (subtree depths differ by at most 1 at every node). Any such tree is accepted.

**Example:** `[-10,-3,0,5,9]` → `[0,-3,9,-10,null,5]` (or another balanced BST)
```

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
# any height-balanced BST with the same in-order sequence is accepted
def balanced_bst(root, vals):
    def walk(n):  # (inorder values, height) or None if unbalanced
        if not n:
            return [], 0
        l, r = walk(n.left), walk(n.right)
        if l is None or r is None or abs(l[1] - r[1]) > 1:
            return None
        return l[0] + [n.val] + r[0], 1 + max(l[1], r[1])
    w = walk(root)
    return w is not None and w[0] == vals
for vals in ([-10, -3, 0, 5, 9], [1, 3], [0], [1, 2, 3, 4, 5, 6, 7, 8]):
    assert balanced_bst(Solution().sortedArrayToBST(vals), vals)
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

```go
func sortedArrayToBST(nums []int) *TreeNode {
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2 // the middle element keeps both sides balanced
		return &TreeNode{Val: nums[mid], Left: build(lo, mid-1), Right: build(mid+1, hi)}
	}
	return build(0, len(nums)-1)
}
// test
// any height-balanced BST with the same in-order sequence is accepted
var walk func(n *TreeNode) ([]int, int, bool) // in-order values, height, balanced
walk = func(n *TreeNode) ([]int, int, bool) {
	if n == nil {
		return []int{}, 0, true
	}
	lv, lh, lok := walk(n.Left)
	rv, rh, rok := walk(n.Right)
	vals := append(append(lv, n.Val), rv...)
	return vals, 1 + max(lh, rh), lok && rok && lh-rh <= 1 && rh-lh <= 1
}
for _, vals := range [][]int{{-10, -3, 0, 5, 9}, {1, 3}, {0}, {1, 2, 3, 4, 5, 6, 7, 8}} {
	got, _, balanced := walk(sortedArrayToBST(vals))
	assertEq(balanced, true)
	assertEq(got, vals)
}
```

## 173. Binary Search Tree Iterator | Medium
related: 230, 94
tags: design, inorder, stack

```problem
Design `BSTIterator(root)`, which walks a BST in sorted (in-order) order: `next()` returns the next smallest value and `hasNext()` says whether any remain.

**Example:** `[7,3,15,null,null,9,20]` → `next()` returns 3, 7, 9, 15, 20

**Constraints:** follow-up: O(1) amortised per call using O(h) memory.
```

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
it = BSTIterator(make_tree([1]))
assert it.hasNext() and it.next() == 1 and not it.hasNext()
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

```go
type BSTIterator struct {
	st []*TreeNode // the path of left children still to visit
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for ; node != nil; node = node.Left {
		this.st = append(this.st, node)
	}
}

func (this *BSTIterator) Next() int {
	node := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	this.pushLeft(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.st) > 0
}
// test
it := Constructor(makeTree(7, 3, 15, nil, nil, 9, 20))
out := []int{}
for i := 0; i < 10 && it.HasNext(); i++ {
	out = append(out, it.Next())
}
assertEq(out, []int{3, 7, 9, 15, 20})
it2 := Constructor(makeTree(1))
assertEq(it2.HasNext(), true)
assertEq(it2.Next(), 1)
assertEq(it2.HasNext(), false)
```

## 863. All Nodes Distance K in Binary Tree | Medium
related: 236, 102
tags: bfs, graph

```problem
Given a binary tree, a `target` node and an integer `k`, return the values of all nodes exactly `k` edges away from `target` (up through parents as well as down), in any order.

**Example:** `[3,5,1,6,2,0,8,null,null,7,4]`, `target = 5, k = 2` → `[7,4,1]`
```

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
assert Solution().distanceK(t, find_node(t, 5), 0) == [5]
assert sorted(Solution().distanceK(t, find_node(t, 7), 3)) == [3, 6]
t1 = make_tree([1])
assert Solution().distanceK(t1, t1, 3) == []
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

```go
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := map[*TreeNode]*TreeNode{}
	st := []*TreeNode{root}
	for len(st) > 0 { // record every node's parent, then BFS outward from target
		node := st[len(st)-1]
		st = st[:len(st)-1]
		for _, c := range []*TreeNode{node.Left, node.Right} {
			if c != nil {
				parent[c] = node
				st = append(st, c)
			}
		}
	}
	q, seen := []*TreeNode{target}, map[*TreeNode]bool{target: true}
	for ; k > 0 && len(q) > 0; k-- {
		next := []*TreeNode{}
		for _, node := range q {
			for _, nb := range []*TreeNode{node.Left, node.Right, parent[node]} {
				if nb != nil && !seen[nb] {
					seen[nb] = true
					next = append(next, nb)
				}
			}
		}
		q = next
	}
	res := []int{}
	for _, node := range q {
		res = append(res, node.Val)
	}
	return res
}
// test
t := makeTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
assertUnordered(distanceK(t, findNode(t, 5), 2), []int{1, 4, 7})
assertEq(distanceK(t, findNode(t, 5), 0), []int{5})
assertUnordered(distanceK(t, findNode(t, 7), 3), []int{3, 6})
t1 := makeTree(1)
assertEq(distanceK(t1, t1, 3), []int{})
```


## 94. Binary Tree Inorder Traversal | Easy
related: 230, 173, 144
tags: inorder, stack

```problem
Return the in-order traversal (left subtree, node, right subtree) of a binary tree.

**Example:** `[1,null,2,3]` → `[1,3,2]`

**Constraints:** follow-up: iteratively.
```

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

```go
func inorderTraversal(root *TreeNode) []int {
	res, st, node := []int{}, []*TreeNode{}, root
	for len(st) > 0 || node != nil {
		for node != nil { // go left as far as possible
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}
// test
assertEq(inorderTraversal(makeTree(1, nil, 2, 3)), []int{1, 3, 2})
assertEq(inorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{4, 2, 6, 5, 7, 1, 3, 9, 8})
```

## 111. Minimum Depth of Binary Tree | Easy
related: 104
tags: bfs

```problem
Return the minimum depth of a binary tree: the number of nodes on the shortest path from the root to a **leaf** (a node with no children).

**Example:** `[3,9,20,null,null,15,7]` → `2`; `[2,null,3,null,4,null,5,null,6]` → `5`
```

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

```go
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q, depth := []*TreeNode{root}, 1
	for { // BFS: the first leaf found is the shallowest
		next := []*TreeNode{}
		for _, n := range q {
			if n.Left == nil && n.Right == nil {
				return depth
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		q, depth = next, depth+1
	}
}
// test
assertEq(minDepth(makeTree(3, 9, 20, nil, nil, 15, 7)), 2)
assertEq(minDepth(makeTree(2, nil, 3, nil, 4, nil, 5, nil, 6)), 5)
```

## 106. Construct Binary Tree from Inorder and Postorder Traversal | Medium
related: 105
tags: divide-and-conquer

```problem
Rebuild a binary tree from its `inorder` and `postorder` traversals (all values are unique) and return the root.

**Example:** `inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]` → `[3,9,20,null,null,15,7]`
```

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
assert tree_vals(Solution().buildTree([-1], [-1])) == [-1]
assert tree_vals(Solution().buildTree([2, 1], [2, 1])) == [1, 2]
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

```go
func buildTree(inorder []int, postorder []int) *TreeNode {
	pos := map[int]int{}
	for i, v := range inorder {
		pos[v] = i
	}
	next := len(postorder) - 1 // postorder read backwards: root, right subtree, left subtree
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo >= hi {
			return nil
		}
		root := &TreeNode{Val: postorder[next]}
		next--
		m := pos[root.Val]
		root.Right = build(m+1, hi)
		root.Left = build(lo, m)
		return root
	}
	return build(0, len(inorder))
}
// test
assertEq(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}), []any{3, 9, 20, nil, nil, 15, 7})
assertEq(buildTree([]int{-1}, []int{-1}), []any{-1})
assertEq(buildTree([]int{2, 1}, []int{2, 1}), []any{1, 2})
```

## 109. Convert Sorted List to Binary Search Tree | Medium
related: 108
tags: inorder-simulation

```problem
Turn a sorted singly linked list into a height-balanced binary search tree. Any balanced BST with the same values is accepted.

**Example:** `-10→-3→0→5→9` → `[0,-3,9,-10,null,5]` (or another balanced BST)

**Constraints:** up to 2·10⁴ nodes.
```

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
# any height-balanced BST with the same in-order sequence is accepted
def balanced_bst(root, vals):
    def walk(n):  # (inorder values, height) or None if unbalanced
        if not n:
            return [], 0
        l, r = walk(n.left), walk(n.right)
        if l is None or r is None or abs(l[1] - r[1]) > 1:
            return None
        return l[0] + [n.val] + r[0], 1 + max(l[1], r[1])
    w = walk(root)
    return w is not None and w[0] == vals
for vals in ([-10, -3, 0, 5, 9], [1, 3], [0], [1, 2, 3, 4, 5, 6, 7, 8]):
    assert balanced_bst(Solution().sortedListToBST(make_list(vals)), vals)
assert Solution().sortedListToBST(None) is None
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

```go
func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	cur := head // consumed in order while building in-order
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2
		left := build(lo, mid-1)
		root := &TreeNode{Val: cur.Val, Left: left}
		cur = cur.Next
		root.Right = build(mid+1, hi)
		return root
	}
	return build(0, n-1)
}
// test
// any height-balanced BST with the same in-order sequence is accepted
var walk func(n *TreeNode) ([]int, int, bool) // in-order values, height, balanced
walk = func(n *TreeNode) ([]int, int, bool) {
	if n == nil {
		return []int{}, 0, true
	}
	lv, lh, lok := walk(n.Left)
	rv, rh, rok := walk(n.Right)
	vals := append(append(lv, n.Val), rv...)
	return vals, 1 + max(lh, rh), lok && rok && lh-rh <= 1 && rh-lh <= 1
}
for _, vals := range [][]int{{-10, -3, 0, 5, 9}, {1, 3}, {0}, {1, 2, 3, 4, 5, 6, 7, 8}} {
	got, _, balanced := walk(sortedListToBST(makeList(vals...)))
	assertEq(balanced, true)
	assertEq(got, vals)
}
assertEq(sortedListToBST(nil) == nil, true)
```

## 117. Populating Next Right Pointers in Each Node II | Medium
related: 116, 102
tags: bfs, level-linking

```problem
In any binary tree, set each node's `next` pointer to the node immediately to its right on the same level, or `null` if it is the rightmost.

**Example:** `[1,2,3,4,5,null,7]` → levels linked as `1→null`, `2→3→null`, `4→5→7→null`

**Constraints:** follow-up: O(1) extra space.
```

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
assert Solution().connect(None) is None
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

```go
type Node struct { // provided by LeetCode
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	for level := root; level != nil; {
		dummy := &Node{} // collects the next level through .Next
		tail := dummy
		for node := level; node != nil; node = node.Next {
			for _, child := range []*Node{node.Left, node.Right} {
				if child != nil {
					tail.Next = child
					tail = child
				}
			}
		}
		level = dummy.Next
	}
	return root
}
// test
n := map[int]*Node{}
for i := 1; i <= 7; i++ {
	n[i] = &Node{Val: i}
}
n[1].Left, n[1].Right = n[2], n[3]
n[2].Left, n[2].Right, n[3].Right = n[4], n[5], n[7]
connect(n[1])
assertEq([]bool{n[2].Next == n[3], n[4].Next == n[5], n[5].Next == n[7], n[7].Next == nil}, []bool{true, true, true, true})
assertEq(connect(nil) == nil, true)
```

## 129. Sum Root to Leaf Numbers | Medium
related: 112, 257
tags: dfs

```problem
Every node holds a digit 0–9, and each root-to-leaf path spells a number (root digit first). Return the sum of all these numbers.

**Example:** `[1,2,3]` → `25` (12 + 13); `[4,9,0,5,1]` → `1026` (495 + 491 + 40)
```

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

```go
func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur = cur*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return cur
		}
		return dfs(node.Left, cur) + dfs(node.Right, cur)
	}
	return dfs(root, 0)
}
// test
assertEq(sumNumbers(makeTree(1, 2, 3)), 25)
assertEq(sumNumbers(makeTree(4, 9, 0, 5, 1)), 1026)
```

## 257. Binary Tree Paths | Easy
related: 113, 129
tags: dfs, backtracking

```problem
Return every root-to-leaf path of a binary tree as a string of values joined by `"->"`, in any order.

**Example:** `[1,2,3,null,5]` → `["1->2->5","1->3"]`
```

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
assert sorted(Solution().binaryTreePaths(make_tree([1, 2, 3, None, 5]))) == sorted(["1->2->5", "1->3"])
assert Solution().binaryTreePaths(make_tree([1])) == ["1"]
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

```go
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		dfs(node.Left, path+"->")
		dfs(node.Right, path+"->")
	}
	dfs(root, "")
	return res
}
// test
assertUnordered(binaryTreePaths(makeTree(1, 2, 3, nil, 5)), []string{"1->2->5", "1->3"})
assertEq(binaryTreePaths(makeTree(1)), []string{"1"})
```

## 222. Count Complete Tree Nodes | Medium
related: 958
tags: binary-search, complete-tree

```problem
Count the nodes of a **complete** binary tree (every level full except possibly the last, whose nodes are packed to the left) in less than O(n) time.

**Example:** `[1,2,3,4,5,6]` → `6`

**Constraints:** up to 5·10⁴ nodes.
```

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

```go
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := 0, 0
	for l := root; l != nil; l = l.Left {
		hl++
	}
	for r := root; r != nil; r = r.Right {
		hr++
	}
	if hl == hr { // perfect tree
		return 1<<hl - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
// test
assertEq(countNodes(makeTree(1, 2, 3, 4, 5, 6)), 6)
vals := []any{}
for i := 1; i <= 15; i++ {
	vals = append(vals, i)
}
assertEq(countNodes(makeTree(vals...)), 15)
assertEq(countNodes(nil), 0)
```

## 337. House Robber III | Medium
related: 198, 213
tags: tree-dp

```problem
Houses form a binary tree; each node holds the money in that house. Robbing two houses joined by an edge (parent and child) sets off the alarm. Return the most money you can rob.

**Example:** `[3,2,3,null,3,null,1]` → `7` (3 + 3 + 1); `[3,4,5,1,3,null,1]` → `9` (4 + 5)
```

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

```go
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int) // (best if node is robbed, best if it is skipped)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lr, ls := dfs(node.Left)
		rr, rs := dfs(node.Right)
		return node.Val + ls + rs, max(lr, ls) + max(rr, rs)
	}
	robbed, skipped := dfs(root)
	return max(robbed, skipped)
}
// test
assertEq(rob(makeTree(3, 2, 3, nil, 3, nil, 1)), 7)
assertEq(rob(makeTree(3, 4, 5, 1, 3, nil, 1)), 9)
```

## 450. Delete Node in a BST | Medium
related: 701, 98
tags: bst

```problem
Delete the node with value `key` from a BST, if it exists, and return the root of a valid BST. Any valid result is accepted.

**Example:** `[5,3,6,2,4,null,7], key = 3` → `[5,4,6,2,null,null,7]`

**Constraints:** follow-up: O(height) time.
```

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

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		succ := root.Right // two children: take the in-order successor's value
		for succ.Left != nil {
			succ = succ.Left
		}
		root.Val = succ.Val
		root.Right = deleteNode(root.Right, succ.Val)
	}
	return root
}
// test
assertEq(deleteNode(makeTree(5, 3, 6, 2, 4, nil, 7), 3), []any{5, 4, 6, 2, nil, nil, 7})
assertEq(deleteNode(makeTree(5, 3, 6, 2, 4, nil, 7), 0), []any{5, 3, 6, 2, 4, nil, 7})
```

## 530. Minimum Absolute Difference in BST | Easy
related: 98, 230
tags: inorder

```problem
Return the minimum absolute difference between the values of any two different nodes of a BST.

**Example:** `[4,2,6,1,3]` → `1`
```

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

```go
func getMinimumDifference(root *TreeNode) int {
	best, prev, havePrev := math.MaxInt, 0, false
	var inorder func(*TreeNode) // in-order visits values sorted: compare neighbours
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if havePrev {
			best = min(best, node.Val-prev)
		}
		prev, havePrev = node.Val, true
		inorder(node.Right)
	}
	inorder(root)
	return best
}
// test
assertEq(getMinimumDifference(makeTree(4, 2, 6, 1, 3)), 1)
assertEq(getMinimumDifference(makeTree(1, 0, 48, nil, nil, 12, 49)), 1)
```

## 637. Average of Levels in Binary Tree | Easy
related: 102
tags: bfs

```problem
Return the average node value of each level of a binary tree, top to bottom.

**Example:** `[3,9,20,null,null,15,7]` → `[3.0,14.5,11.0]`
```

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
assert Solution().averageOfLevels(make_tree([1])) == [1.0]
assert Solution().averageOfLevels(make_tree([2147483647, 2147483647, 2147483647])) == [2147483647.0, 2147483647.0]
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

```go
func averageOfLevels(root *TreeNode) []float64 {
	res, q := []float64{}, []*TreeNode{root}
	for len(q) > 0 {
		sum, next := 0, []*TreeNode{}
		for _, n := range q {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(q)))
		q = next
	}
	return res
}
// test
assertEq(averageOfLevels(makeTree(3, 9, 20, nil, nil, 15, 7)), []float64{3, 14.5, 11})
assertEq(averageOfLevels(makeTree(1)), []float64{1})
assertEq(averageOfLevels(makeTree(2147483647, 2147483647, 2147483647)), []float64{2147483647, 2147483647})
```

## 662. Maximum Width of Binary Tree | Medium
related: 102, 958
tags: bfs, heap-indexing

```problem
The width of a level is the number of positions between its leftmost and rightmost non-null nodes, counting the null positions in between as if the tree were complete. Return the maximum width over all levels.

**Example:** `[1,3,2,5,3,null,9]` → `4`; `[1,3,2,5,null,null,9,6,null,7]` → `7`

**Constraints:** up to 3000 nodes; the answer fits in 32 bits.
```

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

```go
func widthOfBinaryTree(root *TreeNode) int {
	type item struct {
		node *TreeNode
		i    int // heap index within the level
	}
	best, q := 0, []item{{root, 0}}
	for len(q) > 0 {
		base := q[0].i
		best = max(best, q[len(q)-1].i-base+1)
		next := []item{}
		for _, it := range q {
			i := it.i - base // renumber from 0 so indices stay small
			if it.node.Left != nil {
				next = append(next, item{it.node.Left, 2 * i})
			}
			if it.node.Right != nil {
				next = append(next, item{it.node.Right, 2*i + 1})
			}
		}
		q = next
	}
	return best
}
// test
assertEq(widthOfBinaryTree(makeTree(1, 3, 2, 5, 3, nil, 9)), 4)
assertEq(widthOfBinaryTree(makeTree(1, 3, 2, 5, nil, nil, 9, 6, nil, 7)), 7)
```

## 687. Longest Univalue Path | Medium
related: 543, 124
tags: dfs, global-answer

```problem
Return the number of **edges** in the longest path in which every node has the same value. The path need not pass through the root.

**Example:** `[5,4,5,1,1,null,5]` → `2`; `[1,4,5,4,4,null,5]` → `2`
```

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

```go
func longestUnivaluePath(root *TreeNode) int {
	best := 0
	var arm func(*TreeNode) int // longest same-value path going down from node
	arm = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := arm(node.Left), arm(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			l++
		} else {
			l = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			r++
		} else {
			r = 0
		}
		best = max(best, l+r)
		return max(l, r)
	}
	arm(root)
	return best
}
// test
assertEq(longestUnivaluePath(makeTree(5, 4, 5, 1, 1, nil, 5)), 2)
assertEq(longestUnivaluePath(makeTree(1, 4, 5, 4, 4, nil, 5)), 2)
```

## 958. Check Completeness of a Binary Tree | Medium
related: 222, 662
tags: bfs

```problem
Decide whether a binary tree is complete: every level is full except possibly the last, whose nodes are all as far left as possible.

**Example:** `[1,2,3,4,5,6]` → `true`; `[1,2,3,4,5,null,7]` → `false`
```

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

```go
func isCompleteTree(root *TreeNode) bool {
	q, seenNil := []*TreeNode{root}, false
	for len(q) > 0 { // BFS including nils: no node may come after a nil
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNil = true
			continue
		}
		if seenNil {
			return false
		}
		q = append(q, node.Left, node.Right)
	}
	return true
}
// test
assertEq(isCompleteTree(makeTree(1, 2, 3, 4, 5, 6)), true)
assertEq(isCompleteTree(makeTree(1, 2, 3, 4, 5, nil, 7)), false)
```

## 987. Vertical Order Traversal of a Binary Tree | Hard
related: 102, 199
tags: dfs, sort

```problem
Put the root at (row 0, column 0); a node at (r, c) has children at (r + 1, c − 1) and (r + 1, c + 1). Return the values column by column, left to right. Within a column, list them top to bottom, and sort nodes sharing the same row and column by value.

**Example:** `[1,2,3,4,6,5,7]` → `[[4],[2],[1,5,6],[3],[7]]`
```

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

```go
func verticalTraversal(root *TreeNode) [][]int {
	cells := [][3]int{} // (col, row, val)
	var dfs func(node *TreeNode, r, c int)
	dfs = func(node *TreeNode, r, c int) {
		if node != nil {
			cells = append(cells, [3]int{c, r, node.Val})
			dfs(node.Left, r+1, c-1)
			dfs(node.Right, r+1, c+1)
		}
	}
	dfs(root, 0, 0)
	sort.Slice(cells, func(i, j int) bool {
		a, b := cells[i], cells[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[2] < b[2]
	})
	res := [][]int{}
	for i, cell := range cells {
		if i == 0 || cell[0] != cells[i-1][0] {
			res = append(res, []int{})
		}
		res[len(res)-1] = append(res[len(res)-1], cell[2])
	}
	return res
}
// test
assertEq(verticalTraversal(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{9}, {3, 15}, {20}, {7}})
assertEq(verticalTraversal(makeTree(1, 2, 3, 4, 6, 5, 7)), [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}})
```

## 872. Leaf-Similar Trees | Easy
related: 100, 257
tags: dfs

```problem
Decide whether two binary trees have the same leaf value sequence (their leaves read left to right).

**Example:** two trees whose leaves both read `6,7,4,9,8` → `true`
```

Collect each tree's leaves left-to-right with a DFS and compare the sequences.

**Complexity:** O(n + m).

```python
class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        def leaves(node):
            if node:
                if not node.left and not node.right:
                    yield node.val
                yield from leaves(node.left)
                yield from leaves(node.right)

        return list(leaves(root1)) == list(leaves(root2))
# test
a = make_tree([3, 5, 1, 6, 2, 9, 8, None, None, 7, 4])
b = make_tree([3, 5, 1, 6, 7, 4, 2, None, None, None, None, None, None, 9, 8])
assert Solution().leafSimilar(a, b) is True
assert Solution().leafSimilar(make_tree([1, 2, 3]), make_tree([1, 3, 2])) is False
assert Solution().leafSimilar(make_tree([1]), make_tree([2, 1])) is True
```

```java
class Solution {
    public boolean leafSimilar(TreeNode root1, TreeNode root2) {
        List<Integer> a = new ArrayList<>(), b = new ArrayList<>();
        leaves(root1, a);
        leaves(root2, b);
        return a.equals(b);
    }

    private void leaves(TreeNode n, List<Integer> out) {
        if (n == null) return;
        if (n.left == null && n.right == null) out.add(n.val);
        leaves(n.left, out);
        leaves(n.right, out);
    }
}
// test
check(new Solution().leafSimilar(TreeNode.of(1, 2, 3), TreeNode.of(1, 3, 2)), false);
check(new Solution().leafSimilar(TreeNode.of(1), TreeNode.of(2, 1)), true);
```

```go
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves func(node *TreeNode, out []int) []int
	leaves = func(node *TreeNode, out []int) []int {
		if node == nil {
			return out
		}
		if node.Left == nil && node.Right == nil {
			return append(out, node.Val)
		}
		return leaves(node.Right, leaves(node.Left, out))
	}
	a, b := leaves(root1, nil), leaves(root2, nil)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
// test
a := makeTree(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4)
b := makeTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)
assertEq(leafSimilar(a, b), true)
assertEq(leafSimilar(makeTree(1, 2, 3), makeTree(1, 3, 2)), false)
assertEq(leafSimilar(makeTree(1), makeTree(2, 1)), true)
```

## 1372. Longest ZigZag Path in a Binary Tree | Medium
related: 124, 543
tags: tree-dp

```problem
A zigzag path starts at any node, picks a direction (left or right), and alternates direction at every step down. Return the length (in edges) of the longest zigzag path.

**Example:** `[1,1,1,null,1,null,null,1,1,null,1]` → `4`

**Constraints:** up to 5·10⁴ nodes.
```

For every node compute two values: the longest zigzag **starting here going left** (= 1 + the left child's "going right" value) and **going right** (= 1 + the right child's "going left"). A missing child contributes −1 so the step to it counts 0. The answer is the maximum over all nodes.

**Complexity:** O(n).

```python
class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        best = 0

        def dfs(node):                       # -> (zigzag length going left, going right), in edges
            nonlocal best
            if not node:
                return -1, -1
            left = dfs(node.left)[1] + 1
            right = dfs(node.right)[0] + 1
            best = max(best, left, right)
            return left, right

        dfs(root)
        return best
# test
assert Solution().longestZigZag(make_tree([1, None, 1, 1, 1, None, None, 1, 1, None, 1, None, None, None, 1])) == 3
assert Solution().longestZigZag(make_tree([1, 1, 1, None, 1, None, None, 1, 1, None, 1])) == 4
assert Solution().longestZigZag(make_tree([1])) == 0
```

```java
class Solution {
    private int best = 0;

    public int longestZigZag(TreeNode root) {
        dfs(root);
        return best;
    }

    private int[] dfs(TreeNode n) {          // {going left, going right}
        if (n == null) return new int[]{-1, -1};
        int left = dfs(n.left)[1] + 1, right = dfs(n.right)[0] + 1;
        best = Math.max(best, Math.max(left, right));
        return new int[]{left, right};
    }
}
// test
check(new Solution().longestZigZag(TreeNode.of(1, 1, 1, null, 1, null, null, 1, 1, null, 1)), 4);
```

```go
func longestZigZag(root *TreeNode) int {
	best := 0
	var dfs func(*TreeNode) (int, int) // zigzag length (in edges) starting with a left / right step
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return -1, -1
		}
		_, lr := dfs(node.Left)
		rl, _ := dfs(node.Right)
		left, right := lr+1, rl+1
		best = max(best, left, right)
		return left, right
	}
	dfs(root)
	return best
}
// test
assertEq(longestZigZag(makeTree(1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1)), 3)
assertEq(longestZigZag(makeTree(1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1)), 4)
assertEq(longestZigZag(makeTree(1)), 0)
```

## 1161. Maximum Level Sum of a Binary Tree | Medium
related: 102, 637
tags: bfs

```problem
The root is on level 1, its children on level 2, and so on. Return the smallest level whose node values have the maximal sum.

**Example:** `[1,7,0,7,-8,null,null]` → `2` (level sums 1, 7, −1)
```

Level-order BFS, summing each level; keep the first (smallest) level with the strictly largest sum. Values can be negative, so start from −∞.

**Complexity:** O(n).

```python
class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        best, best_level, level, q = float("-inf"), 0, 0, deque([root])
        while q:
            level += 1
            total = 0
            for _ in range(len(q)):
                n = q.popleft()
                total += n.val
                q.extend(c for c in (n.left, n.right) if c)
            if total > best:
                best, best_level = total, level
        return best_level
# test
assert Solution().maxLevelSum(make_tree([1, 7, 0, 7, -8, None, None])) == 2
assert Solution().maxLevelSum(make_tree([989, None, 10250, 98693, -89388, None, None, None, -32127])) == 2
assert Solution().maxLevelSum(make_tree([-1, -2, -3])) == 1
```

```java
class Solution {
    public int maxLevelSum(TreeNode root) {
        long best = Long.MIN_VALUE;
        int bestLevel = 0, level = 0;
        Deque<TreeNode> q = new ArrayDeque<>(List.of(root));
        while (!q.isEmpty()) {
            level++;
            long total = 0;
            for (int i = q.size(); i > 0; i--) {
                TreeNode n = q.poll();
                total += n.val;
                if (n.left != null) q.add(n.left);
                if (n.right != null) q.add(n.right);
            }
            if (total > best) { best = total; bestLevel = level; }
        }
        return bestLevel;
    }
}
// test
check(new Solution().maxLevelSum(TreeNode.of(1, 7, 0, 7, -8, null, null)), 2);
check(new Solution().maxLevelSum(TreeNode.of(-1, -2, -3)), 1);
```

```go
func maxLevelSum(root *TreeNode) int {
	best, bestLevel := math.MinInt, 0
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		total, next := 0, []*TreeNode{}
		for _, n := range q {
			total += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if total > best { // strictly greater: the smallest level wins ties
			best, bestLevel = total, level
		}
		q = next
	}
	return bestLevel
}
// test
assertEq(maxLevelSum(makeTree(1, 7, 0, 7, -8, nil, nil)), 2)
assertEq(maxLevelSum(makeTree(989, nil, 10250, 98693, -89388, nil, nil, nil, -32127)), 2)
assertEq(maxLevelSum(makeTree(-1, -2, -3)), 1)
```

## 700. Search in a Binary Search Tree | Easy
related: 701, 98
tags: bst

```problem
Find the node with value `val` in a BST and return the subtree rooted there, or `null` if it doesn't exist.

**Example:** `[4,2,7,1,3], val = 2` → `[2,1,3]`
```

Walk down from the root: go left if the target is smaller, right if larger. No recursion needed.

**Complexity:** O(h).

```python
class Solution:
    def searchBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        while root and root.val != val:
            root = root.left if val < root.val else root.right
        return root
# test
assert tree_vals(Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 2)) == [2, 1, 3]
assert Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 5) is None
assert tree_vals(Solution().searchBST(make_tree([4, 2, 7, 1, 3]), 7)) == [7]
```

```java
class Solution {
    public TreeNode searchBST(TreeNode root, int val) {
        while (root != null && root.val != val) root = val < root.val ? root.left : root.right;
        return root;
    }
}
// test
check(new Solution().searchBST(TreeNode.of(4, 2, 7, 1, 3), 2), "[2, 1, 3]");
check(new Solution().searchBST(TreeNode.of(4, 2, 7, 1, 3), 5), null);
```

```go
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
// test
assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 2), []any{2, 1, 3})
assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 5) == nil, true)
assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 7), []any{7})
```

## 285. Inorder Successor in BST | Medium
related: 230, 173, 98
tags: bst, premium

```problem
Given a BST and one of its nodes `p`, return the in-order successor of `p` (the node with the smallest value greater than `p.val`), or `null` if `p` holds the largest value.

**Example:** `[5,3,6,2,4,null,null,1]`, `p = 6` → `null`; `p = 4` → `5`
```

**Problem (Premium):** given a BST and one of its nodes `p`, return the node with the smallest key greater than `p.val`, or null.

No in-order traversal needed: walk down from the root. Whenever `p.val < node.val`, this node is a **candidate** (it's bigger than p) — remember it and go left looking for a smaller one; otherwise go right.

**Complexity:** O(h) time, O(1) space.

```python
class Solution:
    def inorderSuccessor(self, root: TreeNode, p: TreeNode) -> Optional[TreeNode]:
        succ = None
        while root:
            if p.val < root.val:
                succ, root = root, root.left
            else:
                root = root.right
        return succ
# test
t = make_tree([2, 1, 3])
assert Solution().inorderSuccessor(t, find_node(t, 1)).val == 2
t = make_tree([5, 3, 6, 2, 4, None, None, 1])
assert Solution().inorderSuccessor(t, find_node(t, 6)) is None
assert Solution().inorderSuccessor(t, find_node(t, 4)).val == 5
assert Solution().inorderSuccessor(t, find_node(t, 1)).val == 2
assert Solution().inorderSuccessor(t, find_node(t, 3)).val == 4
```

```java
class Solution {
    public TreeNode inorderSuccessor(TreeNode root, TreeNode p) {
        TreeNode succ = null;
        while (root != null) {
            if (p.val < root.val) { succ = root; root = root.left; }
            else root = root.right;
        }
        return succ;
    }
}
// test
TreeNode t = TreeNode.of(5, 3, 6, 2, 4, null, null, 1);
check(new Solution().inorderSuccessor(t, TreeNode.find(t, 4)).val, 5);
check(new Solution().inorderSuccessor(t, TreeNode.find(t, 6)), null);
```

```go
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var succ *TreeNode
	for root != nil {
		if p.Val < root.Val { // root is a candidate; look for a smaller one on the left
			succ, root = root, root.Left
		} else {
			root = root.Right
		}
	}
	return succ
}
// test
t := makeTree(2, 1, 3)
assertEq(inorderSuccessor(t, findNode(t, 1)).Val, 2)
t2 := makeTree(5, 3, 6, 2, 4, nil, nil, 1)
assertEq(inorderSuccessor(t2, findNode(t2, 6)) == nil, true)
assertEq(inorderSuccessor(t2, findNode(t2, 4)).Val, 5)
assertEq(inorderSuccessor(t2, findNode(t2, 1)).Val, 2)
assertEq(inorderSuccessor(t2, findNode(t2, 3)).Val, 4)
```

## 427. Construct Quad Tree | Medium
related: 108, 54
tags: divide-and-conquer

```problem
Represent an n×n grid of 0s and 1s (n a power of 2) as a quad tree. If a region is uniform, it becomes a leaf holding that value; otherwise it becomes an internal node with four children for its top-left, top-right, bottom-left and bottom-right quarters. Return the root.

**Example:** `[[0,1],[1,0]]` → an internal root with four leaves `0, 1, 1, 0`
```

Recursively build the four quadrants of an `n × n` square. If all four children are leaves with the same value, collapse them into one leaf; otherwise return an internal node holding the four children. (Internal nodes' `val` can be anything.)

```viz
1 1 0 0      TL = all 1 -> leaf(1)    TR = all 0 -> leaf(0)
1 1 0 0
1 1 1 1      BL = all 1 -> leaf(1)    BR = all 1 -> leaf(1)
1 1 1 1      children differ -> root is internal
```

**Complexity:** O(n²) time — every cell is a base case exactly once and each merge is O(1) (the recursion tree has about 4/3·n² nodes); O(log n) recursion depth.

```python
class Node:  # provided by LeetCode
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight


class Solution:
    def construct(self, grid: List[List[int]]) -> 'Node':
        def build(r, c, n):
            if n == 1:
                return Node(grid[r][c] == 1, True, None, None, None, None)
            h = n // 2
            kids = [build(r, c, h), build(r, c + h, h), build(r + h, c, h), build(r + h, c + h, h)]
            if all(k.isLeaf for k in kids) and len({k.val for k in kids}) == 1:
                return Node(kids[0].val, True, None, None, None, None)
            return Node(True, False, *kids)

        return build(0, 0, len(grid))
# test
def quad(node: "Node"):  # level order: "*" = internal node, "0"/"1" = leaf value
    out, q = [], deque([node])
    while q:
        n = q.popleft()
        out.append(("1" if n.val else "0") if n.isLeaf else "*")
        if not n.isLeaf:
            q.extend([n.topLeft, n.topRight, n.bottomLeft, n.bottomRight])
    return "".join(out)
assert quad(Solution().construct([[0, 1], [1, 0]])) == "*0110"
assert quad(Solution().construct([[1]])) == "1"
assert quad(Solution().construct([[1, 1], [1, 1]])) == "1"
assert quad(Solution().construct([[1, 1, 0, 0], [1, 1, 0, 0], [1, 1, 1, 1], [1, 1, 1, 1]])) == "*1011"
assert quad(Solution().construct([[1, 1, 0, 0], [1, 0, 0, 0], [1, 1, 1, 1], [1, 1, 1, 1]])) == "**0111110"
```

```java
class Node { // provided by LeetCode
    public boolean val, isLeaf;
    public Node topLeft, topRight, bottomLeft, bottomRight;
    public Node() {}
    public Node(boolean val, boolean isLeaf) { this.val = val; this.isLeaf = isLeaf; }
    public Node(boolean val, boolean isLeaf, Node topLeft, Node topRight, Node bottomLeft, Node bottomRight) {
        this.val = val; this.isLeaf = isLeaf;
        this.topLeft = topLeft; this.topRight = topRight; this.bottomLeft = bottomLeft; this.bottomRight = bottomRight;
    }
}

class Solution {
    public Node construct(int[][] grid) { return build(grid, 0, 0, grid.length); }

    private Node build(int[][] g, int r, int c, int n) {
        if (n == 1) return new Node(g[r][c] == 1, true);
        int h = n / 2;
        Node tl = build(g, r, c, h), tr = build(g, r, c + h, h), bl = build(g, r + h, c, h), br = build(g, r + h, c + h, h);
        if (tl.isLeaf && tr.isLeaf && bl.isLeaf && br.isLeaf && tl.val == tr.val && tr.val == bl.val && bl.val == br.val)
            return new Node(tl.val, true);
        return new Node(true, false, tl, tr, bl, br);
    }
}
// test
StringBuilder sb = new StringBuilder();
Deque<Node> q = new ArrayDeque<>(List.of(new Solution().construct(new int[][]{{1, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})));
while (!q.isEmpty()) {
    Node n = q.poll();
    sb.append(n.isLeaf ? (n.val ? "1" : "0") : "*");
    if (!n.isLeaf) q.addAll(List.of(n.topLeft, n.topRight, n.bottomLeft, n.bottomRight));
}
check(sb.toString(), "*1011");
```

```go
type Node struct { // provided by LeetCode
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var build func(r, c, n int) *Node
	build = func(r, c, n int) *Node {
		if n == 1 {
			return &Node{Val: grid[r][c] == 1, IsLeaf: true}
		}
		h := n / 2
		kids := []*Node{build(r, c, h), build(r, c+h, h), build(r+h, c, h), build(r+h, c+h, h)}
		merge := true // four leaves with the same value become one leaf
		for _, k := range kids {
			merge = merge && k.IsLeaf && k.Val == kids[0].Val
		}
		if merge {
			return &Node{Val: kids[0].Val, IsLeaf: true}
		}
		return &Node{Val: true, TopLeft: kids[0], TopRight: kids[1], BottomLeft: kids[2], BottomRight: kids[3]}
	}
	return build(0, 0, len(grid))
}
// test
quad := func(node *Node) string { // level order: "*" = internal node, "0"/"1" = leaf value
	out, q := "", []*Node{node}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			return out + "?"
		}
		switch {
		case !n.IsLeaf:
			out += "*"
			q = append(q, n.TopLeft, n.TopRight, n.BottomLeft, n.BottomRight)
		case n.Val:
			out += "1"
		default:
			out += "0"
		}
	}
	return out
}
assertEq(quad(construct([][]int{{0, 1}, {1, 0}})), "*0110")
assertEq(quad(construct([][]int{{1}})), "1")
assertEq(quad(construct([][]int{{1, 1}, {1, 1}})), "1")
assertEq(quad(construct([][]int{{1, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})), "*1011")
assertEq(quad(construct([][]int{{1, 1, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})), "**0111110")
```

## 144. Binary Tree Preorder Traversal | Easy
related: 94, 145, 589
tags: preorder, stack

```problem
Return the pre-order traversal (node, left subtree, right subtree) of a binary tree.

**Example:** `[1,null,2,3]` → `[1,2,3]`

**Constraints:** follow-up: iteratively.
```

Iterative preorder: pop a node, record it, push its **right** child then its **left** child (so the left is processed first).

**Complexity:** O(n) time, O(h) stack.

```python
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        out, st = [], [root] if root else []
        while st:
            node = st.pop()
            out.append(node.val)
            if node.right:
                st.append(node.right)
            if node.left:
                st.append(node.left)
        return out
# test
assert Solution().preorderTraversal(make_tree([1, None, 2, 3])) == [1, 2, 3]
assert Solution().preorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [1, 2, 4, 5, 6, 7, 3, 8, 9]
assert Solution().preorderTraversal(make_tree([])) == []
```

```java
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> out = new ArrayList<>();
        Deque<TreeNode> st = new ArrayDeque<>();
        if (root != null) st.push(root);
        while (!st.isEmpty()) {
            TreeNode n = st.pop();
            out.add(n.val);
            if (n.right != null) st.push(n.right);
            if (n.left != null) st.push(n.left);
        }
        return out;
    }
}
// test
check(new Solution().preorderTraversal(TreeNode.of(1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9)), List.of(1, 2, 4, 5, 6, 7, 3, 8, 9));
```

```go
func preorderTraversal(root *TreeNode) []int {
	out, st := []int{}, []*TreeNode{}
	if root != nil {
		st = append(st, root)
	}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, node.Val)
		if node.Right != nil { // push right first so left is visited first
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
	return out
}
// test
assertEq(preorderTraversal(makeTree(1, nil, 2, 3)), []int{1, 2, 3})
assertEq(preorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{1, 2, 4, 5, 6, 7, 3, 8, 9})
assertEq(preorderTraversal(makeTree()), []int{})
```

## 145. Binary Tree Postorder Traversal | Easy
related: 94, 144
tags: stack, reverse-trick

```problem
Return the post-order traversal (left subtree, right subtree, node) of a binary tree.

**Example:** `[1,null,2,3]` → `[3,2,1]`

**Constraints:** follow-up: iteratively.
```

Postorder (left, right, root) is the **reverse** of a modified preorder (root, right, left): run preorder pushing left before right, then reverse the output.

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        out, st = [], [root] if root else []
        while st:
            node = st.pop()
            out.append(node.val)
            if node.left:
                st.append(node.left)
            if node.right:
                st.append(node.right)
        return out[::-1]
# test
assert Solution().postorderTraversal(make_tree([1, None, 2, 3])) == [3, 2, 1]
assert Solution().postorderTraversal(make_tree([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9])) == [4, 6, 7, 5, 2, 9, 8, 3, 1]
assert Solution().postorderTraversal(make_tree([])) == []
```

```java
class Solution {
    public List<Integer> postorderTraversal(TreeNode root) {
        LinkedList<Integer> out = new LinkedList<>();
        Deque<TreeNode> st = new ArrayDeque<>();
        if (root != null) st.push(root);
        while (!st.isEmpty()) {
            TreeNode n = st.pop();
            out.addFirst(n.val);                 // prepend = reverse of root-right-left
            if (n.left != null) st.push(n.left);
            if (n.right != null) st.push(n.right);
        }
        return out;
    }
}
// test
check(new Solution().postorderTraversal(TreeNode.of(1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9)), List.of(4, 6, 7, 5, 2, 9, 8, 3, 1));
```

```go
func postorderTraversal(root *TreeNode) []int {
	out, st := []int{}, []*TreeNode{}
	if root != nil {
		st = append(st, root)
	}
	for len(st) > 0 { // root-right-left, then reversed = left-right-root
		node := st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, node.Val)
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
// test
assertEq(postorderTraversal(makeTree(1, nil, 2, 3)), []int{3, 2, 1})
assertEq(postorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{4, 6, 7, 5, 2, 9, 8, 3, 1})
assertEq(postorderTraversal(makeTree()), []int{})
```

## 701. Insert into a Binary Search Tree | Medium
related: 700, 450, 98
tags: bst

```problem
Insert `val` (not already present) into a BST and return the root of a valid BST. Any valid result is accepted.

**Example:** `[4,2,7,1,3], val = 5` → `[4,2,7,1,3,5]`
```

Walk down as in a search until you fall off the tree; attach the new node there as a leaf. (Any valid BST is accepted; leaf insertion is the simplest.)

**Complexity:** O(h).

```python
class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        node = TreeNode(val)
        if not root:
            return node
        cur = root
        while True:
            if val < cur.val:
                if not cur.left:
                    cur.left = node
                    return root
                cur = cur.left
            else:
                if not cur.right:
                    cur.right = node
                    return root
                cur = cur.right
# test
def inorder(n):
    return inorder(n.left) + [n.val] + inorder(n.right) if n else []
for vals, v in (([4, 2, 7, 1, 3], 5), ([40, 20, 60, 10, 30, 50, 70], 25), ([], 5), ([4, 2, 7, 1, 3, None, None, None, None, None, None], 0)):
    got = inorder(Solution().insertIntoBST(make_tree(vals), v))
    assert got == sorted([x for x in vals if x is not None] + [v])     # still a BST holding every value
```

```java
class Solution {
    public TreeNode insertIntoBST(TreeNode root, int val) {
        if (root == null) return new TreeNode(val);
        TreeNode cur = root;
        while (true) {
            if (val < cur.val) {
                if (cur.left == null) { cur.left = new TreeNode(val); return root; }
                cur = cur.left;
            } else {
                if (cur.right == null) { cur.right = new TreeNode(val); return root; }
                cur = cur.right;
            }
        }
    }
}
// test
check(new Solution().insertIntoBST(TreeNode.of(4, 2, 7, 1, 3), 5), "[4, 2, 7, 1, 3, 5]");
```

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	for cur := root; ; {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
}
// test
var inorder func(n *TreeNode) []int
inorder = func(n *TreeNode) []int {
	if n == nil {
		return []int{}
	}
	return append(append(inorder(n.Left), n.Val), inorder(n.Right)...)
}
for _, tc := range []struct {
	vals []any
	v    int
}{{[]any{4, 2, 7, 1, 3}, 5}, {[]any{40, 20, 60, 10, 30, 50, 70}, 25}, {[]any{}, 5}, {[]any{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}, 0}} {
	want := []int{tc.v}
	for _, x := range tc.vals {
		if x != nil {
			want = append(want, x.(int))
		}
	}
	sort.Ints(want)
	assertEq(inorder(insertIntoBST(makeTree(tc.vals...), tc.v)), want) // still a BST holding every value
}
```

## 1325. Delete Leaves With a Given Value | Medium
related: 814, 450
tags: postorder

```problem
Delete every leaf whose value is `target`. If that turns a parent into a leaf with value `target`, delete it too, and keep going until no such leaf remains.

**Example:** `[1,2,3,2,null,2,4], target = 2` → `[1,null,3,null,4]`
```

Post-order: clean both subtrees first, **then** decide about the node — after its children are deleted it may have become a target leaf itself, which cascades upward naturally.

**Complexity:** O(n).

```python
class Solution:
    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        if not root:
            return None
        root.left = self.removeLeafNodes(root.left, target)
        root.right = self.removeLeafNodes(root.right, target)
        if not root.left and not root.right and root.val == target:
            return None
        return root
# test
assert tree_vals(Solution().removeLeafNodes(make_tree([1, 2, 3, 2, None, 2, 4]), 2)) == [1, None, 3, None, 4]
assert tree_vals(Solution().removeLeafNodes(make_tree([1, 3, 3, 3, 2]), 3)) == [1, 3, None, None, 2]
assert tree_vals(Solution().removeLeafNodes(make_tree([1, 2, None, 2, None, 2]), 2)) == [1]
assert Solution().removeLeafNodes(make_tree([1, 1, 1]), 1) is None
```

```java
class Solution {
    public TreeNode removeLeafNodes(TreeNode root, int target) {
        if (root == null) return null;
        root.left = removeLeafNodes(root.left, target);
        root.right = removeLeafNodes(root.right, target);
        return root.left == null && root.right == null && root.val == target ? null : root;
    }
}
// test
check(new Solution().removeLeafNodes(TreeNode.of(1, 2, 3, 2, null, 2, 4), 2), "[1, null, 3, null, 4]");
check(new Solution().removeLeafNodes(TreeNode.of(1, 1, 1), 1), null);
```

```go
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target) // children first: a parent may become a leaf
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
// test
assertEq(removeLeafNodes(makeTree(1, 2, 3, 2, nil, 2, 4), 2), []any{1, nil, 3, nil, 4})
assertEq(removeLeafNodes(makeTree(1, 3, 3, 3, 2), 3), []any{1, 3, nil, nil, 2})
assertEq(removeLeafNodes(makeTree(1, 2, nil, 2, nil, 2), 2), []any{1})
assertEq(removeLeafNodes(makeTree(1, 1, 1), 1) == nil, true)
```
