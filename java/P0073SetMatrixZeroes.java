// 73. Set Matrix Zeroes (Medium)
// https://leetcode.com/problems/set-matrix-zeroes/
// Run (Java 11+): java P0073SetMatrixZeroes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0073SetMatrixZeroes {
    public static void main(String[] args) {
        int[][] m = {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}};
        new Solution().setZeroes(m);
        check(m, new int[][]{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}});
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
    public void setZeroes(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length;
        boolean row0 = false, col0 = false;
        for (int c = 0; c < C; c++) if (matrix[0][c] == 0) row0 = true;
        for (int r = 0; r < R; r++) if (matrix[r][0] == 0) col0 = true;
        for (int r = 1; r < R; r++)
            for (int c = 1; c < C; c++) if (matrix[r][c] == 0) matrix[r][0] = matrix[0][c] = 0;
        for (int r = 1; r < R; r++)
            for (int c = 1; c < C; c++) if (matrix[r][0] == 0 || matrix[0][c] == 0) matrix[r][c] = 0;
        if (row0) Arrays.fill(matrix[0], 0);
        if (col0) for (int r = 0; r < R; r++) matrix[r][0] = 0;
    }
}
