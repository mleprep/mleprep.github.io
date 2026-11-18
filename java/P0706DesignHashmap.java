// 706. Design HashMap (Easy)
// https://leetcode.com/problems/design-hashmap/
// Run (Java 11+): java P0706DesignHashmap.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0706DesignHashmap {
    public static void main(String[] args) {
        MyHashMap hm = new MyHashMap();
        hm.put(1, 1); hm.put(2, 2); hm.put(1010, 7);
        check(hm.get(1), 1);
        check(hm.get(1010), 7);
        hm.remove(2);
        check(hm.get(2), -1);
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

class MyHashMap {
    private static final int SIZE = 1009;
    private final List<int[]>[] b;

    @SuppressWarnings("unchecked")
    public MyHashMap() {
        b = new List[SIZE];
        for (int i = 0; i < SIZE; i++) b[i] = new LinkedList<>();
    }

    public void put(int key, int value) {
        for (int[] p : b[key % SIZE]) if (p[0] == key) { p[1] = value; return; }
        b[key % SIZE].add(new int[]{key, value});
    }

    public int get(int key) {
        for (int[] p : b[key % SIZE]) if (p[0] == key) return p[1];
        return -1;
    }

    public void remove(int key) { b[key % SIZE].removeIf(p -> p[0] == key); }
}
