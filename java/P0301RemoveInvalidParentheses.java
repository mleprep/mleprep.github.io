// 301. Remove Invalid Parentheses (Hard)
// https://leetcode.com/problems/remove-invalid-parentheses/
// Run (Java 11+): java P0301RemoveInvalidParentheses.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0301RemoveInvalidParentheses {
    public static void main(String[] args) {
        List<String> r = new Solution().removeInvalidParentheses("()())()");
        Collections.sort(r);
        check(r, "[(())(), ()()()]");
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
    public List<String> removeInvalidParentheses(String s) {
        Set<String> level = new HashSet<>(List.of(s));
        while (true) {
            List<String> good = new ArrayList<>();
            for (String t : level) if (valid(t)) good.add(t);
            if (!good.isEmpty()) return good;
            Set<String> next = new HashSet<>();
            for (String t : level)
                for (int i = 0; i < t.length(); i++)
                    if (t.charAt(i) == '(' || t.charAt(i) == ')') next.add(t.substring(0, i) + t.substring(i + 1));
            level = next;
        }
    }

    private boolean valid(String t) {
        int bal = 0;
        for (char c : t.toCharArray()) {
            if (c == '(') bal++;
            else if (c == ')' && --bal < 0) return false;
        }
        return bal == 0;
    }
}
