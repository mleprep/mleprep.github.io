// 50. Pow(x, n) (Medium)
// https://leetcode.com/problems/powx-n/
// Run (Java 11+): java P0050PowxN.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0050PowxN {
    public static void main(String[] args) {
        check(new Solution().myPow(2.0, 10), 1024.0);
        check(new Solution().myPow(2.0, -2), 0.25);
        check(new Solution().myPow(1.0, Integer.MIN_VALUE), 1.0);
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
    public double myPow(double x, int n) {
        long e = n;
        if (e < 0) { x = 1 / x; e = -e; }
        double res = 1;
        for (; e > 0; e >>= 1) {
            if ((e & 1) == 1) res *= x;
            x *= x;
        }
        return res;
    }
}
