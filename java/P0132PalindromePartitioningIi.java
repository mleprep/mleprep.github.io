// 132. Palindrome Partitioning II (Hard)
// https://leetcode.com/problems/palindrome-partitioning-ii/
// Run (Java 11+): java P0132PalindromePartitioningIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0132PalindromePartitioningIi {
    public static void main(String[] args) {
        check(new Solution().minCut("aab"), 1);
        check(new Solution().minCut("abccbc"), 2);
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
    public int minCut(String s) {
        int n = s.length();
        int[] cuts = new int[n + 1];
        for (int i = 0; i <= n; i++) cuts[i] = i - 1;
        for (int c = 0; c < n; c++)
            for (int[] lr : new int[][]{{c, c}, {c, c + 1}})
                for (int l = lr[0], r = lr[1]; l >= 0 && r < n && s.charAt(l) == s.charAt(r); l--, r++)
                    cuts[r + 1] = Math.min(cuts[r + 1], cuts[l] + 1);
        return cuts[n];
    }
}
