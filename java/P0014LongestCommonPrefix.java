// 14. Longest Common Prefix (Easy)
// https://leetcode.com/problems/longest-common-prefix/
// Run (Java 11+): java P0014LongestCommonPrefix.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0014LongestCommonPrefix {
    public static void main(String[] args) {
        check(new Solution().longestCommonPrefix(new String[]{"flower", "flow", "flight"}), "fl");
        check(new Solution().longestCommonPrefix(new String[]{"dog", "racecar", "car"}), "");
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
    public String longestCommonPrefix(String[] strs) {
        String first = strs[0];
        for (int i = 0; i < first.length(); i++)
            for (String s : strs)
                if (i == s.length() || s.charAt(i) != first.charAt(i)) return first.substring(0, i);
        return first;
    }
}
