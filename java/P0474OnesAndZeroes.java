// 474. Ones and Zeroes (Medium)
// https://leetcode.com/problems/ones-and-zeroes/
// Run (Java 11+): java P0474OnesAndZeroes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0474OnesAndZeroes {
    public static void main(String[] args) {
        check(new Solution().findMaxForm(new String[]{"10", "0001", "111001", "1", "0"}, 5, 3), 4);
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
    public int findMaxForm(String[] strs, int m, int n) {
        int[][] dp = new int[m + 1][n + 1];
        for (String s : strs) {
            int z = 0, o = 0;
            for (char c : s.toCharArray()) if (c == '0') z++; else o++;
            for (int i = m; i >= z; i--)
                for (int j = n; j >= o; j--) dp[i][j] = Math.max(dp[i][j], dp[i - z][j - o] + 1);
        }
        return dp[m][n];
    }
}
