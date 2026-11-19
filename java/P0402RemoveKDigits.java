// 402. Remove K Digits (Medium)
// https://leetcode.com/problems/remove-k-digits/
// Run (Java 11+): java P0402RemoveKDigits.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0402RemoveKDigits {
    public static void main(String[] args) {
        check(new Solution().removeKdigits("1432219", 3), "1219");
        check(new Solution().removeKdigits("10200", 1), "200");
        check(new Solution().removeKdigits("10", 2), "0");
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
    public String removeKdigits(String num, int k) {
        StringBuilder st = new StringBuilder();
        for (char d : num.toCharArray()) {
            while (k > 0 && st.length() > 0 && st.charAt(st.length() - 1) > d) { st.deleteCharAt(st.length() - 1); k--; }
            st.append(d);
        }
        st.setLength(st.length() - k);
        int i = 0;
        while (i < st.length() && st.charAt(i) == '0') i++;
        return i == st.length() ? "0" : st.substring(i);
    }
}
