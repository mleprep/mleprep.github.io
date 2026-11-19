// 1071. Greatest Common Divisor of Strings (Easy)
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Run (Java 11+): java P1071GreatestCommonDivisorOfStrings.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1071GreatestCommonDivisorOfStrings {
    public static void main(String[] args) {
        check(new Solution().gcdOfStrings("ABABAB", "ABAB"), "AB");
        check(new Solution().gcdOfStrings("LEET", "CODE"), "");
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
    public String gcdOfStrings(String str1, String str2) {
        if (!(str1 + str2).equals(str2 + str1)) return "";
        return str1.substring(0, gcd(str1.length(), str2.length()));
    }

    private int gcd(int a, int b) { return b == 0 ? a : gcd(b, a % b); }
}
