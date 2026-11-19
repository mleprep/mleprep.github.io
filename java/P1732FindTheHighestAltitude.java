// 1732. Find the Highest Altitude (Easy)
// https://leetcode.com/problems/find-the-highest-altitude/
// Run (Java 11+): java P1732FindTheHighestAltitude.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1732FindTheHighestAltitude {
    public static void main(String[] args) {
        check(new Solution().largestAltitude(new int[]{-5, 1, 5, 0, -7}), 1);
        check(new Solution().largestAltitude(new int[]{-4, -3, -2, -1, 4, 3, 2}), 0);
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
    public int largestAltitude(int[] gain) {
        int alt = 0, best = 0;
        for (int g : gain) best = Math.max(best, alt += g);
        return best;
    }
}
