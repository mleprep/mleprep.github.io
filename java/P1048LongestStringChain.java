// 1048. Longest String Chain (Medium)
// https://leetcode.com/problems/longest-string-chain/
// Run (Java 11+): java P1048LongestStringChain.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1048LongestStringChain {
    public static void main(String[] args) {
        check(new Solution().longestStrChain(new String[]{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}), 5);
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
    public int longestStrChain(String[] words) {
        Arrays.sort(words, (a, b) -> a.length() - b.length());
        Map<String, Integer> dp = new HashMap<>();
        int best = 0;
        for (String w : words) {
            int cur = 1;
            for (int i = 0; i < w.length(); i++)
                cur = Math.max(cur, dp.getOrDefault(w.substring(0, i) + w.substring(i + 1), 0) + 1);
            dp.put(w, cur);
            best = Math.max(best, cur);
        }
        return best;
    }
}
