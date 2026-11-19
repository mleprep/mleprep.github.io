// 733. Flood Fill (Easy)
// https://leetcode.com/problems/flood-fill/
// Run (Java 11+): java P0733FloodFill.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0733FloodFill {
    public static void main(String[] args) {
        check(new Solution().floodFill(new int[][]{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), new int[][]{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}});
        check(new Solution().floodFill(new int[][]{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0), new int[][]{{0, 0, 0}, {0, 0, 0}});
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
    public int[][] floodFill(int[][] image, int sr, int sc, int color) {
        int old = image[sr][sc];
        if (old != color) fill(image, sr, sc, old, color);
        return image;
    }

    private void fill(int[][] g, int r, int c, int old, int color) {
        if (r < 0 || c < 0 || r >= g.length || c >= g[0].length || g[r][c] != old) return;
        g[r][c] = color;
        fill(g, r + 1, c, old, color);
        fill(g, r - 1, c, old, color);
        fill(g, r, c + 1, old, color);
        fill(g, r, c - 1, old, color);
    }
}
