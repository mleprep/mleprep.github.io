// 138. Copy List with Random Pointer (Medium)
// https://leetcode.com/problems/copy-list-with-random-pointer/
// Run (Java 11+): java P0138CopyListWithRandomPointer.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0138CopyListWithRandomPointer {
    public static void main(String[] args) {
        Node a = new Node(7), b = new Node(13), c = new Node(11);
        a.next = b; b.next = c; b.random = a; c.random = c;
        Node x = new Solution().copyRandomList(a);
        check(x != a && x.val == 7 && x.next.val == 13 && x.next.random == x && x.next.next.random == x.next.next, true);
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
    int val; Node next, random;
    Node(int val) { this.val = val; }
}

class Solution {
    public Node copyRandomList(Node head) {
        Map<Node, Node> clone = new HashMap<>();
        for (Node cur = head; cur != null; cur = cur.next) clone.put(cur, new Node(cur.val));
        for (Node cur = head; cur != null; cur = cur.next) {
            clone.get(cur).next = clone.get(cur.next);     // get(null) -> null
            clone.get(cur).random = clone.get(cur.random);
        }
        return clone.get(head);
    }
}
