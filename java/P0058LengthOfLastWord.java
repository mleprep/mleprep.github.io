// 58. Length of Last Word (Easy)
// https://leetcode.com/problems/length-of-last-word/
// Run (Java 11+): java P0058LengthOfLastWord.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0058LengthOfLastWord {
    public static void main(String[] args) {
        check(new Solution().lengthOfLastWord("   fly me   to   the moon  "), 4);
        check(new Solution().lengthOfLastWord("a"), 1);
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
    public int lengthOfLastWord(String s) {
        int i = s.length() - 1;
        while (i >= 0 && s.charAt(i) == ' ') i--;
        int j = i;
        while (j >= 0 && s.charAt(j) != ' ') j--;
        return i - j;
    }
}
