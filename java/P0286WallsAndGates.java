// 286. Walls and Gates (Medium)
// https://leetcode.com/problems/walls-and-gates/
// Run (Java 11+): java P0286WallsAndGates.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0286WallsAndGates {
    public static void main(String[] args) {
        int I = Integer.MAX_VALUE;
        int[][] g = {{I, -1, 0, I}, {I, I, I, -1}, {I, -1, I, -1}, {0, -1, I, I}};
        new Solution().wallsAndGates(g);
        check(g, new int[][]{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}});
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
    public void wallsAndGates(int[][] rooms) {
        int R = rooms.length, C = rooms[0].length, INF = Integer.MAX_VALUE;
        Deque<int[]> q = new ArrayDeque<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) if (rooms[r][c] == 0) q.add(new int[]{r, c});
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int[] d : dirs) {
                int nr = cur[0] + d[0], nc = cur[1] + d[1];
                if (nr >= 0 && nc >= 0 && nr < R && nc < C && rooms[nr][nc] == INF) {
                    rooms[nr][nc] = rooms[cur[0]][cur[1]] + 1;
                    q.add(new int[]{nr, nc});
                }
            }
        }
    }
}
