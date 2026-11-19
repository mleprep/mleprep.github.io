// 1926. Nearest Exit from Entrance in Maze (Medium)
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Run (Java 11+): java P1926NearestExitFromEntranceInMaze.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1926NearestExitFromEntranceInMaze {
    public static void main(String[] args) {
        check(new Solution().nearestExit(new char[][]{"++.+".toCharArray(), "...+".toCharArray(), "+++.".toCharArray()}, new int[]{1, 2}), 1);
        check(new Solution().nearestExit(new char[][]{".+".toCharArray()}, new int[]{0, 0}), -1);
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
    public int nearestExit(char[][] maze, int[] entrance) {
        int R = maze.length, C = maze[0].length;
        boolean[][] seen = new boolean[R][C];
        seen[entrance[0]][entrance[1]] = true;
        Deque<int[]> q = new ArrayDeque<>(List.of(entrance));
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int i = q.size(); i > 0; i--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int r = cur[0] + d[0], c = cur[1] + d[1];
                    if (r < 0 || c < 0 || r >= R || c >= C || seen[r][c] || maze[r][c] != '.') continue;
                    if (r == 0 || c == 0 || r == R - 1 || c == C - 1) return steps;
                    seen[r][c] = true;
                    q.add(new int[]{r, c});
                }
            }
        }
        return -1;
    }
}
