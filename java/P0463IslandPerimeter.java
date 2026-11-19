// 463. Island Perimeter (Easy)
// https://leetcode.com/problems/island-perimeter/
// Run (Java 11+): java P0463IslandPerimeter.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0463IslandPerimeter {
    public static void main(String[] args) {
        check(new Solution().islandPerimeter(new int[][]{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16);
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
    public int islandPerimeter(int[][] grid) {
        int land = 0, shared = 0;
        for (int r = 0; r < grid.length; r++)
            for (int c = 0; c < grid[0].length; c++)
                if (grid[r][c] == 1) {
                    land++;
                    if (r + 1 < grid.length && grid[r + 1][c] == 1) shared++;
                    if (c + 1 < grid[0].length && grid[r][c + 1] == 1) shared++;
                }
        return 4 * land - 2 * shared;
    }
}
