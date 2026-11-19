// 1406. Stone Game III (Hard)
// https://leetcode.com/problems/stone-game-iii/
// Run (Java 11+): java P1406StoneGameIii.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1406StoneGameIii {
    public static void main(String[] args) {
        check(new Solution().stoneGameIII(new int[]{1, 2, 3, 7}), "Bob");
        check(new Solution().stoneGameIII(new int[]{1, 2, 3, -9}), "Alice");
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
    public String stoneGameIII(int[] stoneValue) {
        int n = stoneValue.length;
        int[] dp = new int[n + 3];
        for (int i = n - 1; i >= 0; i--) {
            int take = 0, best = Integer.MIN_VALUE;
            for (int k = 0; k < 3 && i + k < n; k++) {
                take += stoneValue[i + k];
                best = Math.max(best, take - dp[i + k + 1]);
            }
            dp[i] = best;
        }
        return dp[0] > 0 ? "Alice" : dp[0] < 0 ? "Bob" : "Tie";
    }
}
