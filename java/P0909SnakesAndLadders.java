// 909. Snakes and Ladders (Medium)
// https://leetcode.com/problems/snakes-and-ladders/
// Run (Java 11+): java P0909SnakesAndLadders.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0909SnakesAndLadders {
    public static void main(String[] args) {
        int[][] B = {{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
                     {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}};
        check(new Solution().snakesAndLadders(B), 4);
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
    public int snakesAndLadders(int[][] board) {
        int n = board.length;
        boolean[] seen = new boolean[n * n + 1];
        Deque<Integer> q = new ArrayDeque<>(List.of(1));
        seen[1] = true;
        for (int moves = 0; !q.isEmpty(); moves++)
            for (int k = q.size(); k > 0; k--) {
                int s = q.poll();
                if (s == n * n) return moves;
                for (int nxt = s + 1; nxt <= Math.min(s + 6, n * n); nxt++) {
                    int r = (nxt - 1) / n, c = (nxt - 1) % n;
                    if (r % 2 == 1) c = n - 1 - c;
                    int v = board[n - 1 - r][c], dest = v == -1 ? nxt : v;
                    if (!seen[dest]) { seen[dest] = true; q.add(dest); }
                }
            }
        return -1;
    }
}
