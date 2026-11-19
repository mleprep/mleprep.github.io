// 705. Design HashSet (Easy)
// https://leetcode.com/problems/design-hashset/
// Run (Java 11+): java P0705DesignHashset.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0705DesignHashset {
    public static void main(String[] args) {
        MyHashSet hs = new MyHashSet();
        hs.add(1); hs.add(1009); hs.remove(1);
        check(hs.contains(1) + " " + hs.contains(1009), "false true");
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

class MyHashSet {
    private static final int SIZE = 1009;
    private final List<List<Integer>> buckets = new ArrayList<>();

    public MyHashSet() { for (int i = 0; i < SIZE; i++) buckets.add(new LinkedList<>()); }

    public void add(int key) { if (!contains(key)) buckets.get(key % SIZE).add(key); }

    public void remove(int key) { buckets.get(key % SIZE).remove(Integer.valueOf(key)); }

    public boolean contains(int key) { return buckets.get(key % SIZE).contains(key); }
}
