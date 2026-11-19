// 28. Find the Index of the First Occurrence in a String (Easy)
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Run (Java 11+): java P0028FindTheIndexOfTheFirstOccurrenceInAString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0028FindTheIndexOfTheFirstOccurrenceInAString {
    public static void main(String[] args) {
        check(new Solution().strStr("sadbutsad", "sad"), 0);
        check(new Solution().strStr("aaaaab", "aab"), 3);
        check(new Solution().strStr("leetcode", "leeto"), -1);
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
    public int strStr(String haystack, String needle) {
        int m = needle.length();
        int[] lps = new int[m];
        for (int i = 1, k = 0; i < m; i++) {
            while (k > 0 && needle.charAt(i) != needle.charAt(k)) k = lps[k - 1];
            if (needle.charAt(i) == needle.charAt(k)) k++;
            lps[i] = k;
        }
        for (int i = 0, k = 0; i < haystack.length(); i++) {
            while (k > 0 && haystack.charAt(i) != needle.charAt(k)) k = lps[k - 1];
            if (haystack.charAt(i) == needle.charAt(k) && ++k == m) return i - m + 1;
        }
        return -1;
    }
}
