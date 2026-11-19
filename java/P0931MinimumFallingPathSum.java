// 931. Minimum Falling Path Sum (Medium)
// https://leetcode.com/problems/minimum-falling-path-sum/
// Run (Java 11+): java P0931MinimumFallingPathSum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0931MinimumFallingPathSum {
    public static void main(String[] args) {
        check(new Solution().minFallingPathSum(new int[][]{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}), 13);
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
    public int minFallingPathSum(int[][] matrix) {
        int n = matrix.length;
        int[] prev = matrix[0].clone();
        for (int r = 1; r < n; r++) {
            int[] cur = new int[n];
            for (int c = 0; c < n; c++) {
                int best = prev[c];
                if (c > 0) best = Math.min(best, prev[c - 1]);
                if (c < n - 1) best = Math.min(best, prev[c + 1]);
                cur[c] = matrix[r][c] + best;
            }
            prev = cur;
        }
        return Arrays.stream(prev).min().getAsInt();
    }
}
