// 25. Reverse Nodes in k-Group (Hard)
// https://leetcode.com/problems/reverse-nodes-in-k-group/
// Run (Java 11+): java P0025ReverseNodesInKGroup.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0025ReverseNodesInKGroup {
    public static void main(String[] args) {
        check(new Solution().reverseKGroup(ListNode.of(1, 2, 3, 4, 5), 2), ListNode.of(2, 1, 4, 3, 5));
        check(new Solution().reverseKGroup(ListNode.of(1, 2, 3, 4, 5), 3), ListNode.of(3, 2, 1, 4, 5));
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
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode(0, head), groupPrev = dummy;
        while (true) {
            ListNode kth = groupPrev;
            for (int i = 0; i < k; i++) {
                kth = kth.next;
                if (kth == null) return dummy.next;
            }
            ListNode groupNext = kth.next, prev = groupNext, cur = groupPrev.next;
            while (cur != groupNext) { ListNode nxt = cur.next; cur.next = prev; prev = cur; cur = nxt; }
            ListNode first = groupPrev.next;
            groupPrev.next = kth;
            groupPrev = first;
        }
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
