// 877. Stone Game (Medium)
// https://leetcode.com/problems/stone-game/
// Run (Java 11+): java P0877StoneGame.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0877StoneGame {
    public static void main(String[] args) {
        check(new Solution().stoneGame(new int[]{5, 3, 4, 5}), true);
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
    public boolean stoneGame(int[] piles) {
        int n = piles.length;
        int[] dp = piles.clone();
        for (int len = 2; len <= n; len++)
            for (int i = 0; i + len <= n; i++) dp[i] = Math.max(piles[i] - dp[i + 1], piles[i + len - 1] - dp[i]);
        return dp[0] > 0;
    }
}
