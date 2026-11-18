// 516. Longest Palindromic Subsequence (Medium)
// https://leetcode.com/problems/longest-palindromic-subsequence/
// Run (Java 11+): java P0516LongestPalindromicSubsequence.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0516LongestPalindromicSubsequence {
    public static void main(String[] args) {
        check(new Solution().longestPalindromeSubseq("bbbab"), 4);
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
    public int longestPalindromeSubseq(String s) {
        int n = s.length();
        int[][] dp = new int[n][n];
        for (int l = n - 1; l >= 0; l--) {
            dp[l][l] = 1;
            for (int r = l + 1; r < n; r++)
                dp[l][r] = s.charAt(l) == s.charAt(r) ? dp[l + 1][r - 1] + 2 : Math.max(dp[l + 1][r], dp[l][r - 1]);
        }
        return dp[0][n - 1];
    }
}
