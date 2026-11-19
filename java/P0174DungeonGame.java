// 174. Dungeon Game (Hard)
// https://leetcode.com/problems/dungeon-game/
// Run (Java 11+): java P0174DungeonGame.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0174DungeonGame {
    public static void main(String[] args) {
        check(new Solution().calculateMinimumHP(new int[][]{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}), 7);
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
    public int calculateMinimumHP(int[][] dungeon) {
        int R = dungeon.length, C = dungeon[0].length;
        int[] need = new int[C + 1];
        Arrays.fill(need, Integer.MAX_VALUE);
        need[C - 1] = 1;
        for (int r = R - 1; r >= 0; r--)
            for (int c = C - 1; c >= 0; c--)
                need[c] = Math.max(1, Math.min(need[c], need[c + 1]) - dungeon[r][c]);
        return need[0];
    }
}
