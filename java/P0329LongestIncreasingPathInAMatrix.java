// 329. Longest Increasing Path in a Matrix (Hard)
// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
// Run (Java 11+): java P0329LongestIncreasingPathInAMatrix.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0329LongestIncreasingPathInAMatrix {
    public static void main(String[] args) {
        check(new Solution().longestIncreasingPath(new int[][]{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}), 4);
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
    private int[][] memo;

    public int longestIncreasingPath(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length, best = 0;
        memo = new int[R][C];
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) best = Math.max(best, dfs(matrix, r, c));
        return best;
    }

    private int dfs(int[][] m, int r, int c) {
        if (memo[r][c] > 0) return memo[r][c];
        int res = 1;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int[] d : dirs) {
            int nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nc >= 0 && nr < m.length && nc < m[0].length && m[nr][nc] > m[r][c])
                res = Math.max(res, 1 + dfs(m, nr, nc));
        }
        return memo[r][c] = res;
    }
}
