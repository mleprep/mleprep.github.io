// 1161. Maximum Level Sum of a Binary Tree (Medium)
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Run (Java 11+): java P1161MaximumLevelSumOfABinaryTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1161MaximumLevelSumOfABinaryTree {
    public static void main(String[] args) {
        check(new Solution().maxLevelSum(TreeNode.of(1, 7, 0, 7, -8, null, null)), 2);
        check(new Solution().maxLevelSum(TreeNode.of(-1, -2, -3)), 1);
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
