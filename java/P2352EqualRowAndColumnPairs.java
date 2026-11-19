// 2352. Equal Row and Column Pairs (Medium)
// https://leetcode.com/problems/equal-row-and-column-pairs/
// Run (Java 11+): java P2352EqualRowAndColumnPairs.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2352EqualRowAndColumnPairs {
    public static void main(String[] args) {
        check(new Solution().equalPairs(new int[][]{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}), 3);
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
    public int equalPairs(int[][] grid) {
        int n = grid.length, res = 0;
        Map<String, Integer> rows = new HashMap<>();
        for (int[] r : grid) rows.merge(Arrays.toString(r), 1, Integer::sum);
        for (int c = 0; c < n; c++) {
            int[] col = new int[n];
            for (int r = 0; r < n; r++) col[r] = grid[r][c];
            res += rows.getOrDefault(Arrays.toString(col), 0);
        }
        return res;
    }
}
