// 63. Unique Paths II (Medium)
// https://leetcode.com/problems/unique-paths-ii/
// Run (Java 11+): java P0063UniquePathsIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0063UniquePathsIi {
    public static void main(String[] args) {
        check(new Solution().uniquePathsWithObstacles(new int[][]{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}), 2);
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
    public int uniquePathsWithObstacles(int[][] grid) {
        int n = grid[0].length;
        int[] row = new int[n];
        row[0] = 1;
        for (int[] r : grid)
            for (int c = 0; c < n; c++) {
                if (r[c] == 1) row[c] = 0;
                else if (c > 0) row[c] += row[c - 1];
            }
        return row[n - 1];
    }
}
