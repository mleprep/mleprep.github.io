// 986. Interval List Intersections (Medium)
// https://leetcode.com/problems/interval-list-intersections/
// Run (Java 11+): java P0986IntervalListIntersections.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0986IntervalListIntersections {
    public static void main(String[] args) {
        int[][] A = {{0, 2}, {5, 10}, {13, 23}, {24, 25}}, B = {{1, 5}, {8, 12}, {15, 24}, {25, 26}};
        check(new Solution().intervalIntersection(A, B), new int[][]{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}});
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
    public int[][] intervalIntersection(int[][] a, int[][] b) {
        List<int[]> res = new ArrayList<>();
        int i = 0, j = 0;
        while (i < a.length && j < b.length) {
            int lo = Math.max(a[i][0], b[j][0]), hi = Math.min(a[i][1], b[j][1]);
            if (lo <= hi) res.add(new int[]{lo, hi});
            if (a[i][1] < b[j][1]) i++; else j++;
        }
        return res.toArray(new int[0][]);
    }
}
