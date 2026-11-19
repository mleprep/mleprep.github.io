// 1094. Car Pooling (Medium)
// https://leetcode.com/problems/car-pooling/
// Run (Java 11+): java P1094CarPooling.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1094CarPooling {
    public static void main(String[] args) {
        check(new Solution().carPooling(new int[][]{{2, 1, 5}, {3, 3, 7}}, 4), false);
        check(new Solution().carPooling(new int[][]{{2, 1, 5}, {3, 5, 7}}, 3), true);
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
    public boolean carPooling(int[][] trips, int capacity) {
        int[] d = new int[1001];
        for (int[] t : trips) { d[t[1]] += t[0]; d[t[2]] -= t[0]; }
        int load = 0;
        for (int x : d) if ((load += x) > capacity) return false;
        return true;
    }
}
