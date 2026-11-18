// 137. Single Number II (Medium)
// https://leetcode.com/problems/single-number-ii/
// Run (Java 11+): java P0137SingleNumberIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0137SingleNumberIi {
    public static void main(String[] args) {
        check(new Solution().singleNumber(new int[]{2, 2, 3, 2}), 3);
        check(new Solution().singleNumber(new int[]{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}), -4);
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
        int res = 0;
        for (int b = 0; b < 32; b++) {
            int cnt = 0;
            for (int x : nums) cnt += (x >> b) & 1;
            if (cnt % 3 != 0) res |= 1 << b;
        }
        return res;
    }
}
