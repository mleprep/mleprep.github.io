// 221. Maximal Square (Medium)
// https://leetcode.com/problems/maximal-square/
// Run (Java 11+): java P0221MaximalSquare.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0221MaximalSquare {
    public static void main(String[] args) {
        char[][] m = Stream.of("10100", "10111", "11111", "10010").map(String::toCharArray).toArray(char[][]::new);
        check(new Solution().maximalSquare(m), 4);
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
    public int maximalSquare(char[][] matrix) {
        int n = matrix[0].length, best = 0;
        int[] dp = new int[n + 1];
        for (char[] row : matrix) {
            int diag = 0;
            for (int c = 1; c <= n; c++) {
                int up = dp[c];
                dp[c] = row[c - 1] == '1' ? 1 + Math.min(diag, Math.min(dp[c], dp[c - 1])) : 0;
                diag = up;
                best = Math.max(best, dp[c]);
            }
        }
        return best * best;
    }
}
