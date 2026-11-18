// 202. Happy Number (Easy)
// https://leetcode.com/problems/happy-number/
// Run (Java 11+): java P0202HappyNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0202HappyNumber {
    public static void main(String[] args) {
        check(new Solution().isHappy(19), true);
        check(new Solution().isHappy(2), false);
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
    public boolean isHappy(int n) {
        int slow = n, fast = next(n);
        while (fast != 1 && slow != fast) { slow = next(slow); fast = next(next(fast)); }
        return fast == 1;
    }

    private int next(int x) {
        int s = 0;
        for (; x > 0; x /= 10) s += (x % 10) * (x % 10);
        return s;
    }
}
