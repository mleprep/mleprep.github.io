// 54. Spiral Matrix (Medium)
// https://leetcode.com/problems/spiral-matrix/
// Run (Java 11+): java P0054SpiralMatrix.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0054SpiralMatrix {
    public static void main(String[] args) {
        check(new Solution().spiralOrder(new int[][]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), "[1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]");
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
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        int top = 0, bottom = matrix.length - 1, left = 0, right = matrix[0].length - 1;
        while (top <= bottom && left <= right) {
            for (int c = left; c <= right; c++) res.add(matrix[top][c]);
            for (int r = top + 1; r <= bottom; r++) res.add(matrix[r][right]);
            if (top < bottom && left < right) {
                for (int c = right - 1; c >= left; c--) res.add(matrix[bottom][c]);
                for (int r = bottom - 1; r > top; r--) res.add(matrix[r][left]);
            }
            top++; bottom--; left++; right--;
        }
        return res;
    }
}
