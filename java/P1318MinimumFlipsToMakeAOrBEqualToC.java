// 1318. Minimum Flips to Make a OR b Equal to c (Medium)
// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// Run (Java 11+): java P1318MinimumFlipsToMakeAOrBEqualToC.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1318MinimumFlipsToMakeAOrBEqualToC {
    public static void main(String[] args) {
        check(new Solution().minFlips(2, 6, 5), 3);
        check(new Solution().minFlips(4, 2, 7), 1);
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
    public int minFlips(int a, int b, int c) {
        return Integer.bitCount((a | b) ^ c) + Integer.bitCount(a & b & ~c);
    }
}
