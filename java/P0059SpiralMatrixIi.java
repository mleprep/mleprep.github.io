// 59. Spiral Matrix II (Medium)
// https://leetcode.com/problems/spiral-matrix-ii/
// Run (Java 11+): java P0059SpiralMatrixIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0059SpiralMatrixIi {
    public static void main(String[] args) {
        check(new Solution().generateMatrix(3), new int[][]{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}});
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
    public int[][] generateMatrix(int n) {
        int[][] m = new int[n][n];
        int top = 0, bottom = n - 1, left = 0, right = n - 1, k = 1;
        while (top <= bottom) {
            for (int c = left; c <= right; c++) m[top][c] = k++;
            for (int r = top + 1; r <= bottom; r++) m[r][right] = k++;
            if (top < bottom) for (int c = right - 1; c >= left; c--) m[bottom][c] = k++;
            if (left < right) for (int r = bottom - 1; r > top; r--) m[r][left] = k++;
            top++; bottom--; left++; right--;
        }
        return m;
    }
}
