// 1109. Corporate Flight Bookings (Medium)
// https://leetcode.com/problems/corporate-flight-bookings/
// Run (Java 11+): java P1109CorporateFlightBookings.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1109CorporateFlightBookings {
    public static void main(String[] args) {
        check(new Solution().corpFlightBookings(new int[][]{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5), new int[]{10, 55, 45, 25, 25});
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
    public int[] corpFlightBookings(int[][] bookings, int n) {
        int[] d = new int[n + 1];
        for (int[] b : bookings) { d[b[0] - 1] += b[2]; d[b[1]] -= b[2]; }
        int[] res = new int[n];
        for (int i = 0, s = 0; i < n; i++) res[i] = s += d[i];
        return res;
    }
}
