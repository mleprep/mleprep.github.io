// 1657. Determine if Two Strings Are Close (Medium)
// https://leetcode.com/problems/determine-if-two-strings-are-close/
// Run (Java 11+): java P1657DetermineIfTwoStringsAreClose.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1657DetermineIfTwoStringsAreClose {
    public static void main(String[] args) {
        check(new Solution().closeStrings("cabbba", "abbccc"), true);
        check(new Solution().closeStrings("uau", "ssx"), false);
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
    public boolean closeStrings(String word1, String word2) {
        if (word1.length() != word2.length()) return false;
        int[] a = new int[26], b = new int[26];
        for (char c : word1.toCharArray()) a[c - 'a']++;
        for (char c : word2.toCharArray()) b[c - 'a']++;
        for (int i = 0; i < 26; i++) if ((a[i] == 0) != (b[i] == 0)) return false;
        Arrays.sort(a);
        Arrays.sort(b);
        return Arrays.equals(a, b);
    }
}
