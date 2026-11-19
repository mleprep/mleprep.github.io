// 1197. Minimum Knight Moves (Medium)
// https://leetcode.com/problems/minimum-knight-moves/
// Run (Java 11+): java P1197MinimumKnightMoves.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1197MinimumKnightMoves {
    public static void main(String[] args) {
        check(new Solution().minKnightMoves(5, 5), 4);
        check(new Solution().minKnightMoves(1, 0), 3);
        check(new Solution().minKnightMoves(300, 0), 150);
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
    public int minKnightMoves(int x, int y) {
        x = Math.abs(x); y = Math.abs(y);
        int off = 2;                                   // shift so coordinates >= -2 index from 0
        boolean[][] seen = new boolean[x + 5][y + 5];
        Deque<int[]> q = new ArrayDeque<>();
        q.add(new int[]{0, 0});
        seen[off][off] = true;
        int[][] moves = {{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}};
        for (int d = 0; ; d++) {
            for (int i = q.size(); i > 0; i--) {
                int[] p = q.poll();
                if (p[0] == x && p[1] == y) return d;
                for (int[] m : moves) {
                    int a = p[0] + m[0], b = p[1] + m[1];
                    if (a < -2 || b < -2 || a > x + 2 || b > y + 2 || seen[a + off][b + off]) continue;
                    seen[a + off][b + off] = true;
                    q.add(new int[]{a, b});
                }
            }
        }
    }
}
