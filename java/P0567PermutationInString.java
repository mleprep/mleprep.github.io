// 567. Permutation in String (Medium)
// https://leetcode.com/problems/permutation-in-string/
// Run (Java 11+): java P0567PermutationInString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0567PermutationInString {
    public static void main(String[] args) {
        check(new Solution().checkInclusion("ab", "eidbaooo"), true);
        check(new Solution().checkInclusion("ab", "eidboaoo"), false);
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
    public boolean checkInclusion(String s1, String s2) {
        int k = s1.length(), n = s2.length();
        if (k > n) return false;
        int[] diff = new int[26];              // diff = count(window) - count(s1)
        for (int i = 0; i < k; i++) { diff[s1.charAt(i) - 'a']--; diff[s2.charAt(i) - 'a']++; }
        int zeros = 0;
        for (int d : diff) if (d == 0) zeros++;
        for (int r = k; ; r++) {
            if (zeros == 26) return true;
            if (r == n) return false;
            int in = s2.charAt(r) - 'a', out = s2.charAt(r - k) - 'a';
            if (diff[in] == 0) zeros--; if (++diff[in] == 0) zeros++;
            if (diff[out] == 0) zeros--; if (--diff[out] == 0) zeros++;
        }
    }
}
