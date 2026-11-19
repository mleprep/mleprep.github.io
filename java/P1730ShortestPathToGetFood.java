// 1730. Shortest Path to Get Food (Medium)
// https://leetcode.com/problems/shortest-path-to-get-food/
// Run (Java 11+): java P1730ShortestPathToGetFood.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1730ShortestPathToGetFood {
    public static void main(String[] args) {
        check(new Solution().getFood(new char[][]{"XXXXXX".toCharArray(), "X*OOOX".toCharArray(), "XOO#OX".toCharArray(), "XXXXXX".toCharArray()}), 3);
        check(new Solution().getFood(new char[][]{"*#".toCharArray()}), 1);
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
    public int getFood(char[][] grid) {
        int R = grid.length, C = grid[0].length;
        Deque<int[]> q = new ArrayDeque<>();
        boolean[][] seen = new boolean[R][C];
        for (int r = 0; r < R; r++) for (int c = 0; c < C; c++) if (grid[r][c] == '*') { q.add(new int[]{r, c}); seen[r][c] = true; }
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int i = q.size(); i > 0; i--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int r = cur[0] + d[0], c = cur[1] + d[1];
                    if (r < 0 || c < 0 || r >= R || c >= C || seen[r][c] || grid[r][c] == 'X') continue;
                    if (grid[r][c] == '#') return steps;
                    seen[r][c] = true;
                    q.add(new int[]{r, c});
                }
            }
        }
        return -1;
    }
}
