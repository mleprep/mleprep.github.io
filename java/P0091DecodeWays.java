// 91. Decode Ways (Medium)
// https://leetcode.com/problems/decode-ways/
// Run (Java 11+): java P0091DecodeWays.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0091DecodeWays {
    public static void main(String[] args) {
        check(new Solution().numDecodings("226"), 3);
        check(new Solution().numDecodings("06"), 0);
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
    public int numDecodings(String s) {
        int a = 1, b = s.charAt(0) != '0' ? 1 : 0;
        for (int i = 2; i <= s.length(); i++) {
            int cur = s.charAt(i - 1) != '0' ? b : 0;
            int two = Integer.parseInt(s.substring(i - 2, i));
            if (two >= 10 && two <= 26) cur += a;
            a = b; b = cur;
        }
        return b;
    }
}
