// 2807. Insert Greatest Common Divisors in Linked List (Medium)
// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Run (Java 11+): java P2807InsertGreatestCommonDivisorsInLinkedList.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2807InsertGreatestCommonDivisorsInLinkedList {
    public static void main(String[] args) {
        check(new Solution().insertGreatestCommonDivisors(ListNode.of(18, 6, 10, 3)), "[18, 6, 6, 2, 10, 1, 3]");
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
    public ListNode insertGreatestCommonDivisors(ListNode head) {
        for (ListNode cur = head; cur != null && cur.next != null; cur = cur.next.next)
            cur.next = new ListNode(gcd(cur.val, cur.next.val), cur.next);
        return head;
    }

    private int gcd(int a, int b) { return b == 0 ? a : gcd(b, a % b); }
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
