// 3133. Minimum Array End (Medium)
// https://leetcode.com/problems/minimum-array-end/
// Run (Java 11+): java P3133MinimumArrayEnd.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P3133MinimumArrayEnd {
    public static void main(String[] args) {
        check(new Solution().minEnd(3, 4), 6L);
        check(new Solution().minEnd(2, 7), 15L);
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
    public long minEnd(int n, int x) {
        long res = x, k = n - 1;
        for (long bit = 1; k > 0; bit <<= 1) {
            if ((x & bit) == 0) {
                if ((k & 1) == 1) res |= bit;
                k >>= 1;
            }
        }
        return res;
    }
}
