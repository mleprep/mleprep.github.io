// 435. Non-overlapping Intervals (Medium)
// https://leetcode.com/problems/non-overlapping-intervals/
// Run (Java 11+): java P0435NonOverlappingIntervals.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0435NonOverlappingIntervals {
    public static void main(String[] args) {
        check(new Solution().eraseOverlapIntervals(new int[][]{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1);
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
    public int eraseOverlapIntervals(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[1], b[1]));
        int keep = 0;
        long end = Long.MIN_VALUE;
        for (int[] iv : intervals) if (iv[0] >= end) { keep++; end = iv[1]; }
        return intervals.length - keep;
    }
}
