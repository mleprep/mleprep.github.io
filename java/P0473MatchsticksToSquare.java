// 473. Matchsticks to Square (Medium)
// https://leetcode.com/problems/matchsticks-to-square/
// Run (Java 11+): java P0473MatchsticksToSquare.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0473MatchsticksToSquare {
    public static void main(String[] args) {
        check(new Solution().makesquare(new int[]{1, 1, 2, 2, 2}), true);
        check(new Solution().makesquare(new int[]{3, 3, 3, 3, 4}), false);
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
    public boolean makesquare(int[] matchsticks) {
        int total = 0;
        for (int m : matchsticks) total += m;
        if (total % 4 != 0) return false;
        Integer[] sticks = Arrays.stream(matchsticks).boxed().sorted(Collections.reverseOrder()).toArray(Integer[]::new);
        if (sticks[0] > total / 4) return false;
        return place(sticks, 0, new int[4], total / 4);
    }

    private boolean place(Integer[] sticks, int i, int[] sides, int side) {
        if (i == sticks.length) return true;
        for (int k = 0; k < 4; k++) {
            if (sides[k] + sticks[i] <= side) {
                sides[k] += sticks[i];
                if (place(sticks, i + 1, sides, side)) return true;
                sides[k] -= sticks[i];
            }
            if (sides[k] == 0) break;
        }
        return false;
    }
}
