// 867. Transpose Matrix (Easy)
// https://leetcode.com/problems/transpose-matrix/
// Run (Java 11+): java P0867TransposeMatrix.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0867TransposeMatrix {
    public static void main(String[] args) {
        check(new Solution().transpose(new int[][]{{1, 2, 3}, {4, 5, 6}}), new int[][]{{1, 4}, {2, 5}, {3, 6}});
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
    public int[][] transpose(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length;
        int[][] t = new int[C][R];
        for (int i = 0; i < R; i++) for (int j = 0; j < C; j++) t[j][i] = matrix[i][j];
        return t;
    }
}
