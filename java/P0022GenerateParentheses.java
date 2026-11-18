// 22. Generate Parentheses (Medium)
// https://leetcode.com/problems/generate-parentheses/
// Run (Java 11+): java P0022GenerateParentheses.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0022GenerateParentheses {
    public static void main(String[] args) {
        check(new Solution().generateParenthesis(3), "[((())), (()()), (())(), ()(()), ()()()]");
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
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();
        bt(n, new StringBuilder(), 0, 0, res);
        return res;
    }

    private void bt(int n, StringBuilder cur, int op, int cl, List<String> res) {
        if (cur.length() == 2 * n) { res.add(cur.toString()); return; }
        if (op < n) { cur.append('('); bt(n, cur, op + 1, cl, res); cur.deleteCharAt(cur.length() - 1); }
        if (cl < op) { cur.append(')'); bt(n, cur, op, cl + 1, res); cur.deleteCharAt(cur.length() - 1); }
    }
}
