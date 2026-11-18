// 452. Minimum Number of Arrows to Burst Balloons (Medium)
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// Run (Java 11+): java P0452MinimumNumberOfArrowsToBurstBalloons.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0452MinimumNumberOfArrowsToBurstBalloons {
    public static void main(String[] args) {
        check(new Solution().findMinArrowShots(new int[][]{{10, 16}, {2, 8}, {1, 6}, {7, 12}}), 2);
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
    public int findMinArrowShots(int[][] points) {
        Arrays.sort(points, (a, b) -> Integer.compare(a[1], b[1]));  // compare() avoids overflow
        int arrows = 0;
        long pos = Long.MIN_VALUE;
        for (int[] p : points) if (p[0] > pos) { arrows++; pos = p[1]; }
        return arrows;
    }
}
