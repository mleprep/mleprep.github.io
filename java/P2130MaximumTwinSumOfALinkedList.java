// 2130. Maximum Twin Sum of a Linked List (Medium)
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
// Run (Java 11+): java P2130MaximumTwinSumOfALinkedList.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2130MaximumTwinSumOfALinkedList {
    public static void main(String[] args) {
        check(new Solution().pairSum(ListNode.of(5, 4, 2, 1)), 6);
        check(new Solution().pairSum(ListNode.of(4, 2, 2, 3)), 7);
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
    public int pairSum(ListNode head) {
        ListNode slow = head, fast = head, prev = null;
        while (fast != null && fast.next != null) { slow = slow.next; fast = fast.next.next; }
        while (slow != null) { ListNode nxt = slow.next; slow.next = prev; prev = slow; slow = nxt; }
        int best = 0;
        for (; prev != null; prev = prev.next, head = head.next) best = Math.max(best, head.val + prev.val);
        return best;
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
