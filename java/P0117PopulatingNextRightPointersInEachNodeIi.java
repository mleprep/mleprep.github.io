// 117. Populating Next Right Pointers in Each Node II (Medium)
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// Run (Java 11+): java P0117PopulatingNextRightPointersInEachNodeIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0117PopulatingNextRightPointersInEachNodeIi {
    public static void main(String[] args) {
        Node[] n = new Node[8];
        for (int i = 1; i < 8; i++) n[i] = new Node(i);
        n[1].left = n[2]; n[1].right = n[3]; n[2].left = n[4]; n[2].right = n[5]; n[3].right = n[7];
        new Solution().connect(n[1]);
        check(n[4].next == n[5] && n[5].next == n[7] && n[7].next == null && n[2].next == n[3], true);
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
