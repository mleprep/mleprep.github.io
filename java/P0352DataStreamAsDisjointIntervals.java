// 352. Data Stream as Disjoint Intervals (Hard)
// https://leetcode.com/problems/data-stream-as-disjoint-intervals/
// Run (Java 11+): java P0352DataStreamAsDisjointIntervals.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0352DataStreamAsDisjointIntervals {
    public static void main(String[] args) {
        SummaryRanges sr = new SummaryRanges();
        for (int v : new int[]{1, 3, 7, 2, 6}) sr.addNum(v);
        check(sr.getIntervals(), new int[][]{{1, 3}, {6, 7}});
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

class SummaryRanges {
    private final TreeMap<Integer, Integer> iv = new TreeMap<>();   // start -> end

    public void addNum(int value) {
        Integer lo = iv.floorKey(value), hi = iv.ceilingKey(value + 1);
        if (lo != null && iv.get(lo) >= value) return;
        boolean left = lo != null && iv.get(lo) == value - 1, right = hi != null && hi == value + 1;
        if (left && right) { iv.put(lo, iv.get(hi)); iv.remove(hi); }
        else if (left) iv.put(lo, value);
        else if (right) { iv.put(value, iv.get(hi)); iv.remove(hi); }
        else iv.put(value, value);
    }

    public int[][] getIntervals() {
        int[][] res = new int[iv.size()][];
        int i = 0;
        for (Map.Entry<Integer, Integer> e : iv.entrySet()) res[i++] = new int[]{e.getKey(), e.getValue()};
        return res;
    }
}
