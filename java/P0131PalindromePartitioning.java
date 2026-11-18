// 131. Palindrome Partitioning (Medium)
// https://leetcode.com/problems/palindrome-partitioning/
// Run (Java 11+): java P0131PalindromePartitioning.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0131PalindromePartitioning {
    public static void main(String[] args) {
        check(new Solution().partition("aab"), "[[a, a, b], [aa, b]]");
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
    public List<List<String>> partition(String s) {
        int n = s.length();
        boolean[][] pal = new boolean[n][n];
        for (int i = n - 1; i >= 0; i--)
            for (int j = i; j < n; j++)
                pal[i][j] = s.charAt(i) == s.charAt(j) && (j - i < 2 || pal[i + 1][j - 1]);
        List<List<String>> res = new ArrayList<>();
        bt(s, 0, pal, new ArrayList<>(), res);
        return res;
    }

    private void bt(String s, int i, boolean[][] pal, List<String> path, List<List<String>> res) {
        if (i == s.length()) { res.add(new ArrayList<>(path)); return; }
        for (int j = i; j < s.length(); j++) {
            if (!pal[i][j]) continue;
            path.add(s.substring(i, j + 1));
            bt(s, j + 1, pal, path, res);
            path.remove(path.size() - 1);
        }
    }
}
