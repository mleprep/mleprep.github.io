// 62. Unique Paths (Medium)
// https://leetcode.com/problems/unique-paths/
// Run (Java 11+): java P0062UniquePaths.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0062UniquePaths {
    public static void main(String[] args) {
        check(new Solution().uniquePaths(3, 7), 28);
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
    public int uniquePaths(int m, int n) {
        int[] row = new int[n];
        Arrays.fill(row, 1);
        for (int r = 1; r < m; r++)
            for (int c = 1; c < n; c++) row[c] += row[c - 1];
        return row[n - 1];
    }
}
