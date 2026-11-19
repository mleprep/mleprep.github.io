// 987. Vertical Order Traversal of a Binary Tree (Hard)
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// Run (Java 11+): java P0987VerticalOrderTraversalOfABinaryTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0987VerticalOrderTraversalOfABinaryTree {
    public static void main(String[] args) {
        check(new Solution().verticalTraversal(TreeNode.of(1, 2, 3, 4, 6, 5, 7)), "[[4], [2], [1, 5, 6], [3], [7]]");
        if (fails > 0) { System.out.println(fails + " test(s) FAILED"); System.exit(1); }
        System.out.println("OK");
    }

    static int fails = 0;

    static void check(Object got, Object expected) {
        String g = str(got), e = str(expected);
        if (!g.equals(e)) { fails++; System.out.println("FAIL: got " + g + ", expected " + e); }
    }

    static String str(Object o) {
        if (o == null) return "null";
        if (o instanceof int[]) return Arrays.toString((int[]) o);
        if (o instanceof long[]) return Arrays.toString((long[]) o);
        if (o instanceof double[]) return Arrays.toString((double[]) o);
        if (o instanceof char[]) return Arrays.toString((char[]) o);
        if (o instanceof boolean[]) return Arrays.toString((boolean[]) o);
        if (o instanceof Object[]) return Arrays.deepToString((Object[]) o);
        return String.valueOf(o);
    }
}

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

// ---- LeetCode's TreeNode + local test helpers ----
class TreeNode {
    int val; TreeNode left, right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) { this.val = val; this.left = left; this.right = right; }
    /** Build from LeetCode level-order, e.g. TreeNode.of(3, 9, 20, null, null, 15, 7). */
    static TreeNode of(Integer... v) {
        if (v.length == 0 || v[0] == null) return null;
        TreeNode root = new TreeNode(v[0]);
        Deque<TreeNode> q = new ArrayDeque<>(List.of(root));
        int i = 1;
        while (!q.isEmpty() && i < v.length) {
            TreeNode n = q.poll();
            if (i < v.length && v[i] != null) q.add(n.left = new TreeNode(v[i])); i++;
            if (i < v.length && v[i] != null) q.add(n.right = new TreeNode(v[i])); i++;
        }
        return root;
    }
    static TreeNode find(TreeNode r, int x) {
        if (r == null || r.val == x) return r;
        TreeNode l = find(r.left, x);
        return l != null ? l : find(r.right, x);
    }
    public String toString() {
        List<String> out = new ArrayList<>();
        Deque<TreeNode> q = new LinkedList<>(); q.add(this);
        while (!q.isEmpty()) {
            TreeNode n = q.poll();
            if (n == null) { out.add("null"); continue; }
            out.add(String.valueOf(n.val)); q.add(n.left); q.add(n.right);
        }
        while (!out.isEmpty() && out.get(out.size() - 1).equals("null")) out.remove(out.size() - 1);
        return "[" + String.join(", ", out) + "]";
    }
}
