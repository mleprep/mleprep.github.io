// 790. Domino and Tromino Tiling (Medium)
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Run (Java 11+): java P0790DominoAndTrominoTiling.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0790DominoAndTrominoTiling {
    public static void main(String[] args) {
        check(new Solution().numTilings(3), 5);
        check(new Solution().numTilings(5), 24);
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
    public int numTilings(int n) {
        final int MOD = 1_000_000_007;
        if (n < 3) return new int[]{1, 1, 2}[n];
        long a = 1, b = 1, c = 2;
        for (int i = 3; i <= n; i++) { long d = (2 * c + a) % MOD; a = b; b = c; c = d; }
        return (int) c;
    }
}
