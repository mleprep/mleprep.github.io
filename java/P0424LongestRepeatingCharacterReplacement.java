// 424. Longest Repeating Character Replacement (Medium)
// https://leetcode.com/problems/longest-repeating-character-replacement/
// Run (Java 11+): java P0424LongestRepeatingCharacterReplacement.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0424LongestRepeatingCharacterReplacement {
    public static void main(String[] args) {
        check(new Solution().characterReplacement("ABAB", 2), 4);
        check(new Solution().characterReplacement("AABABBA", 1), 4);
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
    public int characterReplacement(String s, int k) {
        int[] cnt = new int[26];
        int l = 0, maxf = 0;
        for (int r = 0; r < s.length(); r++) {
            maxf = Math.max(maxf, ++cnt[s.charAt(r) - 'A']);
            if (r - l + 1 - maxf > k) cnt[s.charAt(l++) - 'A']--;
        }
        return s.length() - l; // window never shrinks, so its final size is the best size
    }
}
