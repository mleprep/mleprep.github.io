// 1049. Last Stone Weight II (Medium)
// https://leetcode.com/problems/last-stone-weight-ii/
// Run (Java 11+): java P1049LastStoneWeightIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1049LastStoneWeightIi {
    public static void main(String[] args) {
        check(new Solution().lastStoneWeightII(new int[]{2, 7, 4, 1, 8, 1}), 1);
        check(new Solution().lastStoneWeightII(new int[]{31, 26, 33, 21, 40}), 5);
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
    public int lastStoneWeightII(int[] stones) {
        int total = 0;
        for (int s : stones) total += s;
        boolean[] dp = new boolean[total / 2 + 1];
        dp[0] = true;
        for (int s : stones) for (int t = total / 2; t >= s; t--) dp[t] |= dp[t - s];
        for (int t = total / 2; ; t--) if (dp[t]) return total - 2 * t;
    }
}
