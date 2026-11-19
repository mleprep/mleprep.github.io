// 343. Integer Break (Medium)
// https://leetcode.com/problems/integer-break/
// Run (Java 11+): java P0343IntegerBreak.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0343IntegerBreak {
    public static void main(String[] args) {
        check(new Solution().integerBreak(10), 36);
        check(new Solution().integerBreak(2), 1);
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
    public int integerBreak(int n) {
        if (n <= 3) return n - 1;
        int threes = n / 3, rem = n % 3;
        if (rem == 1) { threes--; rem = 4; }          // 3 + 1 -> 2 + 2
        return (int) Math.pow(3, threes) * (rem == 0 ? 1 : rem);
    }
}
