// 417. Pacific Atlantic Water Flow (Medium)
// https://leetcode.com/problems/pacific-atlantic-water-flow/
// Run (Java 11+): java P0417PacificAtlanticWaterFlow.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0417PacificAtlanticWaterFlow {
    public static void main(String[] args) {
        int[][] H = {{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}};
        check(new Solution().pacificAtlantic(H), "[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]");
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
    public List<List<Integer>> pacificAtlantic(int[][] heights) {
        int R = heights.length, C = heights[0].length;
        boolean[][] pac = new boolean[R][C], atl = new boolean[R][C];
        for (int r = 0; r < R; r++) { dfs(heights, r, 0, pac); dfs(heights, r, C - 1, atl); }
        for (int c = 0; c < C; c++) { dfs(heights, 0, c, pac); dfs(heights, R - 1, c, atl); }
        List<List<Integer>> res = new ArrayList<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++)
                if (pac[r][c] && atl[r][c]) res.add(List.of(r, c));
        return res;
    }

    private void dfs(int[][] h, int r, int c, boolean[][] seen) {
        if (seen[r][c]) return;
        seen[r][c] = true;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int[] d : dirs) {
            int nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nc >= 0 && nr < h.length && nc < h[0].length && h[nr][nc] >= h[r][c]) dfs(h, nr, nc, seen);
        }
    }
}
