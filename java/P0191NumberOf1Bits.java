// 191. Number of 1 Bits (Easy)
// https://leetcode.com/problems/number-of-1-bits/
// Run (Java 11+): java P0191NumberOf1Bits.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0191NumberOf1Bits {
    public static void main(String[] args) {
        check(new Solution().hammingWeight(11), 3);
        check(new Solution().hammingWeight(-3), 31);
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
    public int hammingWeight(int n) {
        int count = 0;
        for (; n != 0; n &= n - 1) count++;   // works for negative (unsigned) inputs too
        return count;
    }
}
