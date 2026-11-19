// 1137. N-th Tribonacci Number (Easy)
// https://leetcode.com/problems/n-th-tribonacci-number/
// Run (Java 11+): java P1137NThTribonacciNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1137NThTribonacciNumber {
    public static void main(String[] args) {
        check(new Solution().tribonacci(25), 1389537);
        check(new Solution().tribonacci(37), 2082876103);
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
    public int tribonacci(int n) {
        int a = 0, b = 1, c = 1;
        for (int i = 0; i < n; i++) { int d = a + b + c; a = b; b = c; c = d; }
        return a;
    }
}
