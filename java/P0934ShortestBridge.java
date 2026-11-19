// 934. Shortest Bridge (Medium)
// https://leetcode.com/problems/shortest-bridge/
// Run (Java 11+): java P0934ShortestBridge.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0934ShortestBridge {
    public static void main(String[] args) {
        check(new Solution().shortestBridge(new int[][]{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}), 2);
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
    public int shortestBridge(int[][] grid) {
        int n = grid.length;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        Deque<int[]> q = new ArrayDeque<>(), st = new ArrayDeque<>();
        outer:
        for (int r = 0; r < n; r++)
            for (int c = 0; c < n; c++)
                if (grid[r][c] == 1) { grid[r][c] = 2; st.push(new int[]{r, c}); break outer; }
        while (!st.isEmpty()) {
            int[] cur = st.pop();
            q.add(cur);
            for (int[] d : dirs) {
                int nr = cur[0] + d[0], nc = cur[1] + d[1];
                if (nr >= 0 && nc >= 0 && nr < n && nc < n && grid[nr][nc] == 1) { grid[nr][nc] = 2; st.push(new int[]{nr, nc}); }
            }
        }
        for (int steps = 0; !q.isEmpty(); steps++)
            for (int k = q.size(); k > 0; k--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int nr = cur[0] + d[0], nc = cur[1] + d[1];
                    if (nr < 0 || nc < 0 || nr >= n || nc >= n) continue;
                    if (grid[nr][nc] == 1) return steps;
                    if (grid[nr][nc] == 0) { grid[nr][nc] = 2; q.add(new int[]{nr, nc}); }
                }
            }
        return -1;
    }
}
