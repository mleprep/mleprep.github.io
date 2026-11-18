// 115. Distinct Subsequences (Hard)
// https://leetcode.com/problems/distinct-subsequences/
// Run (Java 11+): java P0115DistinctSubsequences.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0115DistinctSubsequences {
    public static void main(String[] args) {
        check(new Solution().numDistinct("babgbag", "bag"), 5);
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
    public int numDistinct(String s, String t) {
        long[] dp = new long[t.length() + 1];
        dp[0] = 1;
        for (char ch : s.toCharArray())
            for (int j = t.length(); j > 0; j--)
                if (t.charAt(j - 1) == ch) dp[j] += dp[j - 1];
        return (int) dp[t.length()];
    }
}
