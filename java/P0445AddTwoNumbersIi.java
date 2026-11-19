// 445. Add Two Numbers II (Medium)
// https://leetcode.com/problems/add-two-numbers-ii/
// Run (Java 11+): java P0445AddTwoNumbersIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0445AddTwoNumbersIi {
    public static void main(String[] args) {
        check(new Solution().addTwoNumbers(ListNode.of(7, 2, 4, 3), ListNode.of(5, 6, 4)), ListNode.of(7, 8, 0, 7));
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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        Deque<Integer> a = new ArrayDeque<>(), b = new ArrayDeque<>();
        for (; l1 != null; l1 = l1.next) a.push(l1.val);
        for (; l2 != null; l2 = l2.next) b.push(l2.val);
        ListNode head = null;
        int carry = 0;
        while (!a.isEmpty() || !b.isEmpty() || carry != 0) {
            int s = carry + (a.isEmpty() ? 0 : a.pop()) + (b.isEmpty() ? 0 : b.pop());
            carry = s / 10;
            head = new ListNode(s % 10, head);
        }
        return head;
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
