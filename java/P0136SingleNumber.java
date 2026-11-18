// 136. Single Number (Easy)
// https://leetcode.com/problems/single-number/
// Run (Java 11+): java P0136SingleNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0136SingleNumber {
    public static void main(String[] args) {
        check(new Solution().singleNumber(new int[]{4, 1, 2, 1, 2}), 4);
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
    public int singleNumber(int[] nums) {
        int x = 0;
        for (int v : nums) x ^= v;
        return x;
    }
}
