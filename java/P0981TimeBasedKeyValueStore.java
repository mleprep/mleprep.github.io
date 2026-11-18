// 981. Time Based Key-Value Store (Medium)
// https://leetcode.com/problems/time-based-key-value-store/
// Run (Java 11+): java P0981TimeBasedKeyValueStore.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0981TimeBasedKeyValueStore {
    public static void main(String[] args) {
        TimeMap tm = new TimeMap();
        tm.set("foo", "bar", 1);
        check(tm.get("foo", 3), "bar");
        tm.set("foo", "bar2", 4);
        check(tm.get("foo", 5), "bar2");
        check(tm.get("foo", 0), "");
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

class TimeMap {
    private final Map<String, TreeMap<Integer, String>> data = new HashMap<>();

    public void set(String key, String value, int timestamp) {
        data.computeIfAbsent(key, k -> new TreeMap<>()).put(timestamp, value);
    }

    public String get(String key, int timestamp) {
        TreeMap<Integer, String> t = data.get(key);
        if (t == null) return "";
        Map.Entry<Integer, String> e = t.floorEntry(timestamp); // largest ts <= timestamp
        return e == null ? "" : e.getValue();
    }
}
