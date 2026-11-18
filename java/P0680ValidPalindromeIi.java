// 680. Valid Palindrome II (Easy)
// https://leetcode.com/problems/valid-palindrome-ii/
// Run (Java 11+): java P0680ValidPalindromeIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0680ValidPalindromeIi {
    public static void main(String[] args) {
        check(new Solution().validPalindrome("abca"), true);
        check(new Solution().validPalindrome("abc"), false);
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
    public boolean validPalindrome(String s) {
        for (int l = 0, r = s.length() - 1; l < r; l++, r--)
            if (s.charAt(l) != s.charAt(r)) return pal(s, l + 1, r) || pal(s, l, r - 1);
        return true;
    }

    private boolean pal(String s, int l, int r) {
        while (l < r) if (s.charAt(l++) != s.charAt(r--)) return false;
        return true;
    }
}
