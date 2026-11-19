// 1039. Minimum Score Triangulation of Polygon (Medium)
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Run (Java 11+): java P1039MinimumScoreTriangulationOfPolygon.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1039MinimumScoreTriangulationOfPolygon {
    public static void main(String[] args) {
        check(new Solution().minScoreTriangulation(new int[]{1, 3, 1, 4, 1, 5}), 13);
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
    public int minScoreTriangulation(int[] v) {
        int n = v.length;
        int[][] dp = new int[n][n];
        for (int gap = 2; gap < n; gap++)
            for (int i = 0; i + gap < n; i++) {
                int j = i + gap;
                dp[i][j] = Integer.MAX_VALUE;
                for (int k = i + 1; k < j; k++) dp[i][j] = Math.min(dp[i][j], dp[i][k] + v[i] * v[k] * v[j] + dp[k][j]);
            }
        return dp[0][n - 1];
    }
}
