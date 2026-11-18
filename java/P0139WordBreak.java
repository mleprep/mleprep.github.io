// 139. Word Break (Medium)
// https://leetcode.com/problems/word-break/
// Run (Java 11+): java P0139WordBreak.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0139WordBreak {
    public static void main(String[] args) {
        check(new Solution().wordBreak("leetcode", List.of("leet", "code")), true);
        check(new Solution().wordBreak("catsandog", List.of("cats", "dog", "sand", "and", "cat")), false);
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
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> words = new HashSet<>(wordDict);
        int L = 0;
        for (String w : words) L = Math.max(L, w.length());
        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true;
        for (int i = 1; i <= s.length(); i++)
            for (int j = Math.max(0, i - L); j < i && !dp[i]; j++)
                dp[i] = dp[j] && words.contains(s.substring(j, i));
        return dp[s.length()];
    }
}
