// 1140. Stone Game II (Medium)
// https://leetcode.com/problems/stone-game-ii/
// Run (Java 11+): java P1140StoneGameIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1140StoneGameIi {
    public static void main(String[] args) {
        check(new Solution().stoneGameII(new int[]{2, 7, 9, 4, 4}), 10);
        check(new Solution().stoneGameII(new int[]{1, 2, 3, 4, 5, 100}), 104);
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
    public int stoneGameII(int[] piles) {
        int n = piles.length;
        int[] suffix = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) suffix[i] = suffix[i + 1] + piles[i];
        int[][] dp = new int[n + 1][n + 1];                   // dp[i][m]; M never exceeds n
        for (int i = n - 1; i >= 0; i--) {
            for (int m = 1; m <= n; m++) {
                if (i + 2 * m >= n) { dp[i][m] = suffix[i]; continue; }
                for (int x = 1; x <= 2 * m; x++) dp[i][m] = Math.max(dp[i][m], suffix[i] - dp[i + x][Math.max(m, x)]);
            }
        }
        return dp[0][1];
    }
}
