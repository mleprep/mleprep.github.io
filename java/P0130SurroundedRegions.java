// 130. Surrounded Regions (Medium)
// https://leetcode.com/problems/surrounded-regions/
// Run (Java 11+): java P0130SurroundedRegions.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0130SurroundedRegions {
    public static void main(String[] args) {
        char[][] b = Stream.of("XXXX", "XOOX", "XXOX", "XOXX").map(String::toCharArray).toArray(char[][]::new);
        new Solution().solve(b);
        check(new String(b[1]) + new String(b[3]), "XXXXXOXX");
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
    public void solve(char[][] board) {
        int R = board.length, C = board[0].length;
        for (int r = 0; r < R; r++) { mark(board, r, 0); mark(board, r, C - 1); }
        for (int c = 0; c < C; c++) { mark(board, 0, c); mark(board, R - 1, c); }
        for (char[] row : board)
            for (int c = 0; c < C; c++) row[c] = row[c] == 'S' ? 'O' : 'X';
    }

    private void mark(char[][] b, int r, int c) {
        if (r < 0 || c < 0 || r >= b.length || c >= b[0].length || b[r][c] != 'O') return;
        b[r][c] = 'S';
        mark(b, r + 1, c); mark(b, r - 1, c); mark(b, r, c + 1); mark(b, r, c - 1);
    }
}
