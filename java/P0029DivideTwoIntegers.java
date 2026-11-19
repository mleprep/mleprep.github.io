// 29. Divide Two Integers (Medium)
// https://leetcode.com/problems/divide-two-integers/
// Run (Java 11+): java P0029DivideTwoIntegers.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0029DivideTwoIntegers {
    public static void main(String[] args) {
        check(new Solution().divide(7, -3), -2);
        check(new Solution().divide(-2147483648, -1), 2147483647);
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
    public int divide(int dividend, int divisor) {
        if (dividend == Integer.MIN_VALUE && divisor == -1) return Integer.MAX_VALUE;
        boolean neg = (dividend < 0) != (divisor < 0);
        long a = Math.abs((long) dividend), b = Math.abs((long) divisor), q = 0;
        while (a >= b) {
            int k = 0;
            while (a >= (b << (k + 1))) k++;
            a -= b << k;
            q += 1L << k;
        }
        return (int) (neg ? -q : q);
    }
}
