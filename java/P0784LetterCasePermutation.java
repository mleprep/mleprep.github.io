// 784. Letter Case Permutation (Medium)
// https://leetcode.com/problems/letter-case-permutation/
// Run (Java 11+): java P0784LetterCasePermutation.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0784LetterCasePermutation {
    public static void main(String[] args) {
        check(new Solution().letterCasePermutation("3z4"), "[3z4, 3Z4]");
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
    public List<String> letterCasePermutation(String s) {
        List<String> res = new ArrayList<>();
        bt(s.toCharArray(), 0, res);
        return res;
    }

    private void bt(char[] c, int i, List<String> res) {
        if (i == c.length) { res.add(new String(c)); return; }
        if (Character.isLetter(c[i])) {
            c[i] = Character.toLowerCase(c[i]); bt(c, i + 1, res);
            c[i] = Character.toUpperCase(c[i]); bt(c, i + 1, res);
        } else bt(c, i + 1, res);
    }
}
