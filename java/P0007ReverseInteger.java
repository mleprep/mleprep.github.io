// 7. Reverse Integer (Medium)
// https://leetcode.com/problems/reverse-integer/
// Run (Java 11+): java P0007ReverseInteger.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0007ReverseInteger {
    public static void main(String[] args) {
        check(new Solution().reverse(-123), -321);
        check(new Solution().reverse(1534236469), 0);
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
    public int reverse(int x) {
        int res = 0;
        while (x != 0) {
            int d = x % 10;                      // keeps sign in Java
            if (res > Integer.MAX_VALUE / 10 || res < Integer.MIN_VALUE / 10) return 0;
            res = res * 10 + d;
            x /= 10;
        }
        return res;
    }
}
