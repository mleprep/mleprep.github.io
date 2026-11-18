// 1146. Snapshot Array (Medium)
// https://leetcode.com/problems/snapshot-array/
// Run (Java 11+): java P1146SnapshotArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1146SnapshotArray {
    public static void main(String[] args) {
        SnapshotArray sa = new SnapshotArray(3);
        sa.set(0, 5);
        check(sa.snap(), 0);
        sa.set(0, 6);
        check(sa.get(0, 0), 5);
        check(sa.get(1, 0), 0);
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

class SnapshotArray {
    private final List<TreeMap<Integer, Integer>> hist = new ArrayList<>();
    private int snapId = 0;

    public SnapshotArray(int length) {
        for (int i = 0; i < length; i++) {
            TreeMap<Integer, Integer> t = new TreeMap<>();
            t.put(-1, 0);
            hist.add(t);
        }
    }

    public void set(int index, int val) { hist.get(index).put(snapId, val); }
    public int snap() { return snapId++; }
    public int get(int index, int snap_id) { return hist.get(index).floorEntry(snap_id).getValue(); }
}
