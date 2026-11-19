// 6. Zigzag Conversion (Medium)
// https://leetcode.com/problems/zigzag-conversion/
// Run (Java 11+): java P0006ZigzagConversion.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0006ZigzagConversion {
    public static void main(String[] args) {
        check(new Solution().convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR");
        check(new Solution().convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI");
        check(new Solution().convert("AB", 1), "AB");
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
    public String convert(String s, int numRows) {
        if (numRows == 1 || numRows >= s.length()) return s;
        StringBuilder[] rows = new StringBuilder[numRows];
        for (int i = 0; i < numRows; i++) rows[i] = new StringBuilder();
        int r = 0, step = 1;
        for (char c : s.toCharArray()) {
            rows[r].append(c);
            if (r == 0) step = 1;
            else if (r == numRows - 1) step = -1;
            r += step;
        }
        StringBuilder out = new StringBuilder();
        for (StringBuilder row : rows) out.append(row);
        return out.toString();
    }
}
