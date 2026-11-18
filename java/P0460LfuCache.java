// 460. LFU Cache (Hard)
// https://leetcode.com/problems/lfu-cache/
// Run (Java 11+): java P0460LfuCache.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0460LfuCache {
    public static void main(String[] args) {
        LFUCache c = new LFUCache(2);
        c.put(1, 1); c.put(2, 2);
        check(c.get(1), 1);
        c.put(3, 3);
        check(c.get(2), -1);
        check(c.get(3), 3);
        c.put(4, 4);
        check(c.get(1), -1);
        check(c.get(3), 3);
        check(c.get(4), 4);
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

class LFUCache {
    private final int cap;
    private int minF = 0;
    private final Map<Integer, int[]> kv = new HashMap<>();                    // key -> {value, freq}
    private final Map<Integer, LinkedHashSet<Integer>> buckets = new HashMap<>(); // freq -> keys (LRU order)

    public LFUCache(int capacity) { cap = capacity; }

    private void touch(int key) {
        int[] e = kv.get(key);
        LinkedHashSet<Integer> b = buckets.get(e[1]);
        b.remove(key);
        if (b.isEmpty()) { buckets.remove(e[1]); if (minF == e[1]) minF++; }
        e[1]++;
        buckets.computeIfAbsent(e[1], k -> new LinkedHashSet<>()).add(key);
    }

    public int get(int key) {
        if (!kv.containsKey(key)) return -1;
        touch(key);
        return kv.get(key)[0];
    }

    public void put(int key, int value) {
        if (cap == 0) return;
        if (kv.containsKey(key)) { kv.get(key)[0] = value; touch(key); return; }
        if (kv.size() == cap) {
            LinkedHashSet<Integer> b = buckets.get(minF);
            int old = b.iterator().next();
            b.remove(old);
            if (b.isEmpty()) buckets.remove(minF);
            kv.remove(old);
        }
        kv.put(key, new int[]{value, 1});
        buckets.computeIfAbsent(1, k -> new LinkedHashSet<>()).add(key);
        minF = 1;
    }
}
