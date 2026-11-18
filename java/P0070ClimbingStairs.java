// 70. Climbing Stairs (Easy)
// https://leetcode.com/problems/climbing-stairs/
// Run (Java 11+): java P0070ClimbingStairs.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0070ClimbingStairs {
    public static void main(String[] args) {
        check(new Solution().climbStairs(5), 8);
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
    public int climbStairs(int n) {
        int a = 1, b = 1;
        for (int i = 1; i < n; i++) { int c = a + b; a = b; b = c; }
        return b;
    }
}
