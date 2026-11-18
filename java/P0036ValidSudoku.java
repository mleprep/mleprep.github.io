// 36. Valid Sudoku (Medium)
// https://leetcode.com/problems/valid-sudoku/
// Run (Java 11+): java P0036ValidSudoku.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0036ValidSudoku {
    public static void main(String[] args) {
        char[][] b = Stream.of("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
                "7...2...6", ".6....28.", "...419..5", "....8..79").map(String::toCharArray).toArray(char[][]::new);
        check(new Solution().isValidSudoku(b), true);
        b[0][0] = '8';
        check(new Solution().isValidSudoku(b), false);
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
    public boolean isValidSudoku(char[][] board) {
        boolean[][] row = new boolean[9][9], col = new boolean[9][9], box = new boolean[9][9];
        for (int r = 0; r < 9; r++)
            for (int c = 0; c < 9; c++) {
                if (board[r][c] == '.') continue;
                int v = board[r][c] - '1', b = (r / 3) * 3 + c / 3;
                if (row[r][v] || col[c][v] || box[b][v]) return false;
                row[r][v] = col[c][v] = box[b][v] = true;
            }
        return true;
    }
}
