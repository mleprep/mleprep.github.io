// 1423. Maximum Points You Can Obtain from Cards (Medium)
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
// Run (Java 11+): java P1423MaximumPointsYouCanObtainFromCards.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1423MaximumPointsYouCanObtainFromCards {
    public static void main(String[] args) {
        check(new Solution().maxScore(new int[]{1, 2, 3, 4, 5, 6, 1}, 3), 12);
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
    public int maxScore(int[] cardPoints, int k) {
        int n = cardPoints.length, w = n - k, s = 0, total = 0;
        for (int x : cardPoints) total += x;
        for (int i = 0; i < w; i++) s += cardPoints[i];
        int best = s;
        for (int i = w; i < n; i++) { s += cardPoints[i] - cardPoints[i - w]; best = Math.min(best, s); }
        return total - best;
    }
}
