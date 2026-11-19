// 86. Partition List (Medium)
// https://leetcode.com/problems/partition-list/
// Run (Java 11+): java P0086PartitionList.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0086PartitionList {
    public static void main(String[] args) {
        check(new Solution().partition(ListNode.of(1, 4, 3, 2, 5, 2), 3), ListNode.of(1, 2, 2, 4, 3, 5));
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
    public ListNode partition(ListNode head, int x) {
        ListNode small = new ListNode(), big = new ListNode(), s = small, b = big;
        for (; head != null; head = head.next) {
            if (head.val < x) s = s.next = head; else b = b.next = head;
        }
        b.next = null;
        s.next = big.next;
        return small.next;
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
