// 2707. Extra Characters in a String (Medium)
// https://leetcode.com/problems/extra-characters-in-a-string/
// Run (Java 11+): java P2707ExtraCharactersInAString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2707ExtraCharactersInAString {
    public static void main(String[] args) {
        check(new Solution().minExtraChar("leetscode", new String[]{"leet", "code", "leetcode"}), 1);
        check(new Solution().minExtraChar("sayhelloworld", new String[]{"hello", "world"}), 3);
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
    public int minExtraChar(String s, String[] dictionary) {
        Set<String> words = new HashSet<>(Arrays.asList(dictionary));
        int n = s.length();
        int[] dp = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            dp[i] = 1 + dp[i + 1];
            for (int j = i + 1; j <= n; j++) if (words.contains(s.substring(i, j))) dp[i] = Math.min(dp[i], dp[j]);
        }
        return dp[0];
    }
}
