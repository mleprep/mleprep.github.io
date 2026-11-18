// 97. Interleaving String (Medium)
// https://leetcode.com/problems/interleaving-string/
// Run (Java 11+): java P0097InterleavingString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0097InterleavingString {
    public static void main(String[] args) {
        check(new Solution().isInterleave("aabcc", "dbbca", "aadbbcbcac"), true);
        check(new Solution().isInterleave("aabcc", "dbbca", "aadbbbaccc"), false);
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
    public boolean isInterleave(String s1, String s2, String s3) {
        int m = s1.length(), n = s2.length();
        if (m + n != s3.length()) return false;
        boolean[] dp = new boolean[n + 1];
        for (int i = 0; i <= m; i++)
            for (int j = 0; j <= n; j++) {
                if (i == 0 && j == 0) { dp[j] = true; continue; }
                boolean a = i > 0 && dp[j] && s1.charAt(i - 1) == s3.charAt(i + j - 1);
                boolean b = j > 0 && dp[j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1);
                dp[j] = a || b;
            }
        return dp[n];
    }
}
