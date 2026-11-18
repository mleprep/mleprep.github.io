// 10. Regular Expression Matching (Hard)
// https://leetcode.com/problems/regular-expression-matching/
// Run (Java 11+): java P0010RegularExpressionMatching.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0010RegularExpressionMatching {
    public static void main(String[] args) {
        check(new Solution().isMatch("aab", "c*a*b"), true);
        check(new Solution().isMatch("mississippi", "mis*is*p*."), false);
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
    public boolean isMatch(String s, String p) {
        int m = s.length(), n = p.length();
        boolean[][] dp = new boolean[m + 1][n + 1];
        dp[m][n] = true;
        for (int i = m; i >= 0; i--)
            for (int j = n - 1; j >= 0; j--) {
                boolean first = i < m && (p.charAt(j) == s.charAt(i) || p.charAt(j) == '.');
                if (j + 1 < n && p.charAt(j + 1) == '*') dp[i][j] = dp[i][j + 2] || (first && dp[i + 1][j]);
                else dp[i][j] = first && dp[i + 1][j + 1];
            }
        return dp[0][0];
    }
}
