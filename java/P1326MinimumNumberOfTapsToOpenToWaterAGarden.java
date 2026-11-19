// 1326. Minimum Number of Taps to Open to Water a Garden (Hard)
// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// Run (Java 11+): java P1326MinimumNumberOfTapsToOpenToWaterAGarden.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1326MinimumNumberOfTapsToOpenToWaterAGarden {
    public static void main(String[] args) {
        check(new Solution().minTaps(7, new int[]{1, 2, 1, 0, 2, 1, 0, 1}), 3);
        check(new Solution().minTaps(3, new int[]{0, 0, 0, 0}), -1);
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
    public int minTaps(int n, int[] ranges) {
        int[] reach = new int[n + 1];
        for (int i = 0; i <= n; i++) {
            int lo = Math.max(0, i - ranges[i]), hi = Math.min(n, i + ranges[i]);
            reach[lo] = Math.max(reach[lo], hi);
        }
        int taps = 0, end = 0, far = 0;
        for (int i = 0; i < n; i++) {
            far = Math.max(far, reach[i]);
            if (i == end) {
                if (far <= i) return -1;
                taps++;
                end = far;
            }
        }
        return taps;
    }
}
