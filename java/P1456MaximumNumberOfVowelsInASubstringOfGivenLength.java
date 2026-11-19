// 1456. Maximum Number of Vowels in a Substring of Given Length (Medium)
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Run (Java 11+): java P1456MaximumNumberOfVowelsInASubstringOfGivenLength.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1456MaximumNumberOfVowelsInASubstringOfGivenLength {
    public static void main(String[] args) {
        check(new Solution().maxVowels("abciiidef", 3), 3);
        check(new Solution().maxVowels("leetcode", 3), 2);
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
    public int maxVowels(String s, int k) {
        int cur = 0, best = 0;
        for (int i = 0; i < s.length(); i++) {
            if ("aeiou".indexOf(s.charAt(i)) >= 0) cur++;
            if (i >= k && "aeiou".indexOf(s.charAt(i - k)) >= 0) cur--;
            best = Math.max(best, cur);
        }
        return best;
    }
}
