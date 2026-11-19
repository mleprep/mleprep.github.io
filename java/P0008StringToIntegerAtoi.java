// 8. String to Integer (atoi) (Medium)
// https://leetcode.com/problems/string-to-integer-atoi/
// Run (Java 11+): java P0008StringToIntegerAtoi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0008StringToIntegerAtoi {
    public static void main(String[] args) {
        check(new Solution().myAtoi("   -042"), -42);
        check(new Solution().myAtoi("-91283472332"), -2147483648);
        check(new Solution().myAtoi("2147483648"), 2147483647);
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
    public int myAtoi(String s) {
        int i = 0, n = s.length(), sign = 1, num = 0;
        while (i < n && s.charAt(i) == ' ') i++;
        if (i < n && (s.charAt(i) == '+' || s.charAt(i) == '-')) sign = s.charAt(i++) == '-' ? -1 : 1;
        while (i < n && Character.isDigit(s.charAt(i))) {
            int d = s.charAt(i++) - '0';
            if (num > (Integer.MAX_VALUE - d) / 10) return sign == 1 ? Integer.MAX_VALUE : Integer.MIN_VALUE;
            num = num * 10 + d;
        }
        return sign * num;
    }
}
