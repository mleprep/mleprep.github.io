// 289. Game of Life (Medium)
// https://leetcode.com/problems/game-of-life/
// Run (Java 11+): java P0289GameOfLife.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0289GameOfLife {
    public static void main(String[] args) {
        int[][] b = {{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}};
        new Solution().gameOfLife(b);
        check(b, new int[][]{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}});
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
    public void gameOfLife(int[][] board) {
        int R = board.length, C = board[0].length;
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) {
                int live = -(board[r][c] & 1);
                for (int i = Math.max(0, r - 1); i <= Math.min(R - 1, r + 1); i++)
                    for (int j = Math.max(0, c - 1); j <= Math.min(C - 1, c + 1); j++) live += board[i][j] & 1;
                if (live == 3 || (live == 2 && (board[r][c] & 1) == 1)) board[r][c] |= 2;
            }
        for (int[] row : board)
            for (int c = 0; c < C; c++) row[c] >>= 1;
    }
}
