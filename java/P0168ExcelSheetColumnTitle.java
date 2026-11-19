// 168. Excel Sheet Column Title (Easy)
// https://leetcode.com/problems/excel-sheet-column-title/
// Run (Java 11+): java P0168ExcelSheetColumnTitle.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0168ExcelSheetColumnTitle {
    public static void main(String[] args) {
        check(new Solution().convertToTitle(701), "ZY");
        check(new Solution().convertToTitle(52), "AZ");
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
    public String convertToTitle(int columnNumber) {
        StringBuilder sb = new StringBuilder();
        for (int n = columnNumber; n > 0; n /= 26) {
            n--;
            sb.append((char) ('A' + n % 26));
        }
        return sb.reverse().toString();
    }
}
