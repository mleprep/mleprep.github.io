// 43. Multiply Strings (Medium)
// https://leetcode.com/problems/multiply-strings/
// Run (Java 11+): java P0043MultiplyStrings.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0043MultiplyStrings {
    public static void main(String[] args) {
        check(new Solution().multiply("123", "456"), "56088");
        check(new Solution().multiply("0", "52"), "0");
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
    public String multiply(String num1, String num2) {
        int m = num1.length(), n = num2.length();
        int[] pos = new int[m + n];
        for (int i = m - 1; i >= 0; i--)
            for (int j = n - 1; j >= 0; j--) {
                int s = (num1.charAt(i) - '0') * (num2.charAt(j) - '0') + pos[i + j + 1];
                pos[i + j + 1] = s % 10;
                pos[i + j] += s / 10;
            }
        StringBuilder sb = new StringBuilder();
        for (int d : pos) if (sb.length() > 0 || d != 0) sb.append(d);
        return sb.length() == 0 ? "0" : sb.toString();
    }
}
