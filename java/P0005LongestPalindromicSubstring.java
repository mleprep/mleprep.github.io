// 5. Longest Palindromic Substring (Medium)
// https://leetcode.com/problems/longest-palindromic-substring/
// Run (Java 11+): java P0005LongestPalindromicSubstring.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0005LongestPalindromicSubstring {
    public static void main(String[] args) {
        check(new Solution().longestPalindrome("cbbd"), "bb");
        check(new Solution().longestPalindrome("babad"), "bab");
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
    public String longestPalindrome(String s) {
        int bl = 0, br = 0;
        for (int c = 0; c < 2 * s.length() - 1; c++) {
            int l = c / 2, r = (c + 1) / 2;
            while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) { l--; r++; }
            if (r - l - 1 > br - bl) { bl = l + 1; br = r; }
        }
        return s.substring(bl, br);
    }
}
