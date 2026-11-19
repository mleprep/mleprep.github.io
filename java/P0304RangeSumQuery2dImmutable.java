// 304. Range Sum Query 2D - Immutable (Medium)
// https://leetcode.com/problems/range-sum-query-2d-immutable/
// Run (Java 11+): java P0304RangeSumQuery2dImmutable.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0304RangeSumQuery2dImmutable {
    public static void main(String[] args) {
        NumMatrix nm = new NumMatrix(new int[][]{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}});
        check(nm.sumRegion(2, 1, 4, 3), 8);
        check(nm.sumRegion(1, 2, 2, 4), 12);
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

class NumMatrix {
    private final int[][] P;

    public NumMatrix(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length;
        P = new int[R + 1][C + 1];
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) P[r + 1][c + 1] = matrix[r][c] + P[r][c + 1] + P[r + 1][c] - P[r][c];
    }

    public int sumRegion(int r1, int c1, int r2, int c2) {
        return P[r2 + 1][c2 + 1] - P[r1][c2 + 1] - P[r2 + 1][c1] + P[r1][c1];
    }
}
