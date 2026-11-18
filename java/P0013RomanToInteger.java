// 13. Roman to Integer (Easy)
// https://leetcode.com/problems/roman-to-integer/
// Run (Java 11+): java P0013RomanToInteger.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0013RomanToInteger {
    public static void main(String[] args) {
        check(new Solution().romanToInt("MCMXCIV"), 1994);
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
    public int romanToInt(String s) {
        int total = 0;
        for (int i = 0; i < s.length(); i++) {
            int v = val(s.charAt(i));
            total += (i + 1 < s.length() && val(s.charAt(i + 1)) > v) ? -v : v;
        }
        return total;
    }

    private int val(char c) {
        switch (c) {
            case 'I': return 1; case 'V': return 5; case 'X': return 10; case 'L': return 50;
            case 'C': return 100; case 'D': return 500; default: return 1000;
        }
    }
}
