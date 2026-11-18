// 37. Sudoku Solver (Hard)
// https://leetcode.com/problems/sudoku-solver/
// Run (Java 11+): java P0037SudokuSolver.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0037SudokuSolver {
    public static void main(String[] args) {
        char[][] b = Stream.of("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
                "7...2...6", ".6....28.", "...419..5", "....8..79").map(String::toCharArray).toArray(char[][]::new);
        new Solution().solveSudoku(b);
        check(new String(b[0]) + new String(b[8]), "534678912345286179");
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
    private final boolean[][] row = new boolean[9][10], col = new boolean[9][10], box = new boolean[9][10];

    public void solveSudoku(char[][] board) {
        for (int r = 0; r < 9; r++)
            for (int c = 0; c < 9; c++)
                if (board[r][c] != '.') mark(r, c, board[r][c] - '0', true);
        bt(board, 0);
    }

    private void mark(int r, int c, int d, boolean v) { row[r][d] = col[c][d] = box[r / 3 * 3 + c / 3][d] = v; }

    private boolean bt(char[][] b, int k) {
        if (k == 81) return true;
        int r = k / 9, c = k % 9;
        if (b[r][c] != '.') return bt(b, k + 1);
        for (int d = 1; d <= 9; d++) {
            if (row[r][d] || col[c][d] || box[r / 3 * 3 + c / 3][d]) continue;
            b[r][c] = (char) ('0' + d);
            mark(r, c, d, true);
            if (bt(b, k + 1)) return true;
            mark(r, c, d, false);
        }
        b[r][c] = '.';
        return false;
    }
}
