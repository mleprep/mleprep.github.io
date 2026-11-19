// 1372. Longest ZigZag Path in a Binary Tree (Medium)
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
// Run (Java 11+): java P1372LongestZigzagPathInABinaryTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1372LongestZigzagPathInABinaryTree {
    public static void main(String[] args) {
        check(new Solution().longestZigZag(TreeNode.of(1, 1, 1, null, 1, null, null, 1, 1, null, 1)), 4);
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
