// 1249. Minimum Remove to Make Valid Parentheses (Medium)
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Run (Java 11+): java P1249MinimumRemoveToMakeValidParentheses.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1249MinimumRemoveToMakeValidParentheses {
    public static void main(String[] args) {
        check(new Solution().minRemoveToMakeValid("lee(t(c)o)de)"), "lee(t(c)o)de");
        check(new Solution().minRemoveToMakeValid("))(("), "");
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
    public String minRemoveToMakeValid(String s) {
        Deque<Integer> st = new ArrayDeque<>();
        boolean[] bad = new boolean[s.length()];
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(') st.push(i);
            else if (c == ')') { if (st.isEmpty()) bad[i] = true; else st.pop(); }
        }
        for (int i : st) bad[i] = true;
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < s.length(); i++) if (!bad[i]) sb.append(s.charAt(i));
        return sb.toString();
    }
}
