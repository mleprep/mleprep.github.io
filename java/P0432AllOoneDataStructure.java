// 432. All O`one Data Structure (Hard)
// https://leetcode.com/problems/all-oone-data-structure/
// Run (Java 11+): java P0432AllOoneDataStructure.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0432AllOoneDataStructure {
    public static void main(String[] args) {
        AllOne a = new AllOne();
        a.inc("hello"); a.inc("hello"); a.inc("leet");
        check(a.getMaxKey(), "hello");
        check(a.getMinKey(), "leet");
        a.dec("hello"); a.dec("hello");
        check(a.getMaxKey(), "leet");
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

class AllOne {
    private static class Bucket {
        int cnt; Set<String> keys = new HashSet<>(); Bucket prev, next;
        Bucket(int c) { cnt = c; }
    }

    private final Bucket head = new Bucket(0), tail = new Bucket(0);
    private final Map<String, Bucket> where = new HashMap<>();

    public AllOne() { head.next = tail; tail.prev = head; }

    private Bucket after(Bucket node, int cnt) {
        Bucket b = new Bucket(cnt);
        b.prev = node; b.next = node.next;
        node.next.prev = b; node.next = b;
        return b;
    }

    private void dropIfEmpty(Bucket b) {
        if (b.keys.isEmpty()) { b.prev.next = b.next; b.next.prev = b.prev; }
    }

    public void inc(String key) {
        Bucket cur = where.getOrDefault(key, head), nxt = cur.next;
        if (nxt == tail || nxt.cnt != cur.cnt + 1) nxt = after(cur, cur.cnt + 1);
        nxt.keys.add(key);
        where.put(key, nxt);
        if (cur != head) { cur.keys.remove(key); dropIfEmpty(cur); }
    }

    public void dec(String key) {
        Bucket cur = where.get(key);
        if (cur.cnt == 1) where.remove(key);
        else {
            Bucket prv = cur.prev;
            if (prv == head || prv.cnt != cur.cnt - 1) prv = after(cur.prev, cur.cnt - 1);
            prv.keys.add(key);
            where.put(key, prv);
        }
        cur.keys.remove(key);
        dropIfEmpty(cur);
    }

    public String getMaxKey() { return tail.prev == head ? "" : tail.prev.keys.iterator().next(); }
    public String getMinKey() { return head.next == tail ? "" : head.next.keys.iterator().next(); }
}
