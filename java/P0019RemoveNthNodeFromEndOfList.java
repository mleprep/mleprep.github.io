// 19. Remove Nth Node From End of List (Medium)
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Run (Java 11+): java P0019RemoveNthNodeFromEndOfList.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0019RemoveNthNodeFromEndOfList {
    public static void main(String[] args) {
        check(new Solution().removeNthFromEnd(ListNode.of(1, 2, 3, 4, 5), 2), ListNode.of(1, 2, 3, 5));
        check(new Solution().removeNthFromEnd(ListNode.of(1), 1), null);
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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(0, head), slow = dummy, fast = dummy;
        for (int i = 0; i < n; i++) fast = fast.next;
        while (fast.next != null) { slow = slow.next; fast = fast.next; }
        slow.next = slow.next.next;
        return dummy.next;
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
