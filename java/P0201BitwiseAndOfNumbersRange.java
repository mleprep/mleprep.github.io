// 201. Bitwise AND of Numbers Range (Medium)
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Run (Java 11+): java P0201BitwiseAndOfNumbersRange.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0201BitwiseAndOfNumbersRange {
    public static void main(String[] args) {
        check(new Solution().rangeBitwiseAnd(5, 7), 4);
        check(new Solution().rangeBitwiseAnd(1, 2147483647), 0);
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
    public int rangeBitwiseAnd(int left, int right) {
        while (right > left) right &= right - 1;   // drop lowest set bits until within the prefix
        return right;
    }
}
