// 1482. Minimum Number of Days to Make m Bouquets (Medium)
// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
// Run (Java 11+): java P1482MinimumNumberOfDaysToMakeMBouquets.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1482MinimumNumberOfDaysToMakeMBouquets {
    public static void main(String[] args) {
        check(new Solution().minDays(new int[]{7, 7, 7, 7, 12, 7, 7}, 2, 3), 12);
        check(new Solution().minDays(new int[]{1, 10, 3, 10, 2}, 3, 2), -1);
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
    public int minDays(int[] bloomDay, int m, int k) {
        if ((long) m * k > bloomDay.length) return -1;
        int lo = Integer.MAX_VALUE, hi = 0;
        for (int b : bloomDay) { lo = Math.min(lo, b); hi = Math.max(hi, b); }
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            int made = 0, run = 0;
            for (int b : bloomDay) {
                run = b <= mid ? run + 1 : 0;
                if (run == k) { made++; run = 0; }
            }
            if (made >= m) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
