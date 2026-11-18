// 17. Letter Combinations of a Phone Number (Medium)
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Run (Java 11+): java P0017LetterCombinationsOfAPhoneNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0017LetterCombinationsOfAPhoneNumber {
    public static void main(String[] args) {
        check(new Solution().letterCombinations("23"), "[ad, ae, af, bd, be, bf, cd, ce, cf]");
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
    private static final String[] PAD = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        if (!digits.isEmpty()) bt(digits, 0, new StringBuilder(), res);
        return res;
    }

    private void bt(String d, int i, StringBuilder cur, List<String> res) {
        if (i == d.length()) { res.add(cur.toString()); return; }
        for (char c : PAD[d.charAt(i) - '0'].toCharArray()) {
            cur.append(c);
            bt(d, i + 1, cur, res);
            cur.deleteCharAt(cur.length() - 1);
        }
    }
}
