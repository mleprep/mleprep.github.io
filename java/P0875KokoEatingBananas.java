// 875. Koko Eating Bananas (Medium)
// https://leetcode.com/problems/koko-eating-bananas/
// Run (Java 11+): java P0875KokoEatingBananas.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0875KokoEatingBananas {
    public static void main(String[] args) {
        check(new Solution().minEatingSpeed(new int[]{3, 6, 7, 11}, 8), 4);
        check(new Solution().minEatingSpeed(new int[]{30, 11, 23, 4, 20}, 6), 23);
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
    public int minEatingSpeed(int[] piles, int h) {
        int lo = 1, hi = 0;
        for (int p : piles) hi = Math.max(hi, p);
        while (lo < hi) {
            int k = lo + (hi - lo) / 2;
            long hours = 0;
            for (int p : piles) hours += (p + k - 1) / k;
            if (hours <= h) hi = k; else lo = k + 1;
        }
        return lo;
    }
}
