// 109. Convert Sorted List to Binary Search Tree (Medium)
// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
// Run (Java 11+): java P0109ConvertSortedListToBinarySearchTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0109ConvertSortedListToBinarySearchTree {
    public static void main(String[] args) {
        check(new Solution().sortedListToBST(ListNode.of(-10, -3, 0, 5, 9)), TreeNode.of(0, -10, 5, null, -3, null, 9));
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

// ---- LeetCode's ListNode + local test helpers ----
class ListNode {
    int val; ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    static ListNode of(int... vals) {
        ListNode dummy = new ListNode(), t = dummy;
        for (int v : vals) t = t.next = new ListNode(v);
        return dummy.next;
    }
    public String toString() {
        StringBuilder sb = new StringBuilder("[");
        int guard = 0;
        for (ListNode n = this; n != null && guard++ < 10000; n = n.next) sb.append(n.val).append(n.next != null ? ", " : "");
        return sb.append("]").toString();
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
